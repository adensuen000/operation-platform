package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Router router

type router struct {
}

// 初始化路由规则
func (*router) InitApiRouter(r *gin.Engine) {
	r.GET("/testapi", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "testapi success",
			"data": nil,
		})
	}).
		//用户相关
		POST("/register", User.Register).
		POST("/api/login", User.Login).
		GET("/api/user/getUserList", User.GetUsers).
		//Use(middle.JWTAuth()).
		GET("/api/user/list", User.UserQuery).
		GET("/api/user/getUserID", User.GetUserID).
		//工单接口
		GET("/api/workOrder/list", WorkOrder.GetWorkOrder).
		POST("/api/workOrder/create", WorkOrder.CreateWorkOrder).
		//安装操作系统工单接口
		POST("/api/oswo/create", OSWorkOrder.CreateOSWorkOrder).
		GET("/api/oswo/list", OSWorkOrder.GetOSWorkOrder).
		//整机交付的工单接口
		POST("/api/machineWO/create", MWO.CreateMachineWorkOrder).
		GET("/api/machineWO/list", MWO.GetMachineWorkOrder).
		POST("/api/machineWO/update", MWO.UpdateMachineWorkOrder).
		DELETE("/api/machineWO/del", MWO.DelFromMachineWorkOrder).
		POST("/api/machineWO/deal", MWO.DealMWO).
		//工作流接口
		POST("/api/woAssign/create", WOAssign.CreateWOAssign).
		GET("/api/woAssign/list", WOAssign.GetWOAssign)

}
