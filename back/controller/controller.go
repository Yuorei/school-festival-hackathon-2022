package controller

import (
	"fmt"
	"lendingAndBorrowing/firebaseOperation"
	"lendingAndBorrowing/operateDb"
	"net/http"
	"strconv"

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
	User_id     int    `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image_url   string `json:"image_url"`
	Deadline    string `json:"deadline"`
}
type Upload_image_url struct {
	Image_url string `json:"image_url"`
}

type Res_lists struct {
	Uuid        string `json:"id"`
	User_id     int    `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image_url   string `json:"image_url"`
	Deadline    string `json:"deadline"`
}

// /users
func GetUsers(c *gin.Context) {
	var res User_register
	if err := c.Bind(&res); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	c.JSON(http.StatusOK, res)
}
func DeleteUsers(c *gin.Context) {
	var user User_res
	if err := c.Bind(&user); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	c.JSON(http.StatusOK, "StatusOK")
}

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
func GetLendThing(c *gin.Context) { //uuid
	id := c.Param("uuid")
	var res Rent_lists
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
func PostLendLists(c *gin.Context) {
	var lists Rent_lists
	var Rent_list Res_lists
	if err := c.Bind(&lists); err != nil {
		c.String(http.StatusBadRequest, "bad request")
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
	if err := db.Create(&Rent_list); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	c.JSON(http.StatusOK, Rent_list)
}

func PutLendLists(c *gin.Context) {
	id := c.Param("uuid")
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

// /upload-image
func PostUploadImage(c *gin.Context) {
	var image_url Upload_image_url
	var url = "https://storage.googleapis.com/school-festival-hackathon.appspot.com/"
	var image_extension = ".webp"
	if err := c.Bind(&image_url); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return
	}
	uuid_str := u.String()
	image_url.Image_url = url + uuid_str + image_extension

	// formdataのkeyの"file"を受け取る
	image_file, _, err := c.Request.FormFile("file")
	bucket := firebaseOperation.UseDefaultBacket()
	if err := firebaseOperation.UploadFile(bucket, uuid_str+image_extension, image_file); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	c.JSON(http.StatusOK, image_url)
}
