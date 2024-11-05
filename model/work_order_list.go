package model

import (
	"time"
)

// 工单表结构

type WorkOrderList struct {
	TicketID     int64     `gorm:"primaryKey;autoIncrement;autoIncrement:1" json:"ticket_id"`
	CustomerType string    `json:"customer_type"`
	CustomerName string    `json:"customer_name"`
	PhoneNumber  string    `json:"phone_number"`
	BmcIP        string    `json:"bmc_ip"`
	BusinessIP   string    `json:"business_ip"`
	User         User      `gorm:"ForeignKey:UserID;not null"`
	UserID       int       `json:"user_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Status       string    `json:"status"`
	CreateTime   time.Time `json:"create_time"`
	UpdatedTime  time.Time `json:"updated_time"`
}
