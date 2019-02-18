package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"publisher-end/database"
	"strconv"
)

func HandlePost(c *gin.Context) {
	id := c.PostForm("id")
	title := c.PostForm("title")
	htmlStr := c.PostForm("html")
	rawStr := c.PostForm("raw")
	status := c.PostForm("status")
	fmt.Println("fuck ", title, htmlStr)
	_, err := database.SaveOrPublishPost(id, title, htmlStr, rawStr, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "稍后再试",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "操作成功",
	})
}

func HandleGetPosts(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	pagesize, err := strconv.Atoi(c.Query("pagesize"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}

	count, posts, err := database.GetPosts(page, pagesize);
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "获取数据失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"count": count,
		"posts": posts,
		"page": page,
		"pagesize": pagesize,
	})
}

func HandleGetPostByID(c *gin.Context) {
	id := c.Query("id")
	post,err := database.GetPostByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "获取数据失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}