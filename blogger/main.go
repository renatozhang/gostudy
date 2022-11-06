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
	// 文章评论相关的处理
	router.POST("/comment/submit/", controller.CommentSubmit)
	// 留言相关
	router.GET("/leave/new/", controller.LeaveDetail)
	router.POST("/leave/new/", controller.LeaveSubmit)
	// 分类下面的文章列表
	router.GET("/category/", controller.CategoryList)

	router.Run()

}
