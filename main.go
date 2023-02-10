package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	// Set the router
	router = gin.Default()

	router.GET("/users", getUsers)

	//router.GET("/article/view/:article_id")

	router.Run()
}

func getUsers(ctx *gin.Context) {
	users := GetUsers()
	if users == nil || len(users) == 0 {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		fmt.Println(users)
		ctx.IndentedJSON(http.StatusOK, users)
	}
}
