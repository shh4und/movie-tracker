package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shh4und/movie-tracker/handler"
)

func initRoutes(router *gin.Engine) {
	titles := router.Group("/api/titles")
	{
		// fetch a single title (directly search)
		titles.GET("/title/:titleName", handler.GetTitle)
		// fetch a list of a title search
		titles.GET("/search/:titleName", handler.GetSearch)
	}

	users := router.Group("/api/users")
	{

		// CRUD user routes
		users.POST("/register")
		users.POST("/login")
		users.GET("/profile/:userID")
		users.PUT("/profile/:userID")
		users.DELETE("/profile/:userID")

	}

	comments := router.Group("/api/comments")
	{
		// comments routes
		comments.POST("/:userID/:titleID")
		comments.DELETE("/:userID/:commentID")
		comments.GET("/:titleID")
		comments.GET("/:userID")
	}

}
