package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponsError struct {
	Status  string `json:"status,omitempty"`
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type Response struct {
	Data interface{} `json:"data,omitempty"`
}

func RequestError(ct *gin.Context, msg string, code int) {
	out := ResponsError{
		Status:  http.StatusText(code),
		Code:    code,
		Message: msg,
	}
	ct.JSON(code, out)

}

func Requestok(ct *gin.Context, code int, data interface{}) {
	ct.JSON(code, data)
}
