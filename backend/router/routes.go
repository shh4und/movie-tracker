package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shh4und/movie-tracker/auth"
	"github.com/shh4und/movie-tracker/config"
	"github.com/shh4und/movie-tracker/handler"
)

func initRoutes(router *gin.Engine) {

	handler.InitHandler()
	secret := []byte(config.Envs.JwtToken)

	// Public routes
	public := router.Group("/api/v1")
	{
		public.POST("/register", handler.CreateUser) // create
		public.POST("/login", handler.LoginUser)     // authenticate
		// fetch a list of a title search
		public.GET("/titles/search", handler.GetSearch)
	}

	// Protected routes
	protected := router.Group("/api/v1")
	protected.Use(auth.Authenticate(secret)) // using middleware to verify the authenticate session
	{
		// CRUD user routes
		protected.GET("/users/profile", handler.GetUserProfileByUsername)
		protected.PUT("/users/update", handler.UpdateUser)
		protected.DELETE("/users/delete", handler.DeleteUser)
	}

}
