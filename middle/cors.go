package middle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 处理跨域请求
func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		//添加跨域请求头
		ctx.Header("Content-Type", "application/json")
		ctx.Header("Access-Control-Allow-origin", "*")
		ctx.Header("Access-Control-Max-Age", "86400")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTION, PUT, DELETE, UPDATE")
		ctx.Header("Access-Control-Allow-Headers", "X-Token, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		ctx.Header("Access-Control-Allow-Credentials", "false")
		//放行所有options方法
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
	}
}
