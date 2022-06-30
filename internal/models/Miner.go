/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-10-17 14:17:14
 */
package models

import (
	"time"
)

// type Collections struct {
// 	BaseModle
// 	AccountPubkey string `json:"account_pubkey"`
// 	CPUCores string `json:"cpu_cores"`
// 	CPUUsage string `json:"cpu_usage"`
// 	IPAddress string `json:"ip_address"`
// 	MemTotal string `json:"mem_total"`
// 	MemUsage string `json:"mem_usage"`
// 	MinerCount string `json:"miner_count"`
// 	NodeCount string `json:"node_count"`
// 	NodeGraffi string `json:"node_graffi"`
// 	NodeHeight string `json:"node_height"`
// 	NodeVersion string `json:"node_version"`
// 	UpdateTime string `json:"update_time"`
// }

type Collections struct {
	BaseModle
	MachineId int `json:"machine_id"`
	IPAddress string `json:"ip_address"`
	MachineTag string `json:"machine_tag"`
	WorkStatus string `json:"work_status"`
	CPUUsage string `json:"cpu_usage"`
	MemUsage string `json:"mem_usage"`
	MinerCount string `json:"miner_count"`
	NodeCount string `json:"node_count"`
	NodeGraffi string `json:"node_graffi"`
	NodeHeight string `json:"node_height"`
	NodeVersion string `json:"node_version"`
	UpdateTime string `json:"update_time"`
}


func (a *Collections) TableName() string {
	return "iron_miner"
}

// func (a *Collections) FillData() {
// 	collect := Collections{
// 		MachineId: 0,
// 		IPAddress: "",
// 		MachineTag: "",
// 		WorkStatus: "",
// 		CPUUsage: "",
// 		MemUsage: "",
// 		MinerCount: "",
// 		NodeCount: "",
// 		NodeGraffi: "",
// 		NodeHeight: "",
// 		NodeVersion: "",
// 		UpdateTime: "",
// 	}
// 	Db.Save(&collect)
// }

type Miner struct {
	BaseModle
	MinerId uint   `gorm:"primary_key;auto_increment"`
	Title     string `gorm:"size:100;comment:'标题'"`
	Desc      string `gorm:"size:100;comment:'描述'"`
	CoverImg  string `gorm:"size:100;comment:'封面图'"`
	Content   string `gorm:"type:text;comment:'文章内容'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type MinerIndexReq struct {
	Title     string `form:"title"`
	CreatedAt string `form:"created_at"`
}

type MinerReq struct {
	MinerId int    `form:"article_id"`
	Title     string `form:"title" label:"标题" binding:"required"`
	CoverImg  string `form:"cover_img" label:"封面图" binding:"required"`
	Content   string `form:"content" label:"文章详情" binding:"required"`
	Desc      string `form:"desc" label:"文章描述" binding:"required"`
}

