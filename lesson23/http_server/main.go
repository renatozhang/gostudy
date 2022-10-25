package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数，默认是不会解析的
	fmt.Fprintf(w, "%v\n", r.Form)
	fmt.Fprintf(w, "path:%v\n", r.URL.Path)
	fmt.Fprintf(w, "schema:%v\n", r.URL.Scheme)
	fmt.Fprintf(w, "hello world\n")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, errr:%v\n", err)
	}
}
