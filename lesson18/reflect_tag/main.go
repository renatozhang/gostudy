package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"name" db:"name"`
	Sex   int
	Age   int
	Score float32
}

func (s *Student) SetName(name string) {
	s.Name = name
}

func (s *Student) Print() {
	fmt.Printf("通过反射进行调用：%#v\n", s)
}

func main() {
	var s Student
	s.SetName("xxx")
	v := reflect.ValueOf(&s)
	t := v.Type()
	field0 := t.Elem().Field(0).Tag
	fmt.Printf("tag json=%s\n", field0.Get("json"))
	fmt.Printf("tag json=%s\n", field0.Get("db"))
}
