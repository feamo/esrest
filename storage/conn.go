package storage

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type Engine struct {
	db *elastic.Client
}

func connect() (*Engine, error) {
	const (
		host     = "localhost"
		port     = 9200
		userName = "feamo"
		dbname   = "elasticsearch"
	)

	// connection string
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		userName,
		dbname,
	)
	println(connStr)

	// open database
	db, err := elastic.NewClient()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := db.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

}
