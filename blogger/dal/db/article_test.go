package db

import (
	"testing"
	"time"

	"github.com/renatozhang/gostudy/blogger/model"
)

func init() {
	dsn := "root:123456@tcp(localhost:3306)/blog?parseTime=true"
	err := Init(dsn)
	if err != nil {
		panic(err)
	}
}

func TestInsertArticle(t *testing.T) {
	article := &model.ArticleDetail{}
	article.ArticleInfo.CategoryId = 1
	article.Content = "this a test ak dkdkdkddkddkd"
	article.ArticleInfo.Title = "GOLANG 连接Mysql的时区问题"
	article.ArticleInfo.ViewCount = 1
	article.ArticleInfo.CommentCount = 0
	article.ArticleInfo.UserName = "少林之巅"
	article.ArticleInfo.Summary = `使用mysql的时间字段遇到如下两个问题
	1.使用go-sql-driver来连接mysql数据库，获取的时区默认是UTC +0的，与本地的东八区是有区别，在业务处理中会出现问题
	2.获取mysql中的日期，是string类型，需要在代码中用time.Parse进行转化`
	article.ArticleInfo.CreateTime = time.Now()
	article.Category.CategoryId = 1

	articleId, err := InsertArticle(article)
	if err != nil {
		t.Errorf("inset article failed, err:%v\n", err)
		return
	}
	t.Logf("insert article succ, articleId:%v\n", articleId)
}

func TestGetArticleList(t *testing.T) {
	articleList, err := GetArticleList(1, 15)
	if err != nil {
		t.Errorf("get article failed, err:%v\n", err)
		return
	}
	t.Logf("get article succ, len:%d\n", len(articleList))
}

func TestGetArticleDetail(t *testing.T) {
	articleInfo, err := GetArticleDetail(7)
	if err != nil {
		t.Errorf("get article failed, err:%n\n", err)
		return
	}
	t.Logf("get article succ, article:%#v\n", articleInfo)

}

func TestGetRelativeArticle(t *testing.T) {
	articleList, err := GetRelativeArticle(1)
	if err != nil {
		t.Errorf("get relative article failed, err:%v\n", err)
	}

	for _, v := range articleList {
		t.Logf("id:%d title:%s\n", v.ArticleId, v.Title)
	}

}
