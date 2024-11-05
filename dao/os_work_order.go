package dao

import (
	"errors"
	"fmt"
	"operations-platform/db"
	"operations-platform/model"
)

var OSWorkOrder osWorkOrder

type osWorkOrder struct {
}

// 查询OS工单
func (*osWorkOrder) GetOSWorkOrder(title, customerName string, pageSize, page int) (int, []*model.OSWorkOrder, error) {
	//基于创建时间倒序排序
	var (
		ByTitle            = "title LIKE ?"
		ByCustomer         = "customer_name LIKE ?"
		ByTitleAndCustomer = "title LIKE ? AND customer_name LIKE ?"
		ByCreateTimeDesc   = "create_time DESC"
		oswos              []*model.OSWorkOrder
		total              int
	)

	if (title == "" && customerName == "") && (pageSize <= 0 || page <= 0) {
		db.DB.Order(ByCreateTimeDesc).Find(&oswos).Count(&total)
		db.DB.Order(ByCreateTimeDesc).Limit(pageSize).Find(&oswos)
		return total, oswos, nil
	}
	if (title == "" && customerName == "") && pageSize != 0 && page != 0 {
		db.DB.Order(ByCreateTimeDesc).Find(&oswos).Count(&total)
		db.DB.Order(ByCreateTimeDesc).Offset((page - 1) * pageSize).Limit(pageSize).Find(&oswos)
		return total, oswos, nil
	}
	if title == "" && customerName != "" && pageSize != 0 && page != 0 {
		db.DB.Where(ByCustomer, "%"+customerName+"%").Order(ByCreateTimeDesc).Find(&oswos).Count(&total)
		db.DB.Where(ByCustomer, "%"+customerName+"%").Order(ByCreateTimeDesc).Offset((page - 1) * pageSize).Limit(pageSize).Find(&oswos)
		return total, oswos, nil
	}
	if title != "" && customerName == "" && pageSize != 0 && page != 0 {
		db.DB.Where(ByTitle, "%"+title+"%").Order(ByCreateTimeDesc).Find(&oswos).Count(&total)
		db.DB.Where(ByTitle, "%"+title+"%").Order(ByCreateTimeDesc).Offset((page - 1) * pageSize).Limit(pageSize).Find(&oswos)
		return total, oswos, nil
	}
	if title != "" && customerName != "" && pageSize != 0 && page != 0 {
		db.DB.Where(ByTitleAndCustomer, "%"+title+"%", "%"+customerName+"%").Order(ByCreateTimeDesc).Find(&oswos).Count(&total)
		db.DB.Where(ByTitleAndCustomer, "%"+title+"%", "%"+customerName+"%").Order(ByCreateTimeDesc).Offset((page - 1) * pageSize).Limit(pageSize).Find(&oswos)
		return total, oswos, nil
	}
	return total, nil, nil
}

// 创建OS工单
func (*osWorkOrder) CreateOSWorkOrder(oswo *model.OSWorkOrder) (*model.OSWorkOrder, error) {

	//在os_work_orders表中插入新数据
	fmt.Println("开始向OS工单表中插入数据")
	oswoRes := db.DB.Create(oswo)
	if oswoRes.Error != nil {
		fmt.Println("OS工单表中创建工单失败---------")
		return nil, errors.New(fmt.Sprintf("在OS工单表中创建工单失败: ", oswoRes.Error))
	}
	fmt.Println("OS工单表中创建工单成功")
	return oswo, nil
}

func (*osWorkOrder) UpdateOSWorkOrder(oswo *model.OSWorkOrder) (bool, error) {
	return false, nil
}

func (*osWorkOrder) DelOSWorkOrder(oswo *model.OSWorkOrder) (bool, error) {
	return false, nil
}
