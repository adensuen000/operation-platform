package service

import (
	"operations-platform/dao"
	"operations-platform/model"
)

var WOAssign workOrderAssign

type workOrderAssign struct {
}

// 创建工单流
func (*workOrderAssign) CreateWOAssign(woa *model.WorkOrderAssign) (*model.WorkOrderAssign, error) {
	data, err := dao.WOAssign.CreateWOAssign(woa)
	if err != nil {
		return nil, err
	}
	//工单流有数据更新则通过钉钉发送消息
	_, _ = DingMsg.SendWOMsg(woa)
	return data, nil
}

// 获取工单流
func (*workOrderAssign) GetWOAssign(ticketID int64) ([]*model.WorkOrderAssign, int, error) {
	data, total, err := dao.WOAssign.GetWOAssign(ticketID)
	if err != nil {
		return nil, total, err
	}
	return data, total, nil
}
