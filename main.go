package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jianhan/gogin/config"
	"github.com/jianhan/gogin/google"
	"github.com/jianhan/gogin/handler"
	"net/http"
	"time"
)

func main() {
	r, err := getGinEngine()
	if err != nil {
		panic(err)
	}

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

func getGinEngine() (*gin.Engine, error) {
	r := gin.Default()

	// get google map client
	googleMapClient, err := google.NewGoogleMapClient()
	if err != nil {
		return nil, err
	}

	// register google map handler
	handler.NewAPIHandlersRegister(
		handler.NewGoogleAPIHandlerRegister(google.NewNearbySearch(googleMapClient)),
	).Register(r)

	return r, nil
}
