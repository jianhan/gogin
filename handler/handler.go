package handler

import "github.com/gin-gonic/gin"

func APIHandlers(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		google := v1.Group("/google")
		{
			google.GET("nearby-search", googleNearbySearch)
		}
	}
}
