package controller

import (
	"fmt"
	"lendingAndBorrowing/firebaseOperation"
	"lendingAndBorrowing/operateDb"
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
type Image_file struct {
}
type Upload_image_url struct {
	Image_url string `json:"image_url"`
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
		c.String(http.StatusBadRequest, "bad request1")
		return
	}
	db := operateDb.GetConnect()
	if err := db.Create(&lists); err != nil {
		c.String(http.StatusBadRequest, "bad request2")
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

// /upload-image
func PostUploadImage(c *gin.Context) {
	var file Image_file
	var image_url Upload_image_url
	var url = "https://storage.googleapis.com/school-festival-hackathon.appspot.com/"
	var image_extension = ".jpg"
	if err := c.Bind(&file); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	if err := c.Bind(&image_url); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return
	}
	uu := u.String()
	image_url.Image_url = url + uu + image_extension
	image_file, _, err := c.Request.FormFile("file")
	bucket := firebaseOperation.UseDefaultBacket()
	if err := firebaseOperation.UploadFile(bucket, image_url.Image_url, image_file); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	c.JSON(http.StatusOK, image_url)
}
