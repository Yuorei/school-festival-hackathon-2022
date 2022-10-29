package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/users/", GetUsers)
	router.DELETE("/users", DeleteUsers)
	router.GET("/rent-lists", GetAllRentLists) //実装あとで
	router.GET("/rent-lists/:id", GetSingleRentList)
	router.POST("/rent-lists", PostRentLists)
	router.PUT("/rent-lists/:id", PutRentLists)
	router.DELETE("/rent-lists/:id", DeleteRentLists)
	router.GET("/lend-lists", GetLendLists) //実装あとで
	router.GET("/lend-lists/:id", GetLendThing)
	router.POST("/lend-lists", PostLendLists)
	router.PUT("/lend-lists/:id", PutLendLists)
	router.DELETE("/lend-lists/s:id", DeleteLendList)
	router.Run(":3000")
}
