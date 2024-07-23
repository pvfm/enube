package routes_v1

import (
	"github.com/pvfm/enube/api/controllers/controllers_v1"
	"github.com/gin-gonic/gin"
)

func SetSessionRoutesV1(r *gin.RouterGroup) {
	sessionRoutes := r.Group("/login")
	{
		sessionRoutes.POST("", controllers_v1.AuthenticateSession)
	}
}
