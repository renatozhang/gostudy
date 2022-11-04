package main

import (
	"github.com/gin-gonic/gin"
	"github.com/renatozhang/gostudy/blogger/controller"
	"github.com/renatozhang/gostudy/blogger/dal/db"
)

func main() {
	router := gin.Default()

	dsn := "root:123456@tcp(localhost:3306)/blog?parseTime=true"
	err := db.Init(dsn)
	if err != nil {
		panic(err)
	}

	router.Static("static", "./static")
	router.LoadHTMLGlob("views/*")

	router.GET("/", controller.IndexHandle)

	// 发布文章页面
	router.GET("/article/new/", controller.NewArticle)
	// 文章提交接口
	router.POST("/article/submit/", controller.ArticleSubmit)
	// 获取文章详情
	router.GET("/article/detail/", controller.ArticleDetail)

	router.Run()

}
