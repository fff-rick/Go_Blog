package dao

import (
	"blogs/models"
	"log"
)

func GetUserNameByID(id int) string {
	sql := "SELECT user_name FROM user WHERE uid = ?"
	row := DB.QueryRow(sql, id)
	var username string
	row.Scan(&username)
	return username
}

func GetUser(username, password string) *models.User {
	sql := "SELECT * FROM user WHERE user_name = ? AND passwd = ?"
	//fmt.Println("check: ", username, password)
	row := DB.QueryRow(sql, username, password)
	if row.Err() != nil {
		log.Println(row.Err())
		return nil
	}
	user := &models.User{}
	err := row.Scan(&user.Uid, &user.Username, &user.Password, &user.Avatar, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Println("获取用户信息失败：", err)
		return nil
	}
	return user
}
