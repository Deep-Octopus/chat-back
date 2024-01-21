package service

import "go-chat/models"

func GetMessagesByContact(contact models.Contact) []models.Message {
	msgs := make([]models.Message, 0)
	//发送的消息
	msgs = append(msgs, models.GetMessagesByFromIdAndTargetIdAndType(contact.OwnerId, contact.TargetId, uint(contact.Type))...)
	//收到的消息
	msgs = append(msgs, models.GetMessagesByFromIdAndTargetIdAndType(contact.TargetId, contact.OwnerId, uint(contact.Type))...)
	return msgs
}
