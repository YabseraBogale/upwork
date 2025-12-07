package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	url := "https://effoysira.com/page/1"

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(resp.StatusCode)
}
