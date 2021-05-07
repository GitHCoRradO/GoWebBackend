package routers

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/githcorrado/fake-admagic/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER"))

	r.Use(static.Serve("/", static.LocalFile("./public", true)))
	r.NoRoute(func(c *gin.Context) {
		c.File("./public/index.html")
	})

	AddGroupUser(r)
	AddGroupWSRoutes(r)

	return r
}
