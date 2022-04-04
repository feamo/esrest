package api

import (
	"fmt"
	"github.com/feamo/esrest/models"
	"net/http"
	// "strconv"

	"github.com/gin-gonic/gin"

	// "github.com/feamo/esrest/models"
	"github.com/feamo/esrest/storage"
)

func New(es *storage.Engine) *User {
	return &User{es: es}
}

type User struct {
	es *storage.Engine
}

func (r *User) UserCreate(c *gin.Context) {
	var requestBody models.User
	err := c.BindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	err = r.es.UserInsert(&requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (r *User) UserGetId(c *gin.Context) {
	var requestBody models.User
	err := c.BindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Println(requestBody.Email)
	user, err := r.es.UserGetByID(requestBody.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (r *User) SearchQuery(c *gin.Context) {

	query := c.Query("name")

	user, err := r.es.SearchByQuery(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
