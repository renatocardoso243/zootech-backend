package middleware

import (
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
        ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

        if ctx.Request.Method == "OPTIONS" {
            ctx.AbortWithStatus(204)
            return
        }

        ctx.Next()
    }
}
