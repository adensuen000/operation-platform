package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"operations-platform/model"
	"operations-platform/service"
)

var User user

type user struct {
}

// 注册
func (*user) Register(c *gin.Context) {
	p := &model.User{}
	//专门用于标签为json的参数接收，也就是post,put,delete这种带body的请求
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 90401,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	//创建用户
	if err := service.User.Add(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 90500,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "创建用户成功.",
		"data": p.Username,
	})
}

// 登录
func (*user) Login(c *gin.Context) {
	p := &model.User{}
	//专门用于标签为json的参数接收，也就是post,put,delete这种带body的请求
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 90402,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	//校验账号密码
	if err := service.User.Login(p.Username, p.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 90500,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	//更新token
	if err := service.User.UpdateToken(p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 90500,
			"msg":  err.Error(),
			"data": nil,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "登录成功.",
		"data": nil,
	})
}

// 查询用户
func (*user) UserQuery(ctx *gin.Context) {
	u := &model.User{}
	if err := ctx.ShouldBindJSON(u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 90400,
			"msg":  "请求出错",
			"data": nil,
		})
		return
	}
	data, res, err := service.User.UserQuery(u.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 90500,
			"msg":  "内部出错",
			"data": nil,
		})
		return
	}
	if !res {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 90200,
			"msg":  "未查询到用户",
			"data": nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 90200,
		"msg":  "查询到用户",
		"data": data,
	})
}

// 获取用户userID
func (*user) GetUserID(ctx *gin.Context) {
	u := new(struct {
		Username string `form:"username"`
	})
	if err := ctx.ShouldBindQuery(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   90400,
			"msg":    "请求出错",
			"userID": 0,
		})
		return
	}
	fmt.Println("u.Username: ", u.Username)
	userID, err := service.User.GetUserID(u.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":   90500,
			"msg":    err.Error(),
			"userID": userID,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":   90200,
		"msg":    nil,
		"userID": userID,
	})
}

// 根据字符串查询用户名
func (*user) GetUsers(ctx *gin.Context) {
	p := new(struct {
		FilterString string `form:"filter_string"`
	})

	if err := ctx.ShouldBindQuery(&p); err != nil {
		fmt.Println("400: ", p.FilterString)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 90400,
			"msg":  "请求出错",
			"data": nil,
		})
		return
	}
	data, err := service.User.GetUsers(p.FilterString)
	if err != nil {
		fmt.Println("500: ", p.FilterString)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 90500,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	fmt.Println("200: ", p.FilterString)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 90200,
		"msg":  "请求成功.",
		"data": data,
	})
}
