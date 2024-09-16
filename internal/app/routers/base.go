package routers

import (
	"diy-paas/internal/app/api/base"
	"diy-paas/internal/global"
	"github.com/emicklei/go-restful/v3"
)

func RegisterBaseRoutes() {
	ws := new(restful.WebService)

	ws.Route(ws.GET("/hello").To(base.Hello))
	// Handle User
	ws.Route(ws.GET("/users/{user-id}").To(base.UserResource{}.FindUser))
	ws.Route(ws.POST("/users/").To(base.CreateUser))
	ws.Route(ws.POST("/user/login/").To(base.UserLogin))

	//restful.Add(ws)
	global.DefaultContainer.Add(ws)
}
