package middle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"operations-platform/utils"
)

// 用户token认证
func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.URL.String() == "/api/login" {
			ctx.Next()
		}
		//获取和非空验证
		token := ctx.Request.Header.Get("Authorization")
		if len(token) == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 90403,
				"msg":  "请求未携带token,无权限访问",
				"data": nil,
			})
			ctx.Abort()
			return
		}
		//解析token
		claims, err := utils.JWTToken.ParseToken(token)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 90403,
				"msg":  err.Error(),
				"data": nil,
			})
			ctx.Abort()
			return
		}
		//把claims的数据放到上下文中
		//service想用这个claims的话，就接受gin.context上下文即可（入参）
		ctx.Set("claims", claims)
	}
}
