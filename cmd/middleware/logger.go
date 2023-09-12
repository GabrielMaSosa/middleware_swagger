package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		timerq := time.Now()

		ctx.Next()

		durati := time.Since(timerq).Milliseconds()
		mehod := ctx.Request.Method
		url := ctx.Request.URL.String()
		code := ctx.Writer.Status()

		log.Printf("Method: %s| URL: %s | duration: %d miliseconds |Code:%d", mehod, url, durati, code)
	}
}
