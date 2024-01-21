package models

import (
	"go-chat/utils"
	"gorm.io/gorm"
	"time"
)

// Contact 人员关系
type Contact struct {
	gorm.Model
	OwnerId  uint   // 谁的关系信息
	TargetId uint   // 对应的谁
	Type     int    // 对应的类型 1:好友 2：群 3
	Nickname string // 备注
	Desc     string // 预留
}

// ListMessageDto 消息列表
type ListMessageDto struct {
	ID             uint      `json:"ID"`
	AvatarUrl      string    `json:"avatarUrl"`
	Name           string    `json:"name"`
	Username       string    `json:"username"`
	Online         bool      `json:"online"`
	IsGroup        bool      `json:"isGroup"`
	Num            uint      `json:"num"`
	Abstract       string    `json:"abstract"`
	Datetime       time.Time `json:"datetime"`
	NumOfUnReadMsg uint      `json:"numOfUnReadMsg"`
}

func SearchFriend(userId uint) []UserBasic {
	contacts := make([]Contact, 0)
	users := make([]UserBasic, 0)
	utils.DB.Where("owner_id = ? and type = 1", userId, userId).Find(&contacts)
	for _, c := range contacts {
		user := TakeUserById(c.TargetId)
		user.Password = "******"
		users = append(users, *user)
	}
	return users
}
func SearchFriendAndLastMessage(userId uint) []ListMessageDto {
	contacts := make([]Contact, 0)
	listMsg := make([]ListMessageDto, 0)
	utils.DB.Where("owner_id = ? and type = 1", userId).Find(&contacts)
	for _, c := range contacts {
		user := TakeUserById(c.TargetId)
		//TODO
		//获取最后一条消息
		lastMsg := GetLastMessageByUserIdAndType(userId, user.ID, 1)
		if lastMsg.ID == 0 {
			lastMsg.Content = ""
			lastMsg.CreatedAt = time.Now()
		}
		//找到未读消息条数
		//其他信息
		tmp := &ListMessageDto{
			ID:             user.ID,
			AvatarUrl:      "https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png",
			Name:           c.Nickname,
			Username:       user.Username,
			IsGroup:        false,
			Online:         IsOnline(user.ID), //TODO 后期去map里面判断是否在线
			Abstract:       lastMsg.Content,
			Datetime:       lastMsg.CreatedAt,
			NumOfUnReadMsg: uint(GetNumOfUnreadMessageByUserId(user.ID, userId, 1)),
		}
		if len(c.Nickname) == 0 {
			tmp.Name = user.Name
		}
		listMsg = append(listMsg, *tmp)
	}
	return listMsg
}
func SearchGroup(userId uint) []GroupBasic {
	contacts := make([]Contact, 0)
	groups := make([]GroupBasic, 0)
	utils.DB.Where("owner_id = ? and type = 2", userId).Find(&contacts)
	for _, c := range contacts {
		user := TakeGroupById(c.TargetId)
		groups = append(groups, *user)
	}
	return groups
}
func SearchGroupAndLastMessage(userId uint) []ListMessageDto {
	contacts := make([]Contact, 0)
	listMsg := make([]ListMessageDto, 0)
	utils.DB.Where("owner_id = ? and type = 2", userId).Find(&contacts)
	for _, c := range contacts {
		group := TakeGroupById(c.TargetId)
		//TODO
		//获取最后一条消息
		//其他信息
		tmp := ListMessageDto{
			ID:             group.ID,
			AvatarUrl:      group.Icon,
			Name:           c.Nickname,
			Username:       group.Name,
			IsGroup:        true,
			Abstract:       "摘要内容",
			Datetime:       time.Now(),
			NumOfUnReadMsg: 10,
		}
		if len(c.Nickname) == 0 {
			tmp.Name = group.Name
		}
		listMsg = append(listMsg, tmp)
	}
	return listMsg
}
