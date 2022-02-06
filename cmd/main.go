package main

import (
	"log"
)

func main() {
	es, err := storage.New()
	if err != nil {
		log.Fatalln(err)
	}
}
