package handler

import (
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	gerr "github.com/jianhan/gogin/error"
	"github.com/jianhan/gogin/google"
	"github.com/leebenson/conform"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// googleAPIHandlerRegister register google handler func to handle all google related requests.
type googleAPIHandlerRegister struct {
	nearbySearch google.NearbySearch
}

// Register implements Register interface.
func (g *googleAPIHandlerRegister) Register(r *gin.RouterGroup) {
	store := persistence.NewInMemoryStore(time.Duration(5) * time.Minute)
	googleNearbySearchRouter := r.Group("/google")
	{
		googleNearbySearchRouter.GET("nearby-search", cache.CachePage(store, time.Duration(2)*time.Hour, g.NearbySearch))
	}
}

// NearbySearch contains logic for nearby search.
func (g *googleAPIHandlerRegister) NearbySearch(c *gin.Context) {
	// generate request
	var req google.NearbySearchRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, &gerr.APIError{Details: err.Error()})
		return
	}
	conform.Strings(&req)
	logrus.Info(req)

	res, status, err := g.nearbySearch.Search(c, &req)
	if err != nil {
		c.JSON(status, err)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

// NewGoogleAPIHandlerRegister returns a new google nearby search instance.
func NewGoogleAPIHandlerRegister(nearbySearch google.NearbySearch) Register {
	return &googleAPIHandlerRegister{nearbySearch: nearbySearch}
}
