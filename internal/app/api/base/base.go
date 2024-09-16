package base

import (
	"diy-paas/internal/app/dto"
	"diy-paas/internal/app/model"
	"diy-paas/internal/app/service"
	"diy-paas/pkg/jwt"
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"io"
	"log"
	"net/http"
)

func Hello(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "Hello, World!")
}

type UserResource struct {
	users map[string]model.User
}

// GET http://localhost:8080/users/1
func (u UserResource) FindUser(request *restful.Request, response *restful.Response) {
	log.Println("findUser")
	id := request.PathParameter("user-id")
	fmt.Println(id)
}

func CreateUser(request *restful.Request, response *restful.Response) {
	log.Println("createUser")
	err, _ := service.CreateUserService(request)
	if err != nil {
		dto.Response(response, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.Response(response, http.StatusCreated, "创建成功", nil)
}

func UserLogin(request *restful.Request, response *restful.Response) {
	log.Println("Userlogin")
	err, user := service.UserLoginService(request)
	if err != nil {
		dto.Response(response, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	// 生成jwt令牌返回
	token, _ := jwt.GenerateJWT(user.UserName)
	dto.Response(response, http.StatusOK, "登录成功", dto.UserToken{
		UserName: user.UserName, Email: user.Email, Token: token,
	})
}
