package db

import (
	"fmt"

	"github.com/renatozhang/gostudy/blogger/model"
)

func GetLeaveList(pageNum, pageSize int) (leaveList []*model.Leave, err error) {
	sqlstr := "select id,username,email,content,create_time from `leave` order by id desc limit ?,?"
	err = DB.Select(&leaveList, sqlstr, pageNum, pageSize)
	if err != nil {
		fmt.Printf("exec sql:%s  failed, err:%v\n", sqlstr, err)
		return
	}
	return
}

func InsertLeave(leave *model.Leave) (err error) {
	sqlstr := "insert into `leave` (username, email,content) values(?,?,?)"
	_, err = DB.Exec(sqlstr, leave.Username, leave.Email, leave.Content)
	if err != nil {
		fmt.Printf("exec sql:%s failed,err:%v\n", sqlstr, err)
		return
	}
	return
}
