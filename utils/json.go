package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	InternalServerError  = "Internal server error."
	MethodNotImplemented = errors.New("method not implemented")
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func WriteJson(c *gin.Context, data interface{}) {
	res := Response{
		Message: "Success",
		Data:    data,
	}
	//
	c.JSON(200, res)
}

func WriteError(c *gin.Context, status int, err string) {
	res := Response{
		Message: err,
	}
	//
	c.JSON(status, res)
}

func WriteInternalServerError(c *gin.Context, err ...string) {
	if len(err) == 0 {
		WriteError(c, 500, InternalServerError)
		return
	}
	WriteError(c, 500, err[0])

}
