package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id   string
	Name string
	Sex  string
}

type Class struct {
	Name     string
	Count    int
	Students []*Student
}

var rowJson = `
{"Name":"101","Count":0,"Students":[{"Id":"0","Name":"stu0","Sex":"main"},
{"Id":"1","Name":"stu1","Sex":"main"},{"Id":"2","Name":"stu2","Sex":"main"},
{"Id":"3","Name":"stu3","Sex":"main"},{"Id":"4","Name":"stu4","Sex":"main"},
{"Id":"5","Name":"stu5","Sex":"main"},{"Id":"6","Name":"stu6","Sex":"main"},
{"Id":"7","Name":"stu7","Sex":"main"},{"Id":"8","Name":"stu8","Sex":"main"},
{"Id":"9","Name":"stu9","Sex":"main"}]}
`

func main() {
	c := &Class{
		Name:  "101",
		Count: 0,
	}

	for i := 0; i < 10; i++ {
		stu := &Student{
			Name: fmt.Sprintf("stu%d", i),
			Sex:  "man",
			Id:   fmt.Sprintf("%d", i),
		}
		c.Students = append(c.Students, stu)
	}

	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", string(data))

	// json反序列化
	var c1 *Class = &Class{}
	err = json.Unmarshal([]byte(rowJson), c1)
	if err != nil {
		fmt.Println("Unmarshal failed")
	}
	fmt.Printf("c1:%#v\n", c1)

	for _, v := range c1.Students {
		fmt.Printf("student:%#v\n", v)
	}
}
