package main

import (
	"fmt"
	"os"
)

var (
	AllStudents []*Student
)

func showMenu() {
	fmt.Println("1. add student")
	fmt.Println("2. modify student")
	fmt.Println("3. show all student")
	fmt.Println("4. exited")
}
func InputStudent() (stu *Student) {
	var (
		username string
		sex      int
		grade    string
		score    float32
	)
	fmt.Println("please input username")
	fmt.Scanf("%s\n", &username)
	fmt.Println("please input sex")
	fmt.Scanf("%d\n", &sex)
	fmt.Println("please input grade")
	fmt.Scanf("%s\n", &grade)
	fmt.Println("please input score")
	fmt.Scanf("%f\n", &score)
	stu = NewStudent(username, sex, grade, score)
	return
}

func AddStudent() {
	stu := InputStudent()
	for index, v := range AllStudents {
		if v.Username == stu.Username {
			fmt.Printf("user %s sucess update\n\n", stu.Username)
			AllStudents[index] = stu
			return
		}
	}
	AllStudents = append(AllStudents, stu)
	fmt.Printf("user %s insert\n\n", stu.Username)
}

func ModidyStudent() {
	stu := InputStudent()
	for index, v := range AllStudents {
		if v.Username == stu.Username {
			fmt.Printf("user %s sucess update\n\n", stu.Username)
			AllStudents[index] = stu
			return
		}
	}
	fmt.Printf("user %s not found\n\n", stu.Username)
}

func ShowAllStudent() {
	for _, v := range AllStudents {
		fmt.Printf("user:%s info:%#v\n", v.Username, v)
	}
	fmt.Println()
}

func main() {
	for {
		showMenu()
		var sel int
		fmt.Scanf("%d\n", &sel)
		switch sel {
		case 1:
			AddStudent()
		case 2:
			ModidyStudent()
		case 3:
			ShowAllStudent()
		case 4:
			os.Exit(0)
		}
	}
}
