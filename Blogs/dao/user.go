package dao

import (
	"blogs/models"
	"log"
)

func GetUserNameByID(id int) string {
	var user models.User
	result := DB.Where("uid = ?", id).First(&user)
	if result.Error != nil {
		log.Println("获取用户名失败：", result.Error)
		return ""
	}
	return user.Username
}

func GetUser(username, password string) *models.User {
	var user models.User
	result := DB.Where("user_name = ? AND passwd = ?", username, password).First(&user)
	if result.Error != nil {
		log.Println("获取用户信息失败：", result.Error)
		return nil
	}
	return &user
}
