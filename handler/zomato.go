package handler

import (
	"github.com/gin-gonic/gin"
	gerr "github.com/jianhan/gogin/error"
	"github.com/jianhan/gogin/zomato"
	"github.com/leebenson/conform"
	"net/http"
)

func zomatoCategories(c *gin.Context) {
	categories, err := zomato.NewCommonAPI().Categories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &gerr.APIError{Details: err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func zomatoCities(c *gin.Context) {
	// get request
	var req zomato.CitiesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, &gerr.APIError{Details: err.Error()})
		return
	}
	conform.Strings(&req)

	cities, err := zomato.NewCommonAPI().Cities(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &gerr.APIError{Details: err.Error()})
		return
	}

	c.JSON(http.StatusOK, cities)
}

func zomatoCollections(c *gin.Context) {

}
