package logic

import (
	"fmt"
	"time"

	"github.com/renatozhang/gostudy/blogger/dal/db"
	"github.com/renatozhang/gostudy/blogger/model"
)

func InsertComment(comment, author string, articleId int64) (err error) {
	// 1.首先，要验证article_id是否合法
	exist, err := db.IsArticleExist(articleId)
	if err != nil {
		fmt.Printf("query database failed, err:%v\n", err)
		return
	}
	if exist == false {
		fmt.Errorf("article id:%d not fount", articleId)
		return
	}

	// 2.调用dal InsertComment进行评论内容的插入
	var c model.Comment
	c.Content = comment
	c.Username = author
	c.ArticleId = int(articleId)
	c.CreateTime = time.Now()
	c.Status = 1
	err = db.InsertComment(&c)
	return
}

func GetCommentList(articleId int64) (commentList []*model.Comment, err error) {
	// 1.首先，要验证article_id是否合法
	exist, err := db.IsArticleExist(articleId)
	if err != nil {
		fmt.Printf("query database failed,err:%v\n", err)
		return
	}
	if exist == false {
		err = fmt.Errorf("article id :%d not found", articleId)
		return
	}

	// 2.调用dal InsertComment进行评论内容的插入
	commentList, err = db.GetCommentList(articleId, 0, 100)
	return
}
