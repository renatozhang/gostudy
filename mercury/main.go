package main

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"github.com/renatozhang/gostudy/mercury/controller/account"
	"github.com/renatozhang/gostudy/mercury/controller/answer"
	"github.com/renatozhang/gostudy/mercury/controller/category"
	"github.com/renatozhang/gostudy/mercury/controller/comment"
	"github.com/renatozhang/gostudy/mercury/controller/question"
	"github.com/renatozhang/gostudy/mercury/dal/db"
	"github.com/renatozhang/gostudy/mercury/filter"
	"github.com/renatozhang/gostudy/mercury/id_gen"
	"github.com/renatozhang/gostudy/mercury/logger"
	maccount "github.com/renatozhang/gostudy/mercury/middleware/account"
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

func initFilter() (err error) {
	err = filter.Init("./data/filter.dat.txt")
	if err != nil {
		logger.Debug("init filter failed,err:%v", err)
		return
	}
	logger.Debug("init filter succ")
	return
}

func main() {
	router := gin.Default()

	config := make(map[string]string)
	config["log_lervel"] = "debug"
	logger.InitLogger("console", config)

	err := initDb()
	if err != nil {
		panic(err)
	}

	err = initFilter()
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
	router.POST("/api/ask/submit", maccount.AuthMiddleware, question.QuestionSubmitHandle)
	router.GET("/api/question/list", category.GetQustionListHandle)
	router.GET("/api/question/detail", question.QustionDetailHandle)
	router.GET("/api/answer/list", answer.AnswerListHandle)

	// 评论模块
	commentGroup := router.Group("/api/comment/", maccount.AuthMiddleware)
	{

		// 发表评论接口
		commentGroup.POST("/post_comment", comment.PostCommentHandle)
		// 发表回复接口
		commentGroup.POST("/post_reply", comment.PostReplyHandle)
		//获取评论列表接口
		commentGroup.GET("/list", comment.CommentListHandle)
		// 获取回复列表接口
		commentGroup.GET("/reply_list", comment.ReplyListHandle)
		//点赞接口，
		commentGroup.POST("/like", comment.LikeHandle)
	}
	router.Run(":9090")

}
