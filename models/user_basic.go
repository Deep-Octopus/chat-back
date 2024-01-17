package models

import (
	"errors"
	"go-chat/utils"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Identity      string
	ClintIp       string
	ClientPort    string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LogoutTime    time.Time
	IsLogout      bool
	DeviceInfo    string
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
func TakeUserByName(name string) *UserBasic {
	var user UserBasic
	utils.DB.Where("name = ?", name).Take(&user)
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

func GetTokenByUserName(name string) string {
	return TakeUserByName(name).Identity
}
