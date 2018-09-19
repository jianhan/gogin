package handler

import (
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	gerr "github.com/jianhan/gogin/error"
	"github.com/jianhan/gogin/zomato"
	"github.com/leebenson/conform"
	"net/http"
	"time"
)

// zomatoCommonAPIHandlerRegister register zomato common API funcs to handle all common requests.
type zomatoCommonAPIHandlerRegister struct {
	commonAPI zomato.CommonAPI
}

// Register implements Register interface.
func (z *zomatoCommonAPIHandlerRegister) Register(r *gin.RouterGroup) {
	store := persistence.NewInMemoryStore(time.Duration(5) * time.Minute)
	googleNearbySearchRouter := r.Group("/zomato/common")
	{
		googleNearbySearchRouter.GET("categories", cache.CachePage(store, time.Duration(48)*time.Hour, z.Categories))
		googleNearbySearchRouter.GET("cities", cache.CachePage(store, time.Duration(24)*time.Hour, z.Cities))
		googleNearbySearchRouter.GET("collections", cache.CachePage(store, time.Duration(12)*time.Hour, z.Collections))
		googleNearbySearchRouter.GET("establishments", cache.CachePage(store, time.Duration(12)*time.Hour, z.Establishments))
		googleNearbySearchRouter.GET("cuisines", cache.CachePage(store, time.Duration(12)*time.Hour, z.Cuisines))
	}
}

func (g *zomatoCommonAPIHandlerRegister) Categories(c *gin.Context) {
	res, status, err := g.commonAPI.Categories()
	if err != nil {
		c.JSON(status, err)
		return
	}

	c.JSON(status, res)
}

func (g *zomatoCommonAPIHandlerRegister) Cities(c *gin.Context) {
	// generate request
	var req zomato.CitiesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, &gerr.APIError{Details: err.Error()})
		return
	}
	conform.Strings(&req)

	res, status, err := g.commonAPI.Cities(&req)
	if err != nil {
		c.JSON(status, err)
		return
	}

	c.JSON(status, res)
}

func (g *zomatoCommonAPIHandlerRegister) Collections(c *gin.Context) {
	// generate request
	var req zomato.CollectionsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, &gerr.APIError{Details: err.Error()})
		return
	}
	conform.Strings(&req)

	res, status, err := g.commonAPI.Collections(&req)
	if err != nil {
		c.JSON(status, err)
		return
	}

	c.JSON(status, res)
}

func (g *zomatoCommonAPIHandlerRegister) Establishments(c *gin.Context) {
	// generate request
	var req zomato.EstablishmentsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, &gerr.APIError{Details: err.Error()})
		return
	}
	conform.Strings(&req)

	res, status, err := g.commonAPI.Establishments(&req)
	if err != nil {
		c.JSON(status, err)
		return
	}

	c.JSON(status, res)
}

func (g *zomatoCommonAPIHandlerRegister) Cuisines(c *gin.Context) {
	// generate request
	var req zomato.CuisinesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, &gerr.APIError{Details: err.Error()})
		return
	}
	conform.Strings(&req)

	res, status, err := g.commonAPI.Cuisines(&req)
	if err != nil {
		c.JSON(status, err)
		return
	}

	c.JSON(status, res)
}

// NewZomatoCommonAPIHandlerRegister returns a new zomato common API handler
func NewZomatoCommonAPIHandlerRegister(commonAPI zomato.CommonAPI) Register {
	return &zomatoCommonAPIHandlerRegister{commonAPI: commonAPI}
}
