package handler

import (
	"github.com/gin-gonic/gin"
	gerr "github.com/jianhan/gogin/error"
	"github.com/jianhan/gogin/google"
	"github.com/leebenson/conform"
	"net/http"
)

type GoogleAPIHandler struct {
	nearbySearch google.NearbySearch
}

func (g *GoogleAPIHandler) NearbySearch(c *gin.Context) {
	// generate request
	var req google.NearbySearchRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, &gerr.APIError{Details: err.Error()})
		return
	}
	conform.Strings(&req)

	res, status, err := g.nearbySearch.Search(c, &req)
	if err != nil {
		c.JSON(status, err)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

func googleNearbySearch(c *gin.Context) {

	c.JSON(200, response.Results)
}
