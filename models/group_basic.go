package models

import "gorm.io/gorm"

// GroupBasic 群聊
type GroupBasic struct {
	gorm.Model
	Name    string //群聊名称
	OwnerId uint   //群聊拥有者
	Icon    string //图标
	Type    int
	Desc    string // 预留
}
