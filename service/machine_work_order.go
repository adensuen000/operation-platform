package service

import (
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	"operations-platform/dao"
	"operations-platform/db"
	"operations-platform/model"
)

var MWO machineWO

type machineWO struct {
}

// 查询机器交付的工单
func (*machineWO) GetMachineWorkOrder(title, customerName string, pageSize, page int) (int, []*model.MachineWorkOrder, error) {
	total, data, err := dao.MWO.GetMachineWorkOrder(title, customerName, pageSize, page)
	if err != nil {
		return 0, nil, err
	}
	return total, data, nil
}

// 创建机器交付的工单
func (*machineWO) CreateMachineWorkOrder(mwo *model.MachineWorkOrder) (*model.MachineWorkOrder, error) {
	var (
		wo  = &model.WorkOrderList{}
		woa = &model.WorkOrderAssign{}
	)

	//创建整机交付工单的同时需要在工单总表中创建同样的工单
	//向工单总表中插入数据前的赋值
	wo.CustomerType = mwo.CustomerType
	wo.CustomerName = mwo.CustomerName
	wo.PhoneNumber = mwo.PhoneNumber
	wo.Title = mwo.Title
	wo.Description = mwo.Description
	wo.CreateTime = mwo.CreateTime
	wo.UpdatedTime = mwo.UpdatedTime
	wo.UserID = mwo.UserID
	wo.Status = mwo.Status

	//在work_order表中插入数据
	_, errWO := dao.WorkOrder.CreateWorkOrder(wo)
	if errWO != nil {
		logger.Error(fmt.Sprintf("首先在工单总表中创建工单: ", errWO))
		return nil, errors.New(fmt.Sprintf("首先在工单总表中创建工单: ", errWO))
	}

	//获取ticket_id
	db.DB.Order("create_time DESC").First(&wo, "user_id = ? ", wo.UserID)
	mwo.TicketID = wo.TicketID

	//在machine_work_order表中插入新数据
	dataMWO, errMWO := dao.MWO.CreateMachineWorkOrder(mwo)
	if errMWO != nil {
		logger.Error(fmt.Sprintf("整机交付工单表中创建工单失败后，删除工单总表中的数据失败: ", errWO))
		return nil, errors.New(fmt.Sprintf("整机交付工单表中创建工单失败后，删除工单总表中的数据失败: ", errWO))
		//若machine_work_order表中插入数据失败，则删除刚刚向work_order表中插入的数据
		//_, err := dao.WorkOrder.DelData(wo.TicketID)
		//if err != nil {
		//	logger.Error(fmt.Sprintf("整机交付工单表中创建工单失败后，删除工单总表中的数据失败: ", errWO))
		//	return nil, errors.New(fmt.Sprintf("整机交付工单表中创建工单失败后，删除工单总表中的数据失败: ", errWO))
		//}
	}

	//创建工单时，需要创建一条工作流信息
	//向工单流表中插入数据前的赋值
	woa.TicketID = mwo.TicketID
	woa.RecieveUserID = mwo.UserID
	woa.ArriveTime = mwo.CreateTime
	woa.ForwardUserID = mwo.UserID

	dataWOA, errWOA := dao.WOAssign.CreateWOAssign(woa)
	if errWOA != nil {
		logger.Error(fmt.Sprintf("向工单流表中插入数据失败: ", errWO))
		logger.Info(fmt.Sprintf("插入失败的数据如下: ", dataWOA))
	}

	return dataMWO, nil
}

// 更新工单
func (*machineWO) UpdateMachineWorkOrder(mwo *model.MachineWorkOrder) (bool, error) {
	res, err := dao.MWO.UpdateMachineWorkOrder(mwo)
	if err != nil {
		return res, err
	}
	return res, nil
}

// 删除工单
func (*machineWO) DelFromMachineWorkOrder(ticketID int64) (bool, error) {

	//先删除工单流表中相关工单数据
	resWOA, err := dao.WOAssign.DelData(ticketID)
	if err != nil {
		return resWOA, err
	}

	//然后删除整机交付表中的数据
	resMWO, err := dao.MWO.DelFromMachineWorkOrder(ticketID)
	if err != nil {
		return resMWO, err
	}
	//最后删除工单总表中的数据
	resWO, err := dao.WorkOrder.DelData(ticketID)
	if err != nil {
		return resWO, err
	}

	return true, nil
}

// 创建工单后，后续第二阶段，第三阶段的人处理工单，即补充工单信息
func (*machineWO) DealMWO(mwo *model.MachineWorkOrder) (*model.MachineWorkOrder, error) {
	fmt.Println("service: MachineWorkOrderID: ", mwo.MachineWorkOrderID)
	data, err := dao.MWO.DealMWO(mwo)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// 通过ticketID获取工单数据
func (*machineWO) GetMWOByTID(ticketID int64) (*model.MachineWorkOrder, error) {
	var (
		mwo = &model.MachineWorkOrder{}
	)
	mwo.TicketID = ticketID
	res := db.DB.Where("ticket_id = ? ", ticketID).Find(mwo)
	if res.Error != nil {
		return nil, errors.New(fmt.Sprintf("在机器工单表中查询工单失败: ", res.Error))
	}
	return mwo, nil
}
