package storage

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic/v7"

	"github.com/feamo/esrest/models"
)

type Engine struct {
	client *elastic.Client
}

func New() (*Engine, error) {
	engine, err := connect()
	if err != nil {
		return nil, err
	}

	return engine, nil
}

func connect() (*Engine, error) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetBasicAuth("", ""),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	engine := &Engine{
		client: client,
	}

	return engine, err

}

func (e *Engine) UserInsert(user *models.User) error {
	_, err := e.client.Index().
		Index("users").
		BodyJson(user).
		Id(user.Email).
		Do(context.TODO())
	if err != nil {
		panic(err)
	}

	return err
}

func (e *Engine) UserGetByID(id string) (*models.User, error) {
	res, err := e.client.Get().Index("users").Id(id).Do(context.TODO())
	if err != nil {
		return nil, err
	}

	user := &models.User{}
	err = json.Unmarshal(res.Source, &user)
	if err != nil {
		panic(err)
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (e *Engine) SearchByQuery(query string) (*models.User, error) {

	esQuery := elastic.NewMultiMatchQuery(query)

	res, err := e.client.Search().
		Index("users").
		Query(esQuery).
		Do(context.TODO())

	user := &models.User{}
	for _, hit := range res.Hits.Hits {
		err = json.Unmarshal(hit.Source, &user)
		if err != nil {
			panic(err)
		}
	}

	if err != nil {
		return nil, err
	}

	return user, nil

}
