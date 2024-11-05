package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"operations-platform/model"
	"operations-platform/service"
)

var WorkOrder workOrder

type workOrder struct {
}

// 查询工单
func (w *workOrder) GetWorkOrder(ctx *gin.Context) {
	//匿名结构体，用于定义入参
	params := new(struct {
		Title         string `form:"title"`
		Customer_name string `form:"customer_name"`
		PageSize      int    `form:"page_size"`
		Page          int    `form:"page"`
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
	total, data, err := service.WorkOrder.GetWorkOrder(params.Title, params.Customer_name, params.PageSize, params.Page)
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

// 创建工单
func (w *workOrder) CreateWorkOrder(ctx *gin.Context) {
	params := &model.WorkOrderList{}
	if err := ctx.ShouldBindJSON(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 90400,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	//创建工单
	data, err := service.WorkOrder.CreateWorkOrder(params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 90500,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 90200,
		"msg":  "创建工单成功",
		"data": data,
	})
}
