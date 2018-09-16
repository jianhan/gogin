package handler

import (
	"github.com/gin-gonic/gin"
	gerr "github.com/jianhan/gogin/error"
	"github.com/jianhan/gogin/zomato"
	"net/http"
)

func zomatoCategories(c *gin.Context) {
	categories, err := zomato.NewCommonAPI().Categories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &gerr.APIError{Status: http.StatusInternalServerError, Details: err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func zomatoCities(c *gin.Context) {

}

func zomatoCollections(c *gin.Context) {

}
