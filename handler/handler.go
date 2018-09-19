package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	gerr "github.com/jianhan/gogin/error"
	"github.com/leebenson/conform"
	"gopkg.in/go-playground/validator.v9"
)

func validateRequest(c *gin.Context, r interface{}) error {
	// generate request
	if err := c.BindQuery(r); err != nil {
		return &gerr.APIError{Details: "unable to process request, invalid input"}
	}
	conform.Strings(&r)

	if vErr := validator.New().Struct(r); vErr != nil {
		apiError := &gerr.APIError{Details: "validation error"}
		if _, ok := vErr.(*validator.InvalidValidationError); ok {
			return apiError
		}

		for _, err := range vErr.(validator.ValidationErrors) {
			apiError.AddData(err.Field(), fmt.Sprintf("invalid input for %s", err.Field()))
		}

		return apiError
	}

	return nil
}

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
