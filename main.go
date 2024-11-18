package main

import (
	"log"

	"github.com/clipclock08/news-crud/cmd"
)

func main() {
	if err := cmd.Exec(); err != nil {
		log.Fatal(err)
	}
}
