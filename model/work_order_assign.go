package model

import (
	"time"
)

type WorkOrderAssign struct {
	AssignID       int           `gorm:"primaryKey;autoIncrement;autoIncrement:1" json:"assign_id"`
	WorkOrderList  WorkOrderList `gorm:"ForeignKey:TicketID;AssociationForeignKey:TicketID"`
	TicketID       int64         `json:"ticket_id"`
	ForwardUser    User          `gorm:"ForeignKey:ForwardUserID;references:UserID"`
	ForwardUserID  int           `gorm:"ForeignKey:UserID" json:"forward_user_id"`
	RecieveUser    User          `gorm:"ForeignKey:CurrentUserID;references:UserID"`
	RecieveUserID  int           `gorm:"ForeignKey:UserID" json:"recieve_user_id"`
	ArriveTime     time.Time     `json:"arrive_time"`
	ForwardComment string        `json:"forward_comment"`
}
