package routes_v1

import (
	"github.com/gin-gonic/gin"

	"github.com/pvfm/enube/api/controllers/controllers_v1"
	"github.com/pvfm/enube/api/middleware"
)

func SetImportRoutesV1(r *gin.RouterGroup) {
	importRoutes := r.Group("/imports").Use(middleware.Authenticate())
	{
		importRoutes.GET("", controllers_v1.GetImport)
	}
}
