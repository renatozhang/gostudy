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
	c.HTML(http.StatusOK, "views/index.html", articleRecordList)
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

	c.HTML(http.StatusOK, "views/detail.html", articledetail)

}
