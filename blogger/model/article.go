package model

import (
	"time"
)

type ArticleInfo struct {
	Id           int       `db:"id"`
	CategoryId   int       `db:"category_id"`
	Summary      string    `db:"summary"`
	Title        string    `db:"title"`
	ViewCount    int       `db:"view_count"`
	CreateTime   time.Time `db:"create_time"`
	CommentCount int       `db:"comment_count"`
	UserName     string    `db:"username"`
}

type ArticleDetail struct {
	ArticleInfo
	Content string `db:"content"`
	Category
}

type ArticleRecord struct {
	ArticleInfo
	Category
}
