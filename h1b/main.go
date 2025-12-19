package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("/home/yabsera/Downloads/Alabama_H1B_Jobs.csv")
	if err != nil {
		log.Println(err)
	}

	for _, i := range file {
		fmt.Println(i)
	}
}
