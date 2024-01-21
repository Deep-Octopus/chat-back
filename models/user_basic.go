package models

import (
	"errors"
	"go-chat/utils"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string    `json:"name"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	Phone         string    `valid:"matches(^1[3-9]{1}\\d{9}$)" json:"phone"`
	Email         string    `valid:"email" json:"email"`
	Identity      string    `json:"identity"`
	ClintIp       string    `json:"clintIp"`
	ClientPort    string    `json:"clientPort"`
	LoginTime     time.Time `json:"loginTime"`
	HeartbeatTime time.Time `json:"heartbeatTime"`
	LogoutTime    time.Time `json:"logoutTime"`
	IsLogout      bool      `json:"isLogout"`
	DeviceInfo    string    `json:"deviceInfo"`
	Desc          string    `json:"desc"` //预留字段
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	userList := make([]*UserBasic, 0)
	utils.DB.Find(&userList)
	return userList
}
func TakeUserById(id uint) *UserBasic {
	var user UserBasic
	utils.DB.Where("id = ?", id).Take(&user)
	return &user
}
func TakeGroupById(id uint) *GroupBasic {
	var group GroupBasic
	utils.DB.Where("id = ?", id).Take(&group)
	return &group
}
func TakeUserByName(name string) *UserBasic {
	var user UserBasic
	utils.DB.Where("name = ?", name).Take(&user)
	return &user
}
func TakeUserByUsername(username string) *UserBasic {
	var user UserBasic
	utils.DB.Where("username = ?", username).Take(&user)
	return &user
}
func TakeUserByEmail(email string) *UserBasic {
	var user UserBasic
	utils.DB.Where("email = ?", email).Take(&user)
	return &user
}
func TakeUserByPhone(phone string) *UserBasic {
	var user UserBasic
	utils.DB.Where("phone = ?", phone).Take(&user)
	return &user
}
func CreateUser(user *UserBasic) error {
	return utils.DB.Create(user).Error
}

func DeleteUser(user *UserBasic) error {
	return utils.DB.Delete(user).Error
}

func UpdateUser(user *UserBasic) error {
	oldUser := TakeUserById(user.ID)
	if oldUser == nil {
		return errors.New("用户不存在")
	}
	return utils.DB.Model(&oldUser).Updates(user).Error
}

func GetTokenByUserName(username string) string {
	return TakeUserByUsername(username).Identity
}
