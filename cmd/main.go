package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/feamo/esrest/api"
	"github.com/feamo/esrest/storage"
)

func main() {
	es, err := storage.New()
	if err != nil {
		log.Fatalln(err)
	}

	user := api.New(es)

	router := gin.Default()
	router.POST("/users", user.UserCreate)
	router.POST("/users/byId", user.UserGetId)
	router.GET("/users/search", user.SearchQuery)

	router.Run()
}
