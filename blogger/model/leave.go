package model

import "time"

type Leave struct {
	Id         int       `db:"id"`
	Username   string    `db:"username"`
	Content    string    `db:"content"`
	Createtime time.Time `db:"create_time"`
	Emai       string    `db:"email"`
}
