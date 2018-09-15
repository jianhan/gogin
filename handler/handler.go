package handler

import "github.com/gin-gonic/gin"

func NewGinEngine(opts ...func(*gin.Engine)) *gin.Engine {
	r := gin.Default()
	for _, opt := range opts {
		opt(r)
	}

	return r
}

func APIHandlers(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		google := v1.Group("/google")
		{
			google.GET("nearby-search", googleNearbySearch)
		}
	}
}
