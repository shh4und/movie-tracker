package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shh4und/movie-tracker/handler"
)

func initRoutes(router *gin.Engine) {

	handler.InitHandler()

	titles := router.Group("/api/v1/search")
	{
		// fetch a list of a title search
		titles.GET("", handler.GetSearch)
	}

	users := router.Group("/api/v1/users")
	{

		// CRUD user routes
		users.POST("/register", handler.CreateUser) // create
		users.POST("/login", handler.LoginUser)     // authenticate
		// users.GET("/profile/id", handler.GetUserProfileByID)
		users.GET("/profile/username", handler.GetUserProfileByUsername)
		users.PUT("/profile", handler.UpdateUser)
		users.DELETE("/profile", handler.DeleteUser)

	}

}
