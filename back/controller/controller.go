package controller

import (
	"encoding/base64"
	"lendingAndBorrowing/firebaseOperation"
	"lendingAndBorrowing/operateDb"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type User_register struct {
	Jwd string `json:"jwd"`
}

type User_res struct {
	IsNewUser bool `json:"isNewUser"`
}

type Rent_lists struct {
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
	src := res.Image_url
	dec, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		log.Fatal(err)
	}
	bucket := firebaseOperation.UseDefaultBacket()
	u ,_:= uuid.NewRandom()
	filename := u.String()
	if err := firebaseOperation.UploadFile(bucket, filename, string(dec)); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	res.Image_url = "https://school-festival-hackathon.appspot.com" + filename
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

func PutRentLists(c *gin.Context) {
	id := c.Param("id")
	//stringからintにキャスト
	int_id, _ := strconv.Atoi(id)
	var res Rent_lists
	if err := c.Bind(&res); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	res.User_id = int_id
	db := operateDb.GetConnect()
	if err := db.Save(&res); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	c.JSON(http.StatusOK, res)
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
	db := operateDb.GetConnect()
	// Get the first record ordered by primary key
	if err := db.First(&res); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	c.JSON(http.StatusOK, res)
}
func GetLendThing(c *gin.Context) { //id
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
func PostLendLists(c *gin.Context) {
	var lists Rent_lists
	if err := c.Bind(&lists); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	db := operateDb.GetConnect()
	src := lists.Image_url
	dec, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		log.Fatal(err)
	}
	bucket := firebaseOperation.UseDefaultBacket()
	u,_ := uuid.NewRandom()
	filename := u.String()
	lists.Image_url = "https://school-festival-hackathon.appspot.com" + filename
	if err := firebaseOperation.UploadFile(bucket, filename, string(dec)); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	if err := db.Create(&lists); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	c.JSON(http.StatusOK, lists)
}

func PutLendLists(c *gin.Context) {
	id := c.Param("id")
	//stringからintにキャスト
	int_id, _ := strconv.Atoi(id)
	var res Rent_lists
	if err := c.Bind(&res); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	res.User_id = int_id
	db := operateDb.GetConnect()
	if err := db.Save(&res); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	c.JSON(http.StatusOK, res)
}
func DeleteLendList(c *gin.Context) {
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
