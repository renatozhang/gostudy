package model

import "time"

type Comment struct {
	Id         int       `db:"id"`
	Content    string    `db:"content"`
	Username   string    `db:"username"`
	CreateTime time.Time `db:"create_time"`
	Status     int       `db:"status"`
	ArticleId  int       `db:"article_id"`
}
