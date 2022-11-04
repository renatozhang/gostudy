package logic

import (
	"fmt"

	"github.com/renatozhang/gostudy/blogger/dal/db"
	"github.com/renatozhang/gostudy/blogger/model"
)

func GetALLCategoryList() (categoryList []*model.Category, err error) {
	//1.从数据库中，获取文章分类列表
	categoryList, err = db.GetALLCategoryList()
	if err != nil {
		fmt.Printf("get category list failed, err:%v\n", err)
		return
	}
	return
}
