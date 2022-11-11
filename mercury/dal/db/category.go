package db

import (
	"database/sql"

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
