package main

import "github.com/gin-gonic/gin"

var router *gin.Engine

func main() {

	// Set the router
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Handle Index
	router.GET("/")
	// Handle GET requests at /article/view/some_article_id
	router.GET("/article/view/:article_id")

	router.Run()
}
