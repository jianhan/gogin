package handler

import (
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"time"
)

func APIHandlers(r *gin.Engine) {
	store := persistence.NewInMemoryStore(time.Second)
	v1 := r.Group("/api/v1")
	{
		google := v1.Group("/google")
		{
			google.GET("nearby-search", cache.CachePage(store, time.Hour, googleNearbySearch))
		}
	}
}
