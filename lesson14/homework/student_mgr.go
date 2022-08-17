package main

import "fmt"

type StudentMgr struct {
	allStudent []*Student
}

func (p *StudentMgr) AddStudent(stu *Student) (err error) {
	for index, v := range p.allStudent {
		if v.Username == stu.Username {
			fmt.Println("student %s sucess update\n\n", stu.Username)
			p.allStudent[index] = stu
			return
		}
	}
	p.allStudent = append(p.allStudent, stu)
	fmt.Printf("student %s success insert\n\n", stu.Username)
	return
}

func (p *StudentMgr) ModifyStudent(stu *Student) (err error) {
	for index, v := range p.allStudent {
		if v.Username == stu.Username {
			p.allStudent[index] = stu
			fmt.Printf("student %s success update\n\n", stu.Username)
			return
		}
	}
	fmt.Printf("student %s is not found\n", stu.Username)
	err = fmt.Errorf("student %s is not exists", stu.Username)
	return
}

func (p *StudentMgr) ShowAllStudent() {
	for _, v := range p.allStudent {
		fmt.Printf("student:%s info:%#v\n", v.Username, v)
	}
	fmt.Println()
}
