package main

import (
	"fmt"
	"os"
)

var (
	studentMgr = &StudentMgr{}
)

func showMenu() {
	fmt.Println("1. add student")
	fmt.Println("2. modify student")
	fmt.Println("3. show all student")
	fmt.Printf("4. exited\n\n")
}

func InputStudent() *Student {
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

	stu := NewStudent(username, sex, grade, score)
	return stu
}

func main() {
	for {
		showMenu()
		var sel int
		fmt.Scanf("%d\n", &sel)
		switch sel {
		case 1:
			stu := InputStudent()
			studentMgr.AddStudent(stu)
		case 2:
			stu := InputStudent()
			studentMgr.ModifyStudent(stu)
		case 3:
			studentMgr.ShowAllStudent()
		case 4:
			os.Exit(0)
		}
	}
}
