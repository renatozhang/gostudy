package db

import (
	"database/sql"
	"fmt"

	"github.com/renatozhang/gostudy/blogger/model"
)

func InsertArticle(article *model.ArticleDetail) (articleId int64, err error) {
	if article == nil {
		err = fmt.Errorf("invalid article parameter")
		return
	}
	sqlstr := `insert into article(
		category_id,content,title,
		view_count,comment_count,username,summary) 
		values(?,?,?,?,?,?,?)`

	result, err := DB.Exec(sqlstr, article.Category.CategoryId,
		article.Content, article.Title, article.ViewCount,
		article.CommentCount, article.UserName, article.Summary)
	if err != nil {
		return
	}
	articleId, err = result.LastInsertId()
	return
}

func GetArticleList(pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid parameter, page_num:%d, page_size:%d", pageNum, pageSize)
		return
	}
	sqlstr := `select
				id,category_id,summary,title,view_count,
				create_time,comment_count,username
			from
				article
			where
				status=1
			order by create_time desc
			limit ?,?`
	err = DB.Select(&articleList, sqlstr, pageNum, pageSize)
	if err != nil {
		return
	}
	return
}

func GetArticleListByCategoryId(categoryId, pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid parameter, page_num:%d, page_size:%d", pageNum, pageSize)
		return
	}
	sqlstr := `select
				id,category_id,summary,title,view_count,
				create_time,comment_count,username
			from
				article
			where
				category_id=?
			and
				status=1
			order by create_time desc
			limit ?,?`
	err = DB.Select(&articleList, sqlstr, categoryId, pageNum, pageSize)
	if err != nil {
		return
	}
	return
}

func GetArticleDetail(articleId int) (articleInfo *model.ArticleDetail, err error) {
	if articleId < 0 {
		err = fmt.Errorf("invalid parmeter, article_id:%d", articleId)
		return
	}

	articleInfo = &model.ArticleDetail{}

	sqlstr := `select
					id, summary,title,view_count,content,
					create_time,comment_count,username,category_id
			 from
				 article
			where
				id=? 
			and
				status=1`
	err = DB.Get(articleInfo, sqlstr, articleId)
	return
}

func GetRelativeArticle(articleId int64) (articleList []*model.RelativeArticle, err error) {
	var categoryId int64
	sqlstr := "select category_id from article where id=?"
	err = DB.Get(&categoryId, sqlstr, articleId)
	if err != nil {
		return
	}
	sqlstr = "select id, title from article where category_id=? and id !=? limit 10"
	err = DB.Select(&articleList, sqlstr, categoryId, articleId)
	return
}

func GetPrevArticleById(articleId int64) (info *model.RelativeArticle, err error) {
	info = &model.RelativeArticle{
		ArticleId: -1,
	}
	sqlstr := "select id,title from article where id < ? order by id desc limit 1"
	err = DB.Get(info, sqlstr, articleId)
	if err != nil {
		return
	}
	return
}

func GetNextArticleById(articleId int64) (info *model.RelativeArticle, err error) {
	info = &model.RelativeArticle{
		ArticleId: -1,
	}
	sqlstr := "select id,title from article where id > ? order by id asc limit 1"
	err = DB.Get(info, sqlstr, articleId)
	if err != nil {
		return
	}
	return
}

func IsArticleExist(articleId int64) (exists bool, err error) {
	var id int64
	sqlstr := "select id from article where id=?"
	err = DB.Get(&id, sqlstr, articleId)
	if err == sql.ErrNoRows {
		exists = false
		return
	}
	if err != nil {
		return
	}
	return

}
