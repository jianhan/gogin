package main

import (
	"github.com/jianhan/gogin/config"
	"github.com/jianhan/gogin/handler"
	"net/http"
	"time"
)

func main() {
	s := &http.Server{
		Addr:              config.GetEnvs().Addr,
		Handler:           handler.NewGinEngine(handler.APIHandlers),
		ReadTimeout:       time.Duration(config.GetEnvs().ReadTimeout) * time.Second,
		WriteTimeout:      time.Duration(config.GetEnvs().WriteTimeout) * time.Second,
		ReadHeaderTimeout: time.Duration(config.GetEnvs().ReadHeaderTimeout) * time.Second,
		IdleTimeout:       time.Duration(config.GetEnvs().IdleTimeout) * time.Second,
		MaxHeaderBytes:    config.GetEnvs().MaxHeaderBytes,
	}
	s.ListenAndServe()
}
