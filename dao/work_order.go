package dao

import (
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	"operations-platform/db"
	"operations-platform/model"
)

var WorkOrder workOrder

type workOrder struct {
}

// 查询工单数据，返回结构体数据

func (*workOrder) GetWorkOrder(title, customerName string, pageSize, page int) (int, []*model.WorkOrderList, error) {
	//基于创建时间倒序排序
	var (
		ByTitle            = "title LIKE ?"
		ByCustomer         = "customer_name LIKE ?"
		ByTitleAndCustomer = "title LIKE ? AND customer_name LIKE ?"
		ByCreateTimeDesc   = "create_time DESC"
		wols               []*model.WorkOrderList
		total              int
	)

	if (title == "" && customerName == "") && (pageSize <= 0 || page <= 0) {
		db.DB.Order(ByCreateTimeDesc).Find(&wols).Count(&total)
		db.DB.Order(ByCreateTimeDesc).Limit(pageSize).Find(&wols)
		return total, wols, nil
	}
	if (title == "" && customerName == "") && pageSize != 0 && page != 0 {
		db.DB.Order(ByCreateTimeDesc).Find(&wols).Count(&total)
		db.DB.Order(ByCreateTimeDesc).Offset((page - 1) * pageSize).Limit(pageSize).Find(&wols)
		return total, wols, nil
	}
	if title == "" && customerName != "" && pageSize != 0 && page != 0 {
		db.DB.Where(ByCustomer, "%"+customerName+"%").Order(ByCreateTimeDesc).Find(&wols).Count(&total)
		db.DB.Where(ByCustomer, "%"+customerName+"%").Order(ByCreateTimeDesc).Offset((page - 1) * pageSize).Limit(pageSize).Find(&wols)
		return total, wols, nil
	}
	if title != "" && customerName == "" && pageSize != 0 && page != 0 {
		db.DB.Where(ByTitle, "%"+title+"%").Order(ByCreateTimeDesc).Find(&wols).Count(&total)
		db.DB.Where(ByTitle, "%"+title+"%").Order(ByCreateTimeDesc).Offset((page - 1) * pageSize).Limit(pageSize).Find(&wols)
		return total, wols, nil
	}
	if title != "" && customerName != "" && pageSize != 0 && page != 0 {
		db.DB.Where(ByTitleAndCustomer, "%"+title+"%", "%"+customerName+"%").Order(ByCreateTimeDesc).Find(&wols).Count(&total)
		db.DB.Where(ByTitleAndCustomer, "%"+title+"%", "%"+customerName+"%").Order(ByCreateTimeDesc).Offset((page - 1) * pageSize).Limit(pageSize).Find(&wols)
		return total, wols, nil
	}
	return total, nil, nil
}

// 创建工单
func (*workOrder) CreateWorkOrder(wol *model.WorkOrderList) (*model.WorkOrderList, error) {
	//fmt.Println("==============dao=============")
	//fmt.Println("TicketID: ", wol.TicketID)
	//fmt.Println("CustomerType: ", wol.CustomerType)
	//fmt.Println("title: ", wol.title)
	//fmt.Println("CreateTime: ", wol.CreateTime)
	//fmt.Println("UserID: ", wol.UserID)
	//fmt.Println("UpdatedTime: ", wol.UpdatedTime)
	//fmt.Println("CPUQuantity: ", wol.CPUQuantity)
	//fmt.Println("CustomerName: ", wol.CustomerName)
	//fmt.Println("DataDiskSize: ", wol.DataDiskSize)
	//fmt.Println("GPUQuantity: ", wol.GPUQuantity)
	//fmt.Println("MemorySize: ", wol.MemorySize)
	//fmt.Println("OSDiskSize: ", wol.OSDiskSize)
	//fmt.Println("OSType: ", wol.OSType)
	//fmt.Println("PCIEType: ", wol.PCIEType)
	//fmt.Println("PhoneNumber: ", wol.PhoneNumber)
	//fmt.Println("UtilizationStartDate: ", wol.UtilizationStartDate)
	//fmt.Println("UtilizationEndDate: ", wol.UtilizationEndDate)
	tx := db.DB.Create(wol)
	if tx.Error != nil {
		fmt.Println("AAAAAAA工单总表中创建工单失败: ", wol)
		logger.Error(fmt.Sprintf("工单总表中创建工单失败: %v\n", tx.Error))
		return nil, errors.New(fmt.Sprintf("工单总表中创建工单失败: %v\n", tx.Error))
	}
	fmt.Println("AAAAAAA工单总表中创建工单成功: ", wol)
	return wol, nil
}

// 删除数据
func (*workOrder) DelData(ticketID int64) (bool, error) {
	resWO := db.DB.Where("ticket_id = ? ", ticketID).Delete(&model.WorkOrderList{})
	if resWO.Error != nil {
		logger.Error(fmt.Sprintf("在工单总表中删除数据失败: ", resWO.Error))
		return false, errors.New(fmt.Sprintf("在工单总表中删除数据失败: ", resWO.Error))
	}
	return true, nil
}

// 通过ticketID获取工单数据
func (*machineWO) GetWOByTID(ticketID int64) (*model.WorkOrderList, error) {
	var (
		wo *model.WorkOrderList
		//统计数据条数
		count int
	)
	wo.TicketID = ticketID
	res := db.DB.Where("ticket_id = ? ", ticketID).Find(&wo).Count(&count)
	if res.Error != nil {
		return nil, errors.New(fmt.Sprintf("在工单总表中查询工单失败: ", res.Error))
	}
	if count == 0 {
		return nil, nil
	}
	//当count为1的时候符合预期
	return wo, nil
}
