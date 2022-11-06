package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/renatozhang/gostudy/blogger/logic"
)

func IndexHandle(c *gin.Context) {
	articleRecordList, err := logic.GetArticleRecordList(0, 15)
	if err != nil {
		fmt.Printf("get article failed,err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	// fmt.Printf("articleList:%#v\n", articleRecordList)

	ALLcategoryList, err := logic.GetALLCategoryList()
	if err != nil {
		fmt.Printf("get category failed, err:%v\n", err)
	}
	var data map[string]interface{} = make(map[string]interface{})
	data["article_list"] = articleRecordList
	data["category_list"] = ALLcategoryList
	c.HTML(http.StatusOK, "views/index.html", data)
}

func NewArticle(c *gin.Context) {
	categoryList, err := logic.GetALLCategoryList()
	if err != nil {
		fmt.Printf("get category list failed, err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/post_article.html", categoryList)
}

func ArticleSubmit(c *gin.Context) {
	author := c.PostForm("username")
	title := c.PostForm("title")
	content := c.PostForm("content")
	categoryIdStr := c.PostForm("category_id")

	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	err = logic.InsertArticle(author, title, content, categoryId)
	if err != nil {
		fmt.Printf("insert article failed, err:%v", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/")
}

func ArticleDetail(c *gin.Context) {
	articleIdStr := c.Query("article_id")
	articleId, err := strconv.ParseInt(articleIdStr, 10, 64)
	fmt.Println("articleId:", articleId)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	articledetail, err := logic.GetArticleDetail(int(articleId))
	if err != nil {
		fmt.Printf("get article detail failed, err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	relativeArticle, err := logic.GetRelativeArticle(articleId)
	if err != nil {
		fmt.Printf("get relative article failed, err:%v\n", err)
	}

	prevArticle, nextArticle, err := logic.GetPrevAndNextArticleInfo(articleId)
	if err != nil {
		fmt.Printf("get pre or next article failed, err:%v\n", err)
	}

	categoryList, err := logic.GetALLCategoryList()
	if err != nil {
		fmt.Printf("get category list failed,err:%v\n", err)
	}

	commentList, err := logic.GetCommentList(articleId)
	if err != nil {
		fmt.Printf("get comment list failed:%v\n", err)
	}

	m := make(map[string]interface{})
	m["detail"] = articledetail
	m["relative_article"] = relativeArticle
	m["prev"] = prevArticle
	m["next"] = nextArticle
	m["category"] = categoryList
	m["article_id"] = articleId
	m["comment_list"] = commentList

	c.HTML(http.StatusOK, "views/detail.html", m)
}

func CommentSubmit(c *gin.Context) {
	content := c.PostForm("comment")
	author := c.PostForm("author")
	articleIDStr := c.PostForm("article_id")
	articleId, err := strconv.ParseInt(articleIDStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	err = logic.InsertComment(content, author, articleId)
	if err != nil {
		fmt.Printf("insert comment err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	url := fmt.Sprintf("/article/detail/?article_id=%d", articleId)
	c.Redirect(http.StatusMovedPermanently, url)
}

func LeaveDetail(c *gin.Context) {
	leaveList, err := logic.GetLeaveList(0, 100)
	if err != nil {
		fmt.Printf("get leave list err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/gbook.html", leaveList)
}

func LeaveSubmit(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	content := c.PostForm("content")
	err := logic.InsertLeave(username, email, content)
	if err != nil {
		fmt.Printf("insert leave err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/leave/new/")
}

func CategoryList(c *gin.Context) {
	categoryIdStr := c.Query("category_id")
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	articleRecordList, err := logic.GetArticleRecordListById(int(categoryId), 0, 15)
	if err != nil {
		fmt.Printf("get article failed,err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	// fmt.Printf("articleList:%#v\n", articleRecordList)

	ALLcategoryList, err := logic.GetALLCategoryList()
	if err != nil {
		fmt.Printf("get category failed, err:%v\n", err)
	}
	var data map[string]interface{} = make(map[string]interface{})
	data["article_list"] = articleRecordList
	data["category_list"] = ALLcategoryList
	c.HTML(http.StatusOK, "views/index.html", data)
}
