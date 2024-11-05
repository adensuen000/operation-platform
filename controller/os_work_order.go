package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"operations-platform/model"
	"operations-platform/service"
	"reflect"
)

var OSWorkOrder osWorkOrder

type osWorkOrder struct {
}

// 获取OS工单数据
func (*osWorkOrder) GetOSWorkOrder(ctx *gin.Context) {
	//匿名结构体，用于定义入参
	params := new(struct {
		Title        string `form:"title"`
		CustomerName string `form:"customer_name"`
		PageSize     int    `form:"page_size"`
		Page         int    `form:"page"`
	})

	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 90400,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	//获取数据
	total, data, err := service.OSWorkOrder.GetOSWorkOrder(params.Title, params.CustomerName, params.PageSize, params.Page)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  90500,
			"msg":   err.Error(),
			"data":  nil,
			"total": total,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":  90200,
		"msg":   "数据查询成功",
		"data":  data,
		"total": total,
	})
}

// 创建OS工单
func (*osWorkOrder) CreateOSWorkOrder(ctx *gin.Context) {

	//匿名结构体，用于定义入参
	params := &model.OSWorkOrder{}
	if err := ctx.ShouldBindJSON(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 90400,
			"msg":  "请求错误，请检查.",
			"data": nil,
		})
		return
	}
	fmt.Println("===========================controller===========================")
	fmt.Println("params: ", params)

	//打印结构体的字段

	for i := 0; i < reflect.TypeOf(params).Elem().NumField(); i++ {
		prop := reflect.TypeOf(params).Elem().Field(i).Name
		fmt.Println("变量名称: ", prop, "变量值: ", reflect.ValueOf(params).Elem().FieldByName(prop))
	}

	data, err := service.OSWorkOrder.CreateOSWorkOrder(params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 90500,
			"msg":  "服务内部错误，请检查.",
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 90200,
		"msg":  "请求成功.",
		"data": data,
	})
}
