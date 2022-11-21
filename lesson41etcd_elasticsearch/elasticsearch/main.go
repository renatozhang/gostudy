package main

import (
	"context"
	"fmt"

	elastic "github.com/olivere/elastic/v7"
)

type Twitter struct {
	User    string
	Message string
}

func main() {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://localhost:9200"))
	if err != nil {
		fmt.Println("connect es error", err)
		return
	}
	fmt.Println("connect es success")

	twitter := Twitter{User: "olivere", Message: "Take Five"}
	put1, err := client.Index().Index("twitter").Id("1").BodyJson(twitter).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed twitter: %s to index:%s, type:%s\n", put1.Id, put1.Index, put1.Type)
}
