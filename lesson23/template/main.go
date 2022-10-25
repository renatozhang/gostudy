package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

type UserInfo struct {
	Name string
	Sex  string
	Age  int
}

func login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./login.html")
	if err != nil {
		fmt.Fprintf(w, "load login.html failed, err:%v\n", err)
		return
	}
	/*
		user := UserInfo{
			Name: "marry",
			Sex:  "男",
			Age:  18,
		}
	*/
	m := make(map[string]string)
	m["Username"] = "Marry"
	m["Sex"] = "男"
	m["Age"] = "18"
	// t.Execute(w, "marry")
	// t.Execute(w, user)
	t.Execute(w, m)
	t.Execute(os.Stdout, m)
}

func main() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("listen server failed, err:%v\n", err)
		return
	}
}
