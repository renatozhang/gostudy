package model

type RelativeArticle struct {
	ArticleId int    `db:"id"`
	Title     string `db:"title"`
}
