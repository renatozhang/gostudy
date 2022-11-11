package main

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"github.com/renatozhang/gostudy/mercury/controller/account"
	"github.com/renatozhang/gostudy/mercury/controller/ask"
	"github.com/renatozhang/gostudy/mercury/controller/category"
	"github.com/renatozhang/gostudy/mercury/dal/db"
	"github.com/renatozhang/gostudy/mercury/id_gen"
	"github.com/renatozhang/gostudy/mercury/session"
)

func initTemplate(router *gin.Engine) {
	router.StaticFile("/", "./static/index.html")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")
	router.Static("/css/", "./static/css/")
	router.Static("/fonts/", "./static/fonts/")
	router.Static("/img/", "./static/img/")
	router.Static("/js/", "./static/js/")
}

func initDb() (err error) {
	dsn := "root:123456@tcp(localhost:3306)/mercury?parseTime=true"
	err = db.Init(dsn)
	if err != nil {
		return
	}
	return
}

func initSession() (err error) {
	err = session.Init("memory", "")
	if err != nil {
		return
	}
	return
}

func main() {
	router := gin.Default()

	err := initDb()
	if err != nil {
		panic(err)
	}

	err = initSession()
	if err != nil {
		panic(err)
	}

	err = id_gen.Init(1)
	if err != nil {
		panic(err)
	}

	ginpprof.Wrapper(router)
	initTemplate(router)

	router.POST("/api/user/register", account.RegisterHandle)
	router.POST("/api/user/login", account.LoginHandle)
	router.GET("/api/category/list", category.GetCategoryListHandle)
	router.GET("/api/ask/submit", ask.QuestionSubmitHandle)

	router.Run(":9090")

}
