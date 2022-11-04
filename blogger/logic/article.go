package logic

import (
	"fmt"
	"math"

	"github.com/renatozhang/gostudy/blogger/dal/db"
	"github.com/renatozhang/gostudy/blogger/model"
)

func getCategoryIds(articleInfoList []*model.ArticleInfo) (ids []int64) {
LABEL:
	for _, article := range articleInfoList {
		categoryId := article.CategoryId
		for _, id := range ids {
			if id == int64(categoryId) {
				continue LABEL
			}
		}
		ids = append(ids, int64(categoryId))
	}
	return
}

func GetArticleRecordList(page_num, page_size int) (articleRecordList []*model.ArticleRecord, err error) {
	//1.从数据库中，获取文章列表
	articleInfoList, err := db.GetArticleList(page_num, page_size)
	if err != nil {
		fmt.Printf("1 get article list failed, err:%v\n", err)
		return
	}
	if len(articleInfoList) == 0 {
		return
	}

	//2.从数据库中，获取文章对应的分类列表
	categoryIds := getCategoryIds(articleInfoList)
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		fmt.Printf("2 get catefory list failed, err:%v\n", err)
		return
	}

	//聚合数据
	for _, article := range articleInfoList {
		fmt.Printf("content:%s\n", article.Summary)
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		categoryId := article.CategoryId
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}

func GetArticleDetail(articleId int) (articleDetail *model.ArticleDetail, err error) {
	//1.获取文章的信息
	articleDetail, err = db.GetArticleDetail(articleId)
	if err != nil {
		return
	}
	// 2.获取文章对应分类的信息
	category, err := db.GetCategoryById(articleDetail.ArticleInfo.CategoryId)
	if err != nil {
		return
	}
	articleDetail.Category = *category
	return
}

func InsertArticle(author, title, content string, categoryId int64) (err error) {
	articleDetail := model.ArticleDetail{}
	articleDetail.Title = title
	articleDetail.Content = content
	articleDetail.ArticleInfo.CategoryId = int(categoryId)

	contetUtf8 := []rune(content)
	minLength := int(math.Min(float64(len(contetUtf8)), 128.0))
	articleDetail.Summary = string([]rune(content)[:minLength])
	id, err := db.InsertArticle(&articleDetail)
	fmt.Printf("insert article succ, id:%d, err:%v\n", id, err)
	return
}

func GetRelativeArticle(articleId int64) (articleList []*model.RelativeArticle, err error) {
	articleList, err = db.GetRelativeArticle(articleId)
	return
}
