package dao

import (
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	"operations-platform/db"
	"operations-platform/model"
)

var WOAssign workOrderAssign

type workOrderAssign struct {
}

// 向工单流表中插入数据
func (*workOrderAssign) CreateWOAssign(woa *model.WorkOrderAssign) (*model.WorkOrderAssign, error) {
	errWOA := db.DB.Create(woa)
	if errWOA.Error != nil {
		logger.Error(fmt.Sprintf("向工单流表中插入数据失败: ", errWOA))
		return nil, errors.New(fmt.Sprintf("向工单流表中插入数据失败: ", errWOA))
	}
	return woa, nil
}

// 获取工单流
func (*workOrderAssign) GetWOAssign(ticketID int64) ([]*model.WorkOrderAssign, int, error) {
	var (
		woa               []*model.WorkOrderAssign
		total             int
		orderByArriveTime = "arrive_time DESC"
	)
	//获取总条数
	db.DB.Where("ticket_id = ?", ticketID).Order(orderByArriveTime).Find(&woa).Count(&total)
	resWOA := db.DB.Where("ticket_id = ?", ticketID).Order(orderByArriveTime).Find(&woa)
	if resWOA.Error != nil {
		logger.Error(fmt.Sprintf("获取工单流失败: ", resWOA.Error))
		return nil, total, errors.New(fmt.Sprintf("获取工单流失败: ", resWOA.Error))
	}

	return woa, total, nil
}

// 删除工单流表中关于某个工单的所有数据
func (*workOrderAssign) DelData(ticketID int64) (bool, error) {
	resWOA := db.DB.Where("ticket_id = ? ", ticketID).Delete(&model.WorkOrderAssign{})
	if resWOA.Error != nil {
		return false, errors.New(fmt.Sprintf("删除MachineWorkOrder表数据失败: ", resWOA.Error))
	}
	return true, nil
}
