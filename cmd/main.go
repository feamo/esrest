package main

import (
	"github.com/olivere/elastic/v7"
	"log"
)

func main() {

	es, err := elastic.NewClient()
	if err != nil {
		log.Fatalln(err)
	}
	defer es.Stop()

}
