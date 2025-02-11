package service

import (
	"demo1/model"
	"demo1/pkg/utils"
	"demo1/serializer"

	"github.com/jinzhu/gorm"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16"`
}

func (service *UserService) Register() serializer.Response {
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).
		First(&user).Count(&count)
	if count == 1 {
		return serializer.Response{
			Status: 400,
			Msg:    "已经有这个人了，不用再注册",
		}
	}
	user.UserName = service.UserName
	//加密
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    err.Error(),
		}
	}
	//创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "数据库操作错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "用户注册成功",
	}
}

func (service *UserService) Login() serializer.Response {
	var user model.User
	//先去找一下这个user，看看数据库有没有这个人
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Msg:    "用户不存在，请先注册",
			}
		}
		//如果不是用户不存在，是其他不可抗拒的因素
		return serializer.Response{
			Status: 500,
			Msg:    "数据库错误" + err.Error(),
		}
	}
	if user.CheckPassword(service.Password) == false {
		return serializer.Response{
			Status: 400,
			Msg:    "密码错误",
		}
	}
	//发一个token，为了其他功能需要身份验证所给前端存储的
	//创建一个备忘录，需要token来确定是谁创建的
	token, err := utils.GenerateToken(user.ID, service.UserName, service.Password)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "token签发错误" + err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    "登录成功",
	}
}
