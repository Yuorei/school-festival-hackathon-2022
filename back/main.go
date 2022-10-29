package main

import (
	"lendingAndBorrowing/controller"
	"lendingAndBorrowing/operateDb"

	"github.com/gin-gonic/gin"
	//"lendingAndBorrowing/operateDb"
)

func main() {
	router := gin.Default()
	if err:=operateDb.Init();err!=nil{
		panic("DBerror")
	}

	router.POST("/users/", controller.GetUsers)
	router.DELETE("/users", controller.DeleteUsers)
	router.GET("/rent-lists", controller.GetAllRentLists) //実装あとで
	router.GET("/rent-lists/:id", controller.GetSingleRentList)
	router.POST("/rent-lists", controller.PostRentLists)
	router.PUT("/rent-lists/:id", controller.PutRentLists)
	router.DELETE("/rent-lists/:id", controller.DeleteRentLists)
	router.GET("/lend-lists", controller.GetLendLists) //実装あとで
	router.GET("/lend-lists/:id", controller.GetLendThing)
	router.POST("/lend-lists", controller.PostLendLists)
	router.PUT("/lend-lists/:id", controller.PutLendLists)
	router.DELETE("/lend-lists/s:id", controller.DeleteLendList)
	router.Run(":3000")
}
