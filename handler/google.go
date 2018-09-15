package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jianhan/gogin/error"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type googleNearbySearchRequest struct {
	Lat       float64 `form:"lat" json:"lat" validate:"required,latitude"`
	Lon       float64 `form:"lon" json:"lon" validate:"required,longitude"`
	Radius    uint    `form:"radius" json:"radius" validate:"required,gte=100,lte=10000"`
	Keyword   string  `form:"keyword" json:"keyword"`
	MinPrice  uint    `form:"min_price" json:"min_price" validate:"gte=0,lte=4"`
	MaxPrice  uint    `form:"max_price" json:"max_price" validate:"gte=0,lte=4"`
	RankBy    string  `form:"rankby" json:"rankby" validate:"oneof=prominence distance"`
	PageToken string  `form:"page_token" json:"page_token"`
}

func googleNearbySearch(c *gin.Context) {
	var req googleNearbySearchRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	// validation
	if vErr := validator.New().Struct(req); vErr != nil {
		apiError := &error.APIError{Status: http.StatusBadRequest, Details: "validation error"}
		if _, ok := vErr.(*validator.InvalidValidationError); ok {
			c.JSON(http.StatusBadRequest, apiError)
			return
		}

		for _, err := range vErr.(validator.ValidationErrors) {
			apiError.AddData(err.Field(), fmt.Sprintf("invalid input for %s", err.Field()))
		}

		c.JSON(http.StatusBadRequest, apiError)
		return
	}

	// make API request
	c.String(200, "pong")
}
