package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"publisher-end/handler"
	"publisher-end/middleware"
	"time"
)

func GetRouter() *gin.Engine {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))
	router.Use(middleware.RecordTimeMiddleware())

	router.Static("/static", "static")

	// 验证相关
	auth := router.Group("/auth")
	{
		auth.POST("/login", handler.HandleLogin)
		auth.POST("/user", handler.GetUserInfo)
	}

	//　demo
	router.GET("/secret", middleware.JWTMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	// 编辑文件相关
	editor := router.Group("/media")
	{
		editor.POST("/upload", handler.HandleImageUpload)
		editor.POST("/editor", handler.HandlePost)
	}

	//获取博客相关
	blog := router.Group("/blog")
	{
		blog.GET("/posts", handler.HandleGetPosts)
		blog.GET("/post", handler.HandleGetPostByID)
	}

	//评论相关
	comment := router.Group("/comment")
	{
		comment.POST("/save", handler.HandleSaveComment)
		comment.GET("/list", handler.HandleGetComments)
	}



	return router
}
