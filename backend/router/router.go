package router

import (
	"github.com/gin-gonic/gin"
)

func Init() {

	// Initialize router (gin.Engine) with the default configs
	router := gin.Default()

	// Serve static files from the "static" directory
	router.Static("/static", "./static")
	// Serve the main HTML files directly
	router.StaticFile("/index", "./static/index.html")
	router.StaticFile("/register", "./static/register.html")
	router.StaticFile("/login", "./static/login.html")
	// Initialize routes
	initRoutes(router)

	// Run server
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
