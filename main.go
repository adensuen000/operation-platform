package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"operations-platform/config"
	"operations-platform/controller"
	"operations-platform/db"
	"operations-platform/middle"
	"operations-platform/model"
)

var (
	GORM *gorm.DB
	err  error
)

func main() {
	r := gin.Default()
	//数据库初始化
	db.Init()
	//检查表并建表
	db.DB.AutoMigrate(&model.User{})
	db.DB.AutoMigrate(&model.WorkOrderList{})
	//db.DB.AutoMigrate(&model.OSWorkOrder{})
	db.DB.AutoMigrate(&model.MachineWorkOrder{})
	db.DB.AutoMigrate(&model.WorkOrderAssign{})

	//中间件必须加载路由注册前
	//中间件的全局注册
	r.Use(middle.Cors())
	//初始化路由
	controller.Router.InitApiRouter(r)

	//params := &model.MachineWorkOrder{
	//	MachineWorkOrderID: 1010,
	//	//Status:             "",
	//	//GPUQuantity:        "",
	//	OSDiskSize: "500000G",
	//}
	//
	//data, err2 := service.MWO.DealMWO(params)
	//fmt.Println("data:", data)
	//fmt.Println("err:", err2)

	//启动
	r.Run(config.ListAddr)
}
