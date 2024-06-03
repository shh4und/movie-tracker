package router

import (
	"github.com/gin-gonic/gin"
)

func Init() {

	// Initialize router (gin.Engine) with the default configs
	router := gin.Default()

	// Initialize routes
	initRoutes(router)

	// Run server
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
