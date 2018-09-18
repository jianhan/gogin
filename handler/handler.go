package handler

import (
	"github.com/gin-gonic/gin"
)

type Register interface {
	Register(r *gin.RouterGroup)
}

type APIHandlersRegister struct {
	registers []Register
}

func (a *APIHandlersRegister) Register(r *gin.Engine) {
	// api v1 version
	v1 := r.Group("/api/v1")
	{
		for _, v := range a.registers {
			v.Register(v1)
		}
		//// google nearby search
		//google := v1.Group("/google")
		//{
		//	google.GET("nearby-search", cache.CachePage(store, time.Duration(2)*time.Hour, googleNearbySearch))
		//}
		//
		//// zomato search
		//zomato := v1.Group("/zomato")
		//{
		//	zomato.GET("categories", cache.CachePage(store, time.Duration(96)*time.Hour, zomatoCategories))
		//	zomato.GET("cities", cache.CachePage(store, time.Duration(96)*time.Hour, zomatoCities))
		//	zomato.GET("collections", cache.CachePage(store, time.Duration(24)*time.Hour, zomatoCollections))
		//	zomato.GET("establishments", cache.CachePage(store, time.Duration(24)*time.Hour, zomatoEstablishments))
		//	zomato.GET("cuisines", cache.CachePage(store, time.Duration(48)*time.Hour, zomatoCuisines))
		//	zomato.GET("search-restaurants", cache.CachePage(store, time.Duration(5)*time.Minute, zomatoSearchRestaurants))
		//}
	}
}

func NewAPIHandlersRegister(registers ...Register) *APIHandlersRegister {
	return &APIHandlersRegister{registers: registers}
}

//// APIHandlers receive an gin engine and register API routes.
//func APIHandlers(r *gin.Engine) {
//	// store for cache purpose
//	store := persistence.NewInMemoryStore(time.Duration(5) * time.Minute)
//
//}
