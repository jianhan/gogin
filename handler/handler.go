package handler

import "github.com/gin-gonic/gin"

func NewGinEngine(opts ...func(*gin.Engine)) *gin.Engine {
	r := gin.Default()
	for _, opt := range opts {
		opt(r)
	}

	return r
}
