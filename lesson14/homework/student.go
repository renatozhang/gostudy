package main

type Student struct {
	Username string
	Score    float32
	Grade    string
	Sex      int
}

func NewStudent(username string, sex int, grade string, score float32) (stu *Student) {
	stu = &Student{
		Username: username,
		Score:    score,
		Grade:    grade,
		Sex:      sex,
	}
	return
}
