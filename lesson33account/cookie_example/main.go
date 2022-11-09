package main

import (
	"fmt"
	"net/http"
)

func indexHandle(w http.ResponseWriter, r *http.Request) {
	/*
		cookies := r.Cookies()
		for index, cookie := range cookies {
			fmt.Printf("index:%d cookies:%#v\n", index, cookie)
		}
	*/

	c, err := r.Cookie("sessionid")
	fmt.Printf("cookie:%#v, err:%#v\n", c, err)
	/*
		cookie := &http.Cookie{
			Name:   "sessionid",
			Value:  "lkjsfdsfsdfsfsfsfs",
			MaxAge: 3600,
			Domain: "localhost",
			Path:   "/",
		}
		http.SetCookie(w, cookie)
	*/
	w.Write([]byte("hello"))
}

func main() {
	http.HandleFunc("/", indexHandle)
	http.ListenAndServe(":9090", nil)
}
