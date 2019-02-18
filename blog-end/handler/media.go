package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
	"path"
)

//　用于编辑文本时图片的上传
func HandleImageUpload(c *gin.Context) {
	image, _ := c.FormFile("image")
	ext := path.Ext(image.Filename)
	var imageName = image.Filename
	u2, err := uuid.NewV4()
	if err == nil {
		imageName = u2.String() + ext
	}
	err = c.SaveUploadedFile(image, "static/" + imageName)
	if err != nil {
		fmt.Println("Error %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "文件无法保存",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"url": "http://" + c.Request.Host + "/static/" + imageName,
		})
	}
}
