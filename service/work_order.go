package service

import (
	"operations-platform/dao"
	"operations-platform/model"
)

var WorkOrder workOrder

type workOrder struct {
}

// 查询工单
func (w workOrder) GetWorkOrder(title, customerName string, pageSize, page int) (int, []*model.WorkOrderList, error) {
	total, data, err := dao.WorkOrder.GetWorkOrder(title, customerName, pageSize, page)
	if err != nil {
		return 0, nil, err
	}
	return total, data, nil
}

// 创建工单
func (*workOrder) CreateWorkOrder(wol *model.WorkOrderList) (*model.WorkOrderList, error) {
	data, err := dao.WorkOrder.CreateWorkOrder(wol)
	if err != nil {
		return nil, err
	}
	return data, nil
}
