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
	// get request
	var req zomato.CollectionsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, &gerr.APIError{Details: err.Error()})
		return
	}
	conform.Strings(&req)

	collections, err := zomato.NewCommonAPI().Collections(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &gerr.APIError{Details: err.Error()})
		return
	}

	c.JSON(http.StatusOK, collections)
}

func zomatoEstablishments(c *gin.Context) {
	// get request
	var req zomato.EstablishmentsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, &gerr.APIError{Details: err.Error()})
		return
	}
	conform.Strings(&req)

	establishments, err := zomato.NewCommonAPI().Establishments(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &gerr.APIError{Details: err.Error()})
		return
	}

	c.JSON(http.StatusOK, establishments)
}

func zomatoCuisines(c *gin.Context) {
	// get request
	var req zomato.CuisinesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, &gerr.APIError{Details: err.Error()})
		return
	}
	conform.Strings(&req)

	cuisines, err := zomato.NewCommonAPI().Cuisines(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &gerr.APIError{Details: err.Error()})
		return
	}

	c.JSON(http.StatusOK, cuisines)
}

func zomatoSearchRestaurants(c *gin.Context) {
	// get request
	var req zomato.RestaurantsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, &gerr.APIError{Details: err.Error()})
		return
	}
	conform.Strings(&req)

	restaurants, err := zomato.NewRestaurantAPI().SearchRestaurants(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &gerr.APIError{Details: err.Error()})
		return
	}

	c.JSON(http.StatusOK, restaurants)
}
