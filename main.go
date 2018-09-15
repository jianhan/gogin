package main

import (
	"github.com/jianhan/gogin/handler"
	"net/http"
	"time"
)

func main() {
	s := &http.Server{
		Addr:           ":8888",
		Handler:        handler.NewGinEngine(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}