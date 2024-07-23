package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"context"
	"strings"
	"time"

	"github.com/thedatashed/xlsxreader"
	"github.com/joho/godotenv"
	"github.com/jackc/pgx/v5"
)


func main() {
	k := time.Now()
	file, err := xlsxreader.OpenFile(os.Args[1])
	defer file.Close()
	if err != nil {
		log.Fatal("Error while reading the file ", err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	conn, err := pgx.Connect(context.Background(), databaseUrl)
	defer conn.Close(context.Background())

	rows, _ := conn.Query(context.Background(), "select exists(select 1 from pg_tables where schemaname = 'public' and tablename = 'file_imports')")
	result, _ := pgx.CollectRows(rows, pgx.RowTo[bool])

	if result [0] == false {
		_, err := conn.Exec(context.Background(), "create table file_imports (filename text)")
		if err != nil {
			log.Fatal("Error to create table", err)
		}
	}

	validateFile := fmt.Sprintf("select exists(select 1 from file_imports where filename = '%s')", os.Args[1])
	rows, _ = conn.Query(context.Background(), validateFile)
	result, _ = pgx.CollectRows(rows, pgx.RowTo[bool])

	if result[0] == false {
		insertFile := fmt.Sprintf("insert into file_imports (filename) values ('%s')", os.Args[1])
		conn.Exec(context.Background(), insertFile)
	} else {
		log.Fatal("File is already imported")
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	matchers     := make(map[string]string)
	ch           := make(chan map[string]string, 20)
	batchChannel := make(chan []map[string]string)

	go func() {
		defer close(ch)
		for row := range file.ReadRows(file.Sheets[0]) {
			if row.Index == 1 {
				SetMatchers(row, matchers)
			} else {
				ParseRow(row, matchers, ch)
			}
		}
	}()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		batch := []map[string]string{}

		for value := range ch {
			batch = append(batch, value)
			if len(batch) == 10 {
				batchChannel <- batch
				batch = []map[string]string{}
			}
		}

		if len(batch) > 0 {
			batchChannel <- batch
		}
	}()


	go func() {
		wg.Wait()
		close(batchChannel)
	}()

	for value := range batchChannel {
		WriteInDatabase(value, matchers, conn)
	}

	fmt.Println(k)
	fmt.Println(time.Now())
}

func ParseRow(row xlsxreader.Row, matcher map[string]string, ch chan <- map[string]string) {
	obj := map[string]string{}

	for _, cell := range row.Cells {
		obj[matcher[cell.Column]] = cell.Value
	}

	ch <- obj
}

func SetMatchers(row xlsxreader.Row, matcher map[string]string) {
	for _, cell := range row.Cells {
		matcher[cell.Column] = cell.Value
	}
}

func SanitizedColumns(columns []string) []string {
	santizedColumns := []string{}
	for _, c := range columns {
		santizedColumns = append(santizedColumns, strings.ToLower(c))
	}

	return santizedColumns
}

func WriteInDatabase(data []map[string]string, matchers map[string]string, conn *pgx.Conn) {
	rows, _ := conn.Query(context.Background(), "select exists(select 1 from pg_tables where schemaname = 'public' and tablename = 'imports')")
	result, _ := pgx.CollectRows(rows, pgx.RowTo[bool])

	columns := []string{}
	tableColumns := ""

	for _, value := range matchers {
		tableColumns += fmt.Sprintf("%s text ", value)
		columns = append(columns, value)
	}

	if result[0] == false {
		createTable := fmt.Sprintf("create table imports (id serial primary key, %s)", strings.Replace(tableColumns, "text ", "text, ", len(matchers) - 1))

		_, err := conn.Exec(context.Background(), createTable)

		if err != nil {
			log.Fatal("Error to create table", err)
		}
	}

	insertRows := [][]any{}

	for _, d := range data {
		row := []any{}

		fmt.Println(d)
		fmt.Println(columns)

		for _, column := range columns {
			row = append(row, d[column] )
		}
		insertRows = append(insertRows, row)
	}

	sanitizedColumns := SanitizedColumns(columns)

	copyCount, err := conn.CopyFrom(
		context.Background(),
		pgx.Identifier{"imports"},
		sanitizedColumns,
		pgx.CopyFromRows(insertRows),
	)


	if err != nil {
		log.Fatal("Error to insert data in database: ", err)
	}

	fmt.Println(copyCount)
}

