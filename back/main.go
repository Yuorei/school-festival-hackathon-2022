package main

import (
	"lendingAndBorrowing/controller"
	"lendingAndBorrowing/firebaseOperation"
	"lendingAndBorrowing/middleware"
	"lendingAndBorrowing/operateDb"

	"github.com/gin-gonic/gin"
	//"lendingAndBorrowing/operateDb"
	firebase "firebase.google.com/go/v4"
)

var firebase_app *firebase.App

func main() {
	firebase_app = firebaseOperation.Init()
	// DB open
	if err := operateDb.Init(); err != nil {
		panic("DBerror")
	}

	router := gin.Default()
	router.Use(middleware.Middleware)
	// /users
	router.POST("/users/", controller.GetUsers)
	router.DELETE("/users", controller.DeleteUsers)

	// /rent-lists
	router.GET("/rent-lists", controller.GetAllRentLists) //実装あとで
	router.GET("/rent-lists/:id", controller.GetSingleRentList)
	router.POST("/rent-lists", controller.PostRentLists)
	router.PUT("/rent-lists/:id", controller.PutRentLists)
	router.DELETE("/rent-lists/:id", controller.DeleteRentLists)

	// /lend-lists
	router.GET("/lend-lists", controller.GetLendLists) //実装あとで
	router.GET("/lend-lists/:id", controller.GetLendThing)
	router.POST("/lend-lists", controller.PostLendLists)
	router.PUT("/lend-lists/:id", controller.PutLendLists)
	router.DELETE("/lend-lists/s:id", controller.DeleteLendList)

	// sever
	router.Run(":3000")
}

func Firebase_app() *firebase.App {
	return firebase_app
}
