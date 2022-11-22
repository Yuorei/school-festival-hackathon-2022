package controller

import (
	"fmt"
	"lendingAndBorrowing/operateDb"
	"net/http"
	"strconv"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

// /rent-lists
func GetAllRentLists(c *gin.Context) {
	var res Res_lists
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

func GetSingleRentList(c *gin.Context) { //uuid
	id := c.Param("uuid")
	var res Res_lists
	if err := c.Bind(&res); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	// Get the first record ordered by primary key
	db := operateDb.GetConnect()
	if err := db.Where("uuid = ?", id).First(&res); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	c.JSON(http.StatusOK, res)
}

func PostRentLists(c *gin.Context) {
	var lists operateDb.Rent_list
	var Rent_list Res_lists
	if err := c.Bind(&lists); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return
	}
	uuid_str := u.String()

	Rent_list.Uuid = uuid_str
	Rent_list.User_id = lists.User_id
	Rent_list.Name = lists.Name
	Rent_list.Description = lists.Description
	Rent_list.Image_url = lists.Image_url
	Rent_list.Deadline = lists.Deadline

	db := operateDb.GetConnect()
	if err := db.Create(&lists).Error; err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, Rent_list)
}

func PutRentLists(c *gin.Context) {
	id := c.Param("uuid")
	//stringからintにキャスト
	int_id, _ := strconv.Atoi(id)
	var res Rent_lists
	if err := c.Bind(&res); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	res.User_id = int_id
	db := operateDb.GetConnect()
	if err := db.Save(&res).Error; err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
func DeleteRentLists(c *gin.Context) {
	id := c.Param("uuid")
	var lists Rent_lists
	if err := c.Bind(&lists); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	db := operateDb.GetConnect()
	if err := db.Where("uuid = ?", id).Delete(&lists); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	c.JSON(http.StatusOK, "StatusOK")
}
