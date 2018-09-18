package handler

import (
	"github.com/gin-gonic/gin"
	gerr "github.com/jianhan/gogin/error"
	"github.com/leebenson/conform"
	"net/http"
)

type googleNearbySearchRequest struct {
	Lat       float64 `form:"lat" json:"lat" validate:"required"`
	Lng       float64 `form:"lng" json:"lon" validate:"required"`
	Radius    uint    `form:"radius" json:"radius" validate:"required,gte=100,lte=10000"`
	Keyword   string  `conform:"trim" form:"keyword" json:"keyword"`
	MinPrice  uint    `form:"min_price" json:"min_price" validate:"gte=0,lte=4"`
	MaxPrice  uint    `form:"max_price" json:"max_price" validate:"gte=0,lte=4"`
	OpenNow   bool    `form:"open_now" json:"open_now,omitempty"`
	RankBy    string  `conform:"trim" form:"rankby" json:"rankby" validate:"oneof=prominence distance"`
	PageToken string  `conform:"trim" form:"page_token" json:"page_token"`
}

func googleNearbySearch(c *gin.Context) {
	var req googleNearbySearchRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, &gerr.APIError{Details: err.Error()})
		return
	}
	conform.Strings(&req)

	c.JSON(200, response.Results)
}
