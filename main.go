package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jianhan/gogin/config"
	"github.com/jianhan/gogin/handler"
	"net/http"
	"time"
)

func main() {

	// config server
	s := &http.Server{
		Addr:              config.GetEnvs().Addr,
		Handler:           newGinEngine(handler.APIHandlers),
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
