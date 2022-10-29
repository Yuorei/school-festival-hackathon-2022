package controller

import (
	"lendingAndBorrowing/operateDb"
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
type Put_rent_lists struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image_url   string    `json:"image_url"`
	Deadline    time.Time `json:"deadline"`
}

type Rent_lists struct {
	Uuid        string    `json:"uuid"`
	User_id     int       `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image_url   string    `json:"image_url"`
	Deadline    time.Time `json:"deadline"`
}

// /users
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

// /rent-lists
func GetAllRentLists(c *gin.Context) {
	var res Rent_lists
	if err := c.Bind(&res); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	db := operateDb.GetConnect()
	// Get the first record ordered by primary key
	if err := db.First(&res); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	c.JSON(http.StatusOK, res)
}

func GetSingleRentList(c *gin.Context) {
	id := c.Param("id")
	var lists Rent_lists
	if err := c.Bind(&lists); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	// Get the first record ordered by primary key
	db := operateDb.GetConnect()
	if err := db.Where("user_id = ?", id).First(&lists); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	c.JSON(http.StatusOK, lists)
}

func PostRentLists(c *gin.Context) {
	var lists Rent_lists
	if err := c.Bind(&lists); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	db := operateDb.GetConnect()
	if err := db.Create(&lists); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	c.JSON(http.StatusOK, lists)
}

// 聞くidつかっていなくないか？
func PutRentLists(c *gin.Context) { //id
	var part_lists Put_rent_lists
	var res Rent_lists
	if err := c.Bind(&part_lists); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	if err := c.Bind(&res); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	db := operateDb.GetConnect()
	if err := db.First(&part_lists); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	if err := db.Save(&res); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

}
func DeleteRentLists(c *gin.Context) {
	id := c.Param("id")
	var lists Rent_lists
	if err := c.Bind(&lists); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	db := operateDb.GetConnect()
	if err := db.Where("user_id = ?", id).Delete(&lists); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	c.JSON(http.StatusOK, "StatusOK")
}

// /lend-lists
func GetLendLists(c *gin.Context) {
	var res Rent_lists
	if err := c.Bind(&res); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	c.JSON(http.StatusOK, res)
}
func GetLendThing(c *gin.Context) { //id
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

func PutLendLists(c *gin.Context) { //id
	var res Rent_lists
	if err := c.Bind(&res); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	c.JSON(http.StatusOK, res)
}
func DeleteLendList(c *gin.Context) { //id
	c.JSON(http.StatusOK, "StatusOK")
}
