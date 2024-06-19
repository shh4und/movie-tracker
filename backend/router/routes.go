package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shh4und/movie-tracker/handler"
)

func initRoutes(router *gin.Engine) {

	handler.InitHandler()

	titles := router.Group("/api/search")
	{
		// fetch a single title (directly search)
		titles.GET("/title/:title", handler.GetTitle)
		// fetch a list of a title search
		titles.GET("/:title", handler.GetSearch)
	}

	users := router.Group("/api/users")
	{

		// CRUD user routes
		users.POST("/register", handler.CreateUser) // create
		users.POST("/login")                        // authenticate
		users.GET("/profile/id", handler.GetUserProfileByID)
		users.GET("/profile/username", handler.GetUserProfileByUsername)
		users.PUT("/profile", handler.UpdateUser)
		users.DELETE("/profile", handler.DeleteUser)

	}

	comments := router.Group("/api/comments")
	{
		// comments routes
		comments.POST("/:userID/:titleID")
		comments.DELETE("/:userID/:commentID")
		comments.GET("/title/:titleID")
		comments.GET("/user/:userID")
	}

}
