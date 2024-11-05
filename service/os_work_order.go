package service

import (
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	"operations-platform/dao"
	"operations-platform/db"
	"operations-platform/model"
)

var OSWorkOrder osWorkOrder

type osWorkOrder struct {
}

// 查询工单
func (*osWorkOrder) GetOSWorkOrder(title, customerName string, pageSize, page int) (int, []*model.OSWorkOrder, error) {

	total, data, err := dao.OSWorkOrder.GetOSWorkOrder(title, customerName, pageSize, page)
	if err != nil {
		return 0, nil, err
	}
	return total, data, nil
}

// 创建工单
func (*osWorkOrder) CreateOSWorkOrder(oswo *model.OSWorkOrder) (*model.OSWorkOrder, error) {

	//创建OS工单的同时需要在工单总表中创建同样的工单
	wo := &model.WorkOrderList{}
	wo.CustomerType, wo.CustomerName = oswo.CustomerType, oswo.CustomerName
	wo.Title, wo.Description, wo.CreateTime, wo.UpdatedTime = oswo.Title, oswo.Description, oswo.CreateTime, oswo.UpdatedTime
	wo.UserID, wo.Status = oswo.UserID, oswo.Status

	//先向工单总表插入数据
	woData, err := WorkOrder.CreateWorkOrder(wo)
	if err != nil {
		logger.Error(fmt.Sprintf("工单总表中创建工单失败: %v\n", err))
		return nil, errors.New(fmt.Sprintf("工单总表中创建工单失败: %v\n", err))
	}

	//再创建OS工单
	//获取ticket_id
	db.DB.Order("create_time DESC").First(&wo, "user_id = ? ", wo.UserID)
	oswo.TicketID = wo.TicketID

	oswoData, err := dao.OSWorkOrder.CreateOSWorkOrder(oswo)
	if err != nil {
		//若os_work_orders表中插入数据失败，也删除工单总表中的数据
		_, woDelRes := dao.WorkOrder.DelData(woData.TicketID)
		return nil, errors.New(fmt.Sprintf("OS工单表中创建工单失败,删除工单总表中的数据: %v\n", woDelRes))
	}
	return oswoData, nil
}
