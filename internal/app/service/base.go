package service

import (
	"diy-paas/internal/app/dto"
	"diy-paas/internal/app/model"
	"diy-paas/internal/global"
	"diy-paas/pkg/mysql"
	"errors"
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"gorm.io/gorm"
)

func CreateUserService(request *restful.Request) (err error, result *gorm.DB) {
	var udto dto.UserCreateInfo
	e := request.ReadEntity(&udto)
	if e != nil {
		return errors.New("序列化失败"), nil
	}

	// repo
	user := model.User{UserName: udto.UserName, Password: udto.Password, Email: udto.Email}
	result = global.Mysql.Create(&user)

	if result.Error != nil {
		if mysql.IsDuplicateEntry(result.Error.Error()) {
			return errors.New("用户名已经存在"), nil
		} else {
			return err, nil
		}
	}
	return nil, result
}

func UserLoginService(request *restful.Request) (err error, user model.User) {
	var input dto.UserLogin
	e := request.ReadEntity(&input)
	if e != nil {
		return errors.New("序列化失败"), model.User{}
	}

	// 查用户
	target := model.User{}
	result := global.Mysql.Model(&model.User{}).Select("email").Select("user_name").
		Where("user_name = ? AND password = ?", input.UserName, input.Password).First(&target)
	if result.Error != nil {
		fmt.Println(result.Statement)
		return result.Error, model.User{}
	}

	return nil, target
}
