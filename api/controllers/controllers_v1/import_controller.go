package controllers_v1

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/pvfm/enube/api/database"
)

// Create user godos
// @Sumarry     List Imports
// @Description List Imports
// @Tags        imports
// @Accept      json
// @Produce     json
// @Param       page          query     string            false  "Page of result"     default(0)
// @Param       limit         query     string            false  "Query limit"        default(50)
// @Param       offset        query     string            false  "Offset result"      default(0)
// @Param       Authorization header    string            true   "With the bearer token"
// @Success     200           {object}  map[string]string
// @Failure     500           {object}  map[string]string "{"message": "failure"}"
// @Router      /imports [get]
func GetImport(c *gin.Context) {
	 db := database.GetDatabase()

	 limitQuery := c.DefaultQuery("limit", "50")
	 offsetQuery := c.DefaultQuery("offset", "0")
	 pageQuery := c.DefaultQuery("page", "0")

	 limit, _ := strconv.Atoi(limitQuery)
	 offset, _ := strconv.Atoi(offsetQuery)
	 page, _ := strconv.Atoi(pageQuery)

	 rows, err := db.Table("imports").Order("id asc").Limit(limit).Offset(offset + limit * page).Rows()

	 if err != nil {
		 c.JSON(400, gin.H{
			 "message": "Something is happend contact the support",
		 })
	 }

	 var columns []string
	 columns, err = rows.Columns()

	 colNum := len(columns)
	 var results []map[string]interface{}

	 for rows.Next() {
		 r := make([]interface{}, colNum)
		 for i := range r {
			 r[i] = &r[i]
		 }

		 err = rows.Scan(r...)
		 if err != nil {
			 c.JSON(400, gin.H{
				 "message": "Something is happend contact the support",
			 })
		 }

		 var row = map[string]interface{}{}
		 for i := range r {
			 row[columns[i]] = r[i]
		 }

		 results = append(results, row)
	 }

	 c.JSON(200, gin.H{
		 "data": results,
	 })
}
