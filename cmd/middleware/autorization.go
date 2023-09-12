package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AutorizationMiddleware() gin.HandlerFunc {
	token := os.Getenv("TOKEN")
	return func(ctx *gin.Context) {
		//before request
		if ctx.GetHeader("token") != token {
			code := http.StatusUnauthorized
			body := gin.H{"message": "Unauthorize"}

			ctx.JSON(code, body)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
