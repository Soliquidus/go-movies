package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes() *gin.Engine {
	// Use Gin as router
	router := gin.Default()
	router.Use(CORS)

	// Set route for index
	router.GET("/", indexView)

	// Set routes for API
	router.GET("/v1/movies", GetMovies)
	router.GET("/v1/movies/:id", GetMovie)

	// Set up Gin Server
	return router
}

// Function called for index
func indexView(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3006")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, GET, PUT")

	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, "OK")
		return
	}

	c.Next()
}

// CORS Middleware
func CORS(c *gin.Context) {

	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {

		c.Next()

	} else {

		// Everytime we receive an OPTIONS request,
		// we just return an HTTP 200 Status Code
		// Like this, Angular can now do the real
		// request using any other method than OPTIONS
		c.AbortWithStatus(http.StatusOK)
	}
}
