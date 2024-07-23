package routes

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/pvfm/enube/api/docs"

	"github.com/pvfm/enube/api/routes/routes_v1"
)

func SetRoutesV1(r *gin.RouterGroup) {
  docs.SwaggerInfo.BasePath = "/api/v1"

	routesV1 := r.Group("/v1")
	{
		routes_v1.SetUserRoutesV1(routesV1)
		routes_v1.SetSessionRoutesV1(routesV1)
		routes_v1.SetImportRoutesV1(routesV1)

		routesV1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
}

