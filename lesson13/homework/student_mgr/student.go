package main

type Student struct {
	Username string
	Sex      int
	Grade    string
	Score    float32
}

func NewStudent(username string, sex int, grade string, score float32) (stu *Student) {
	stu = &Student{
		Username: username,
		Sex:      sex,
		Grade:    grade,
		Score:    score,
	}
	return
}
