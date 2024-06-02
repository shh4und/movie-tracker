package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shh4und/movie-tracker/handler"
)

func initRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/title/:title", handler.GetTitle)
		v1.GET("/search/:title", handler.GetSearch)
	}
}
