package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println(ctx.Request.RequestURI)
		ctx.Redirect(http.StatusMovedPermanently, "http://localhost:3000")
		return
	}
}