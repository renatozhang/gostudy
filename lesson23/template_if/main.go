package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Address struct {
	City     string
	Province string
}

type UserInfo struct {
	Name    string
	Sex     string
	Age     int
	Address Address
}

func login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Fprintf(w, "load index.html failed, err:%v\n", err)
		return
	}
	user := UserInfo{
		Name: "marry",
		Sex:  "男",
		Age:  18,
		Address: Address{
			City:     "北京",
			Province: "北京市",
		},
	}
	/*
		m := make(map[string]interface{})
		m["Name"] = "Marry"
		m["Sex"] = "男"
		m["Age"] = 18

		err = t.Execute(w, m)
	*/
	err = t.Execute(w, user)
	if err != nil {
		fmt.Printf("execute template failes,err:%v\n", err)
	}
}

func main() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("listen server failed, err:%v\n", err)
		return
	}
}
