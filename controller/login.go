package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
	"operations-platform/service"
)

var Login login

type login struct {
}

func (l *login) Auth(ctx *gin.Context) {
	params := new(struct {
		Username string `json: "username"`
		Password string `json: "password"`
	})
	if err := ctx.ShouldBindJSON(params); err != nil {
		logger.Error("bind请求参数失败: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":  "请求出错: " + err.Error(),
			"data": nil,
		})
		return
	}
	if err := service.Login.Auth(params.Username, params.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "登录成功...",
		"data": nil,
	})
}
