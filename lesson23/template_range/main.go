package main

import (
	"fmt"
	"math/rand"
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
		fmt.Fprintf(w, "load inde.html failed, err:%v\n", err)
		return
	}
	var userlist []*UserInfo
	for i := 0; i < 30; i++ {
		user := UserInfo{
			Name: fmt.Sprintf("marry%d", rand.Intn(10000)),
			Sex:  "男",
			Age:  rand.Intn(100),
			Address: Address{
				City:     "北京",
				Province: "北京市",
			},
		}
		userlist = append(userlist, &user)
	}

	err = t.Execute(w, userlist)
	if err != nil {
		fmt.Printf("execute template failed, err:%v\n", err)
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
