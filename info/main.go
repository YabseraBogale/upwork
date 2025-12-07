package main

import (
	"log"
	"net/http"
)

func main() {
	url := "https://effoysira.com/page"

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	print(resp.StatusCode)
}
