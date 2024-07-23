package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pvfm/enube/api/database"
	"github.com/pvfm/enube/api/models"
	"github.com/pvfm/enube/api/services"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		data, err := services.DecodeToken(tokenString)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized access" )
		}

		db := database.GetDatabase()
		var currentUser models.User

		db.Where("email = ?", data["email"]).First(&currentUser)
		c.Set("currentUser", currentUser)
	}
}
