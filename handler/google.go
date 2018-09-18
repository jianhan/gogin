package handler

import (
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/contrib/cache"
	"github.com/gin-gonic/gin"
	gerr "github.com/jianhan/gogin/error"
	"github.com/jianhan/gogin/google"
	"github.com/leebenson/conform"
	"net/http"
	"time"
)

type googleAPIHandlerRegister struct {
	nearbySearch google.NearbySearch
}

func (g *googleAPIHandlerRegister) Register(r *gin.RouterGroup) {
	store := persistence.NewInMemoryStore(time.Duration(5) * time.Minute)
	googleNearbySearch := r.Group("/google")
	{
		googleNearbySearch.GET("nearby-search", cache.CachePage(store, time.Duration(2)*time.Hour, googleNearbySearch))
	}
}

func (g *googleAPIHandlerRegister) NearbySearch(c *gin.Context) {
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

func NewGoogleAPIHandlerRegister(nearbySearch google.NearbySearch) Register {
	return &googleAPIHandlerRegister{nearbySearch: nearbySearch}
}
