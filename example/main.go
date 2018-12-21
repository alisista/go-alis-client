package main

import (
	"fmt"
	"os"
	"time"

	"github.com/alisista/go-alis-client/client"
	"github.com/alisista/go-alis-client/client/operations"
)

func main() {
	p := operations.GetSearchArticlesParams{}
	p.WithQuery("結婚するメリット")
	p.WithTimeout(2 * time.Second)

	resp, err := client.Default.Operations.GetSearchArticles(&p)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resp.Payload[0])
}
