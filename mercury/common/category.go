package common

type Category struct {
	CategoryId   int    `json:"id" db:"category_id"`
	CategoryName string `json:"name" db:"category_name"`
}
