package main

import (
	"log"

	"github.com/olivere/elastic/v7"
)

func main() {
	es, err := elastic.NewClient()
	if err != nil {
		log.Fatalln(err)
	}
	defer es.Stop()
}
