package controllers_v1

import (
	"fmt"
	"github.com/pvfm/enube/api/database"
	"github.com/pvfm/enube/api/models"
	"github.com/pvfm/enube/api/services"

	"github.com/gin-gonic/gin"
)

type UserSession struct {
	Email     string `json:"email"`
	Password string `json:"password"`
}


// Create user godos
// @Sumarry     Initialize Session
// @Description Initialize Session
// @Tags        session
// @Accept      json
// @Produce     json
// @Param       request             body      UserSession       true                     "user request"
// @Success     200                 {object}  map[string]string "{"token": "success"}"
// @Failure     500                 {object}  map[string]string "{"message": "failure"}"
// @Router      /login [post]
func AuthenticateSession(c *gin.Context) {
	db := database.GetDatabase()

	var userSession UserSession
	var user models.User

	err := c.ShouldBindJSON(&userSession)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Password or email is incorrect",
		})
		return
	}

	err = db.Where("email = ?", userSession.Email).First(&user).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Password or email are incorrect",
		})
		return
	}

	fmt.Println(user)
	validSession := services.CheckPasswordHash(userSession.Password, user.Password)

	if validSession == false {
		c.JSON(500, gin.H{
			"message": "Password or email are incorrect",
		})
		return
	}

	tokenizeData := map[string]string{
		"email": user.Email,
		"name":  user.Name,
		"id":    string(user.ID),
	}

	token, err := services.GenerateToken(tokenizeData)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Password or email is incorrect",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": token,
	})
}
