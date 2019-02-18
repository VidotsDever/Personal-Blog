package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"publisher-end/database"
	"strconv"
)

func HandleSaveComment(c *gin.Context) {
	content := c.PostForm("content")
	post_id := c.PostForm("post_id")
	commenter_id := c.PostForm("commenter_id")
	replyer_id := c.PostForm("replyer_id")
	replyer_name := c.PostForm("replyer_name")
	parent_id_str := c.PostForm("parent_id")
	parent_id, err := strconv.Atoi(parent_id_str)
	if err != nil {
		fmt.Println("HandleSaveComment - %v", err)
		parent_id = 0
	}
	comment_id, err := database.SaveComment(content, post_id, commenter_id, replyer_id, replyer_name, parent_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "内部错误",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"comment_id": comment_id,
	})
}

func HandleGetComments(c *gin.Context) {
	post_id := c.Query("post_id")
	comments := database.GetComments(post_id)
	if comments != nil {
		c.JSON(http.StatusOK, gin.H{
			"comments": comments,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"comments": nil,
		})
	}
}
