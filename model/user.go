package model

import (
	"database/sql"
	"time"
)

type User struct {
	UserID      int          `gorm:"primaryKey;autoIncrement;autoIncrement:1" json:"user_id"`
	Username    string       `gorm:"not null" json:"username"`
	PhoneNumber string       `json:"phone_number"`
	Password    string       `gorm:"not null" json:"password"`
	Token       string       `json:"token"`
	Role        string       `json:"role"`
	CreateTime  time.Time    `json:"create_time"`
	UpdatedTime time.Time    `json:"update_time"`
	DeletedAt   sql.NullTime `gorm:"type:TIMESTAMP NULL" json:"delete_at"`
}
