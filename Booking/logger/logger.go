package logger

import "github.com/gin-gonic/gin"

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
