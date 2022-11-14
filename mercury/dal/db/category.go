package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/renatozhang/gostudy/logger"
	"github.com/renatozhang/gostudy/mercury/common"
)

func GetCategoryList() (categoryList []*common.Category, err error) {
	sqlstr := "select category_id,category_name from category"
	err = DB.Select(&categoryList, sqlstr)
	if err == sql.ErrNoRows {
		return
	}
	if err != nil {
		return
	}
	return
}

func MGetCategory(categoryIds []int64) (categoryMap map[int64]*common.Category, err error) {
	sqlstr := "select category_id,category_name from category where id in (?)"
	var interfaceSlice []interface{}
	for _, c := range categoryIds {
		interfaceSlice = append(interfaceSlice, c)
	}
	insqlStr, params, err := sqlx.In(sqlstr, interfaceSlice...)
	if err != nil {
		logger.Error("sqlx.in failed, sqlstr:%v, err:%v", sqlstr, err)
		return
	}
	categoryMap = make(map[int64]*common.Category, len(categoryIds))
	var categoryList []*common.Category
	err = DB.Select(&categoryList, insqlStr, params...)
	if err == sql.ErrNoRows {
		err = nil
		return
	}
	if err != nil {
		logger.Error("MGetCategory failed, sqlstr:%v, category_ids:%v, err:%v", sqlstr, categoryIds, err)
		return
	}

	for _, v := range categoryList {
		categoryMap[v.CategoryId] = v

	}
	return
}
