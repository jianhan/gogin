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
	}
}

func NewAPIHandlersRegister(registers ...Register) *APIHandlersRegister {
	return &APIHandlersRegister{registers: registers}
}
