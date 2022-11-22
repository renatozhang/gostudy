package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

type ResPack struct {
	r   *http.Response
	err error
}

func work(ctx context.Context) {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	defer wg.Done()
	c := make(chan ResPack, 1)
	req, _ := http.NewRequest("GET", "http://localhost:8000", nil)
	go func() {
		resp, err := client.Do(req)
		pack := ResPack{r: resp, err: err}
		c <- pack
	}()
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-c
		fmt.Println("timeout")
	case res := <-c:
		if res.err != nil {
			fmt.Println(res.err)
			return
		}
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("server Response: %s", out)
	}
}

// func work(ctx context.Context) {
// 	client := &http.Client{}
// 	defer wg.Done()
// 	// c := make(chan ResPack, 1)
// 	req, _ := http.NewRequestWithContext(ctx, "GET", "http://localhost:8000", nil)
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return
// 	}
// 	out, _ := ioutil.ReadAll(resp.Body)
// 	fmt.Printf("server Response: %s", out)
// 	defer resp.Body.Close()

// }

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	wg.Add(1)
	go work(ctx)
	wg.Wait()
	fmt.Println("Finished")
}
