package routers

import (
	"server/api"
	"server/setting"

	"github.com/gin-gonic/gin"
)

// InitRouter router init regsiter
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiR := r.Group("/api")

	{
		apiR.GET("/tags", api.GetTags)

		apiR.POST("/tags", api.AddTag)

		apiR.PUT("/tags/:id", api.EditTag)

		apiR.DELETE("/tags/:id", api.DeleteTag)
	}

	return r
}
