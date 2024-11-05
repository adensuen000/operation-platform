package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"operations-platform/model"
	"operations-platform/service"
)

var WOAssign workOrderAssign

type workOrderAssign struct {
}

// 创建工单流
func (*workOrderAssign) CreateWOAssign(ctx *gin.Context) {
	params := &model.WorkOrderAssign{}
	if err := ctx.ShouldBindJSON(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 90400,
			"msg":  "请求错误: " + err.Error(),
			"data": nil,
		})
		return
	}
	//fmt.Println("=============controller======================")
	//fmt.Println("params.TicketID: ", params.TicketID)
	//fmt.Println("params.ForwardUserID: ", params.ForwardUserID)
	//fmt.Println("params.RecieveUserID: ", params.RecieveUserID)
	//fmt.Println("params.ArriveTime: ", params.ArriveTime)
	//fmt.Println("params.ForwardComment: ", params.ForwardComment)
	data, err := service.WOAssign.CreateWOAssign(params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 90500,
			"msg":  "请求错误: " + err.Error(),
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

// 获取工单流
func (*workOrderAssign) GetWOAssign(ctx *gin.Context) {
	params := new(struct {
		TicketID int64 `form:"ticket_id"`
	})

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  90400,
			"msg":   "请求错误: " + err.Error(),
			"data":  nil,
			"total": 0,
		})
		return
	}
	data, total, err := service.WOAssign.GetWOAssign(params.TicketID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  90500,
			"msg":   "请求错误: " + err.Error(),
			"data":  nil,
			"total": total,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":  90200,
		"msg":   "请求成功.",
		"data":  data,
		"total": total,
	})
}
