package logic

import (
	"fmt"

	"github.com/renatozhang/gostudy/blogger/dal/db"
	"github.com/renatozhang/gostudy/blogger/model"
)

func GetLeaveList(pageNum, pageSize int) (leaveList []*model.Leave, err error) {
	leaveList, err = db.GetLeaveList(pageNum, pageSize)
	if err != nil {
		fmt.Printf("get leave list failed, err:%v\n", err)
		return
	}
	return
}

func InsertLeave(username, email, content string) (err error) {
	var leave model.Leave
	leave.Username = username
	leave.Email = email
	leave.Content = content
	err = db.InsertLeave(&leave)
	if err != nil {
		fmt.Printf("insert leave failed, err:%v\n", err)
		return
	}
	return
}
