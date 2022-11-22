package controller

import (
	"fmt"
	"lendingAndBorrowing/firebaseOperation"
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

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
