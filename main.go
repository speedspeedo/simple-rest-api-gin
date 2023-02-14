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

	router.POST("/register", addUser)

	// var users = []User{
	// 	{Id: 1, Name: "Randy Cummings", Password: "123456", Email: "speedodeveloper1004@gmail.com"},
	// 	{Id: 2, Name: "Andy Cummings", Password: "123456", Email: "speedodeveloper1004@gmail.com"},
	// 	{Id: 3, Name: "Pandy Cummings", Password: "123456", Email: "speedodeveloper1004@gmail.com"},
	// }

	// b, _ := json.Marshal(users)

	// fmt.Println(string(b))

	router.Run()
}

func getUsers(ctx *gin.Context) {
	users := GetUsers()
	if users == nil || len(users) == 0 {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.IndentedJSON(http.StatusOK, users)
	}
}

// curl -X POST http://localhost:8080/users -H 'content-type:application/json' -d "{\"id\":2,\"name\":\"asd\",\"email\":\"asd\",\"password\":\"sd\"}"

func addUser(ctx *gin.Context) {
	var user User

	if err := ctx.BindJSON(&user); err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
	} else {
		AddUser(user)
		ctx.IndentedJSON(http.StatusCreated, user)
	}
}
