package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jianhan/gogin/config"
	"github.com/jianhan/gogin/google"
	"github.com/jianhan/gogin/handler"
	"googlemaps.github.io/maps"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	googleMapClient, err := maps.NewClient(maps.WithAPIKey(config.GetEnvs().GoogleMapAPIKey))
	if err != nil {
		panic(err)
		log.Fatalf("fatal error: %s", err)
	}
	handler.NewAPIHandlersRegister(
		handler.NewGoogleAPIHandlerRegister(google.NewNearbySearch(googleMapClient)),
	).Register(r)

	// config server
	s := &http.Server{
		Addr:              config.GetEnvs().Addr,
		Handler:           r,
		ReadTimeout:       time.Duration(config.GetEnvs().ReadTimeout) * time.Second,
		WriteTimeout:      time.Duration(config.GetEnvs().WriteTimeout) * time.Second,
		ReadHeaderTimeout: time.Duration(config.GetEnvs().ReadHeaderTimeout) * time.Second,
		IdleTimeout:       time.Duration(config.GetEnvs().IdleTimeout) * time.Second,
		MaxHeaderBytes:    config.GetEnvs().MaxHeaderBytes,
	}
	s.ListenAndServe()
}

// newGinEngine return a new gin engine with optional functions.
func newGinEngine(opts ...func(*gin.Engine)) *gin.Engine {
	r := gin.Default()
	for _, opt := range opts {
		opt(r)
	}

	return r
}
