package routes_v1

import (
	"github.com/gin-gonic/gin"

	"github.com/pvfm/enube/api/controllers/controllers_v1"
	"github.com/pvfm/enube/api/middleware"
)


func SetUserRoutesV1(r *gin.RouterGroup) {
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("", controllers_v1.RegisterUser)
	}

	userRoutesAuth := r.Group("/user").Use(middleware.Authenticate())
	{
		userRoutesAuth.GET("", controllers_v1.ShowUser)
	}
}
