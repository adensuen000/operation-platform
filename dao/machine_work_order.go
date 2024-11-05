package dao

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/wonderivan/logger"
	"operations-platform/db"
	"operations-platform/model"
)

var MWO machineWO

type machineWO struct {
}

// 获取工单数据
func (*machineWO) GetMachineWorkOrder(title, customerName string, pageSize, page int) (int, []*model.MachineWorkOrder, error) {
	//基于创建时间倒序排序
	var (
		ByTitle            = "title LIKE ?"
		ByCustomer         = "customer_name LIKE ?"
		ByTitleAndCustomer = "title LIKE ? AND customer_name LIKE ?"
		ByCreateTimeDesc   = "create_time DESC"
		mwos               []*model.MachineWorkOrder
		total              int
	)

	if (title == "" && customerName == "") && (pageSize <= 0 || page <= 0) {
		//获取total的值用于分页
		db.DB.Order(ByCreateTimeDesc).Find(&mwos).Count(&total)
		db.DB.Order(ByCreateTimeDesc).Limit(pageSize).Find(&mwos)
		return total, mwos, nil
	}
	if (title == "" && customerName == "") && pageSize != 0 && page != 0 {
		//获取total的值用于分页
		db.DB.Order(ByCreateTimeDesc).Find(&mwos).Count(&total)
		//预加载user表数据，获取username，并进行关联查询
		db.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("user_id,username")
		}).Find(&mwos).
			Order(ByCreateTimeDesc).
			Offset((page - 1) * pageSize).
			Limit(pageSize).
			Find(&mwos)
		return total, mwos, nil
	}
	if title == "" && customerName != "" && pageSize != 0 && page != 0 {
		//获取total的值用于分页
		db.DB.Where(ByCustomer, "%"+customerName+"%").Order(ByCreateTimeDesc).Find(&mwos).Count(&total)
		//预加载user表数据，获取username，并进行关联查询
		db.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("user_id,username")
		}).Find(&mwos).
			Where(ByCustomer, "%"+customerName+"%").
			Order(ByCreateTimeDesc).
			Offset((page - 1) * pageSize).
			Limit(pageSize).
			Find(&mwos)
		return total, mwos, nil
	}
	if title != "" && customerName == "" && pageSize != 0 && page != 0 {
		//获取total的值用于分页
		db.DB.Where(ByTitle, "%"+title+"%").Order(ByCreateTimeDesc).Find(&mwos).Count(&total)
		//预加载user表数据，获取username，并进行关联查询
		db.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("user_id,username")
		}).Find(&mwos).
			Where(ByTitle, "%"+title+"%").
			Order(ByCreateTimeDesc).
			Offset((page - 1) * pageSize).
			Limit(pageSize).
			Find(&mwos)
		return total, mwos, nil
	}
	if title != "" && customerName != "" && pageSize != 0 && page != 0 {
		//获取total的值用于分页
		db.DB.Where(ByTitleAndCustomer, "%"+title+"%", "%"+customerName+"%").Order(ByCreateTimeDesc).Find(&mwos).Count(&total)
		//预加载user表数据，获取username，并进行关联查询
		db.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("user_id,username")
		}).Find(&mwos).
			Where(ByTitleAndCustomer, "%"+title+"%", "%"+customerName+"%").
			Order(ByCreateTimeDesc).
			Offset((page - 1) * pageSize).
			Limit(pageSize).
			Find(&mwos)
		return total, mwos, nil
	}
	return total, nil, nil
}

// 在MachineWorkOrder表中插入新数据
func (*machineWO) CreateMachineWorkOrder(mwo *model.MachineWorkOrder) (*model.MachineWorkOrder, error) {
	mwoRes := db.DB.Create(mwo)
	if mwoRes.Error != nil {
		logger.Error(fmt.Sprintf("整机交付工单表中创建工单失败: %v\n", mwoRes.Error))
		return nil, errors.New(fmt.Sprintf("整机交付工单表中创建工单失败: ", mwoRes.Error))
	}
	return mwo, nil
}

// 更新数据
func (*machineWO) UpdateMachineWorkOrder(mwo *model.MachineWorkOrder) (bool, error) {
	//更新整机交付工单的同时需要在工单总表中创建同样的工单
	wo := &model.WorkOrderList{}
	wo.CustomerType, wo.CustomerName = mwo.CustomerType, mwo.CustomerName
	wo.UserID = mwo.UserID
	wo.TicketID = mwo.TicketID

	res := db.DB.Save(wo)
	if res.Error != nil {
		logger.Error(fmt.Sprintf("MachineWorkOrder表中更新工单失败: %v\n", res.Error))
		return false, errors.New(fmt.Sprintf("MachineWorkOrder表中更新工单失败: %v\n", res.Error))
	}
	return true, nil
}

// 删除MachineWorkOrder表数据
func (*machineWO) DelData(ticketID int64) (bool, error) {
	resMWO := db.DB.Where("ticket_id = ? ", ticketID).Delete(&model.MachineWorkOrder{})
	if resMWO.Error != nil {
		return false, errors.New(fmt.Sprintf("删除MachineWorkOrder表数据失败: ", resMWO.Error))
	}
	return true, nil
}

// 以机器工单表为入口删除相关工单数据
func (*machineWO) DelFromMachineWorkOrder(ticketID int64) (bool, error) {

	resMWO, err := MWO.DelData(ticketID)
	if err != nil {
		return resMWO, err
	}
	return resMWO, nil
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

//创建工单后，后续第二阶段，第三阶段的人处理工单，即补充工单信息

func (*machineWO) DealMWO(mwo *model.MachineWorkOrder) (*model.MachineWorkOrder, error) {
	if err := db.DB.Model(&model.MachineWorkOrder{}).Where("ticket_id = ?", mwo.TicketID).Updates(mwo).Error; err != nil {
		logger.Error(fmt.Sprintf("处理工单：MachineWorkOrder表中更新工单失败: %v\n", err.Error))
		return nil, errors.New(fmt.Sprintf("处理工单：MachineWorkOrder表中更新工单失败: %v\n", err.Error))
	}
	return mwo, nil
}
