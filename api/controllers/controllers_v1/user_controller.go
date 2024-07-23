package controllers_v1

import (
	"github.com/gin-gonic/gin"

	"github.com/pvfm/enube/api/database"
	"github.com/pvfm/enube/api/models"
)


// Create user godos
// @Sumarry     Create User
// @Description Create User
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       request       body      models.User   true  "user request"
// @Success     201           {object}  models.User
// @Failure     500                 {object}  map[string]string "{"message": "failure"}"
// @Router      /users [post]
func RegisterUser(c *gin.Context) {
	db := database.GetDatabase()

	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Something happend contact support",
		})
		return
	}

	err = db.Create(&user).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Somithing happend contact support",
		})
		return
	}

	c.JSON(201, gin.H{
		"data": user,
	})
}

// Create user godos
// @Sumarry     Show User
// @Description Show User
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       Authorization header    string      true "With the bearer token"
// @Success     200           {object}  models.User
// @Router      /user [get]
func ShowUser(c *gin.Context) {
	currentUser := c.Value("currentUser")

	c.JSON(201, gin.H{
		"data": currentUser,
	})
}
