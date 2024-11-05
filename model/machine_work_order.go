package model

import (
	"time"
)

// 机器交付表
//type MachineWorkOrder struct {
//	MachineWorkOrderID int64         `gorm:"primaryKey;autoIncrement;autoIncrement;not null" json:"machine_work_order_id"`
//	WorkOrderList      WorkOrderList `gorm:"ForeignKey:TicketID;AssociationForeignKey:TicketID;not null"`
//	TicketID           int64         `gorm:"not null" json:"ticket_id"`
//	Title              string        `gorm:"not null" json:"title"`
//	//客户信息
//	CustomerType         string    `gorm:"not null" json:"customer_type"`
//	CustomerName         string    `gorm:"not null" json:"customer_name"`
//	PhoneNumber          string    `gorm:"not null" json:"phone_number"`
//	GPUQuantity          string    `gorm:"not null" json:"gpu_quantity"`
//	PCIEType             string    `gorm:"not null" json:"pcie_type"`
//	CPUQuantity          string    `gorm:"not null" json:"cpu_quantity"`
//	MemorySize           string    `gorm:"not null" json:"memory_size"`
//	OSDiskSize           string    `gorm:"not null" json:"os_disk_size"`
//	DataDiskSize         string    `gorm:"not null" json:"data_disk_size"`
//	OSType               string    `gorm:"not null" json:"os_type"`
//	GraphicsCardType     string    `gorm:"not null" json:"graphics_card_type"`
//	UtilizationStartDate time.Time `gorm:"not null" json:"utilization_start_date"`
//	UtilizationEndDate   time.Time `gorm:"not null" json:"utilization_end_date"`
//	User                 User      `gorm:"ForeignKey:UserID;AssociationForeignKey:UserID"`
//	UserID               int64     `gorm:"not null" json:"user_id"`
//	Description          string    `gorm:"not null" json:"description"`
//	Status               string    `gorm:"not null" json:"status"`
//	CreateTime           time.Time `json:"create_time"`
//	UpdatedTime          time.Time `json:"updated_time"`
//}

// 整机交付表（虚拟机，物理机都是这张表）
type MachineWorkOrder struct {
	MachineWorkOrderID int64         `gorm:"primaryKey;autoIncrement;autoIncrement;not null" json:"machine_work_order_id"`
	WorkOrderList      WorkOrderList `gorm:"ForeignKey:TicketID;AssociationForeignKey:TicketID;not null"`
	TicketID           int64         `gorm:"not null" json:"ticket_id"`

	//客户信息
	CustomerType         string    `json:"customer_type"`
	CustomerName         string    `json:"customer_name"`
	PhoneNumber          string    `json:"phone_number"`
	UsageType            string    `json:"usage_type"`
	SSHPubNetMapping     string    `json:"ssh_pubnet_mapping"`
	OtherPubNetMapping   string    `json:"other_pubnet_mapping"`
	UtilizationStartDate time.Time `json:"utilization_start_date"`
	UtilizationEndDate   time.Time `json:"utilization_end_date"`

	//资源信息
	MachineRoomLocation string `json:"machine_room_location"`
	MachineType         string `json:"machine_type"`
	BusinessIP          string `json:"business_ip"`
	PCIEType            string `json:"pcie_type"`
	GraphicsCardType    string `json:"graphics_card_type"`
	GPUQuantity         string `json:"gpu_quantity"`
	CPUQuantity         string `json:"cpu_quantity"`
	MemorySize          string `json:"memory_size"`
	OSDiskSize          string `json:"os_disk_size"`
	DataDiskSize        string `json:"data_disk_size"`
	OSType              string `json:"os_type"`
	BmcSwitchPort       string `json:"bmc_switch_port"`
	BusipSwitchPort     string `json:"busip_switch_port"`

	//带外信息
	BmcIP       string `json:"bmc_ip"`
	BmcUser     string `json:"bmc_user"`
	BmcPassword string `json:"bmc_password"`

	//网络信息
	BondMode string `json:"bond_mode"`
	VlanId   string `json:"vlan_id"`
	Gateway  string `json:"gateway"`
	Netmask  string `json:"netmask"`

	//映射数据
	SSHMappingData   string `json:"ssh_mapping_data"`
	OtherMappingData string `json:"other_mapping_data"`

	//vpn 信息
	VpnAddress  string     `json:"vpn_address"`
	VpnUser     string     `json:"vpn_user"`
	VpnPassword string     `json:"vpn_password"`
	VpnDDL      *time.Time `json:"vpn_ddl"`

	//系统信息
	OSUser     string `json:"os_user"`
	OSPassword string `json:"os_password"`
	SSHRsa     string `json:"ssh_rsa"`
	SSHRsaPub  string `json:"ssh_rsa_pub"`

	//工单描述
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`

	User        User      `gorm:"ForeignKey:UserID;AssociationForeignKey:UserID"`
	UserID      int       `json:"user_id"`
	CreateTime  time.Time `json:"create_time"`
	UpdatedTime time.Time `json:"updated_time"`
}
