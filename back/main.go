package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/users/", hoge)
	router.DELETE("/users", hoge)
	router.GET("/rent-lists", hoge)
	router.GET("/rent-lists/:id", hoge)
	router.POST("/rent-lists", head)
	router.PUT("/rent-lists/:id", hoge)
	router.DELETE("/rent-lists/:id", hoge)
	router.GET("/lend-lists", hoge)
	router.GET("/lend-lists/:id", hoge)
	router.POST("/lend-lists", hoge)
	router.PUT("/lend-lists/:id", hoge)
	router.DELETE("/lend-lists/:id", hoge)

	router.Run(":3000")
}
