package web

import "github.com/gin-gonic/gin"

type ResponsError struct {
	Status  int    `json:"status,omitempty"`
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type Response struct {
	Data interface{} `json:"data,omitempty"`
}

func RequestError(ct *gin.Context, msg string, code int) {

}

func Requestok(ct *gin.Context, code int, data interface{}) {
	ct.JSON(code, data)
}
