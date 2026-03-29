package service

import (
	"blogs/dao"
	models "blogs/models"
	"blogs/utils"
	"errors"
)

func Login(username, password string) (*models.LoginResponse, error) {
	password = utils.Md5Crypt(password)
	//fmt.Println("MD5: ", password)
	user := dao.GetUser(username, password)
	if user == nil {
		return nil, errors.New("用户名或密码错误")
	}
	token, err := utils.Award(&user.Uid)
	if err != nil {
		return nil, errors.New("获取Token失败")
	}
	lr := &models.LoginResponse{
		Token: token,
		UserInfo: models.UserInfo{
			Avatar:   user.Avatar,
			Uid:      user.Uid,
			Username: user.Username,
		},
	}
	return lr, nil
}
