package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"operations-platform/model"
	"operations-platform/service"
)

var MWO machineWO

type machineWO struct {
}

// 查询机器交付的工单
func (*machineWO) GetMachineWorkOrder(ctx *gin.Context) {
	//匿名结构体，用于定义入参
	params := new(struct {
		Title        string `form:"title"`
		CustomerName string `form:"customer_name"`
		PageSize     int    `form:"page_size"`
		Page         int    `form:"page"`
	})
	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  90400,
			"msg":   "请求出错: " + err.Error(),
			"data":  nil,
			"total": 0,
		})
		return
	}
	//获取数据
	total, data, err := service.MWO.GetMachineWorkOrder(params.Title, params.CustomerName, params.PageSize, params.Page)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  90500,
			"msg":   "数据查询失败:" + err.Error(),
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

// 创建机器交付的工单
func (*machineWO) CreateMachineWorkOrder(ctx *gin.Context) {
	params := &model.MachineWorkOrder{}

	if err := ctx.ShouldBindJSON(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 90400,
			"msg":  "请求出错: " + err.Error(),
			"data": nil,
		})
		return
	}
	//fmt.Printf("%+v\n", params)
	data, err := service.MWO.CreateMachineWorkOrder(params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 90500,
			"msg":  "创建工单失败.",
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 90200,
		"msg":  "创建工单成功.",
		"data": data,
	})
}

// 更新工单
func (*machineWO) UpdateMachineWorkOrder(ctx *gin.Context) {
	params := &model.MachineWorkOrder{}
	if err := ctx.ShouldBindBodyWith(params, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 90400,
			"msg":  "请求出错: " + err.Error(),
			"res":  false,
		})
		return
	}
	res, err := service.MWO.UpdateMachineWorkOrder(params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 90500,
			"msg":  "更新工单失败.",
			"res":  res,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 90200,
		"msg":  "更新工单成功.",
		"res":  res,
	})
}

// 删除工单
func (*machineWO) DelFromMachineWorkOrder(ctx *gin.Context) {
	params := new(struct {
		TicketID int64 `json:"ticket_id"`
	})
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 90400,
			"msg":  "请求出错: " + err.Error(),
			"res":  false,
		})
		return
	}
	res, err := service.MWO.DelFromMachineWorkOrder(params.TicketID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 90500,
			"msg":  "删除工单失败.",
			"res":  res,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 90200,
		"msg":  "删除工单成功.",
		"res":  res,
	})
}

// 创建工单后，后续第二阶段，第三阶段的人处理工单，即补充工单信息
func (*machineWO) DealMWO(ctx *gin.Context) {
	params := &model.MachineWorkOrder{}
	if err := ctx.ShouldBindJSON(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 90400,
			"msg":  "请求出错: " + err.Error(),
			"data": nil,
		})
		return
	}

	data, err := service.MWO.DealMWO(params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 90500,
			"msg":  "服务内部错误: " + err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 90200,
		"msg":  "处理工单成功.",
		"res":  data,
	})
}
