package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sliceMap() {
	rand.Seed(time.Now().UnixNano())
	var s []map[string]int
	s = make([]map[string]int, 5, 10)
	for index, value := range s {
		fmt.Printf("slice[%d]=%v\n", index, value)
	}
	s[0] = make(map[string]int)
	s[0]["stu01"] = 100
	s[0]["stu02"] = 200
	s[0]["stu03"] = 300

	for index, val := range s {
		fmt.Printf("slice[%d]=%#v\n", index, val)
	}
}

func mapSlice() {
	rand.Seed(time.Now().UnixNano())
	var s map[string][]int
	s = make(map[string][]int, 16)
	key := "stu01"
	value, ok := s[key]
	if !ok {
		s[key] = make([]int, 0, 16)
		value = s[key]
	}
	value = append(value, 100)
	value = append(value, 200)
	value = append(value, 300)
	s[key] = value
	fmt.Printf("map:%v\n", s)
}

func main() {
	// sliceMap()
	mapSlice()
}
