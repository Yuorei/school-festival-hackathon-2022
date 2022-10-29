package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type User_register struct {
	Jwd string `json:"jwd"`
}

type User_res struct {
	IsNewUser bool `json:"isNewUser"`
}

type Rent_lists struct {
	Uuid        string    `json:"uuid"`
	User_id     int       `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image_url   string    `json:"image_url"`
	Deadline    time.Time `json:"deadline"`
}

func GetUsers(c *gin.Context) {
	var user User_register
	var res User_res
	if err := c.Bind(&user); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	if err := c.Bind(&res); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	c.JSON(http.StatusOK, res)
}
func DeleteUsers(c *gin.Context) {
	var user User_register
	if err := c.Bind(&user); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	c.JSON(http.StatusOK, "StatusOK")
}
func GetAllRentLists(c *gin.Context) {
	var res Rent_lists
	if err := c.Bind(&res); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	c.JSON(http.StatusOK, res)
}
func GetSingleRentList(c *gin.Context) { //todo funcName
	var lists Rent_lists
	if err := c.Bind(&lists); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	c.JSON(http.StatusOK, lists)
}

func PostRentLists(c *gin.Context) {
	var lists Rent_lists
	var res Rent_lists
	if err := c.Bind(&lists); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	if err := c.Bind(&res); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	c.JSON(http.StatusOK, res)
}
func PutRentLists(c *gin.Context) {
	var lists Rent_lists
	if err := c.Bind(&lists); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

}
func DeleteRentLists(c *gin.Context) {
	c.JSON(http.StatusOK, "StatusOK")
}
func GetLendLists(c *gin.Context) {
	var res Rent_lists
	if err := c.Bind(&res); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	c.JSON(http.StatusOK, res)
}
func PostLendLists(c *gin.Context) {
	var lists Rent_lists
	var res Rent_lists
	if err := c.Bind(&lists); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	if err := c.Bind(&lists); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	c.JSON(http.StatusOK, res)
}

func PutLendLists(c *gin.Context) {
	var res Rent_lists
	if err := c.Bind(&res); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	c.JSON(http.StatusOK, res)
}
func DeleteLendList(c *gin.Context) {
	c.JSON(http.StatusOK, "StatusOK")
}
