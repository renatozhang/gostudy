package db

import "testing"

func init() {
	dsn := "root:123456@tcp(localhost:3306)/blog?parseTime=true"
	err := Init(dsn)
	if err != nil {
		panic(err)
	}
}

func TestGetCategoryList(t *testing.T) {
	var categoryIds []int64
	categoryIds = append(categoryIds, 1, 2, 3)
	categoryList, err := GetCategoryList(categoryIds)
	if err != nil {
		t.Errorf("get category list failed, err:%v\n", err)
		return
	}

	if len(categoryList) != len(categoryIds) {
		t.Errorf("get category list failed, len of categorylist:%d ids len:%d\n", len(categoryList), len(categoryIds))
	}
	for _, v := range categoryList {
		t.Logf("id: %d category:%#v\n", v.CategoryId, v)
	}
}

func TestGetCategoryById(t *testing.T) {
	category, err := GetCategoryById(1)
	if err != nil {
		t.Errorf("get category failed, err:%n\n", err)
		return
	}
	t.Logf("get category succ, article:%#v\n", category)

}

func TestGetAllCategoryList(t *testing.T) {
	categoryList, err := GetALLCategoryList()
	if err != nil {
		t.Errorf("get category list failed, err:%v\n", err)
		return
	}
	t.Logf("category:%#v\n", categoryList)
}
