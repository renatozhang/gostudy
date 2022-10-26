package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}

func writeJson(filename string) (err error) {
	var persons []*Person
	for i := 0; i < 10; i++ {
		p := &Person{
			Name: fmt.Sprintf("name%d", i),
			Age:  rand.Intn(100),
			Sex:  "man",
		}
		persons = append(persons, p)
	}
	data, err := json.Marshal(persons)
	if err != nil {
		fmt.Printf("=marshal failed, err:%v\n", err)
		return
	}

	err = ioutil.WriteFile(filename, data, 0755)
	if err != nil {
		fmt.Printf("write file failed, err:%v\n", err)
		return
	}
	return
}

func readJson(filename string) (err error) {
	var person []*Person
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("read file failed, err:%v\n", err)
		return
	}

	err = json.Unmarshal(data, &person)
	if err != nil {
		return
	}

	for _, v := range person {
		fmt.Printf("%#v\n", v)
	}
	// fmt.Printf("%#v\n", person)
	return
}

func main() {
	filename := "/tmp/person.json"
	err := writeJson(filename)
	if err != nil {
		fmt.Printf("write json file failed, err:%v\n", err)
		return
	}
	err = readJson(filename)
	if err != nil {
		fmt.Printf("read json file failed, err:%v\n", err)
		return
	}
}
