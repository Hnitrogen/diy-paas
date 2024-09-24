package routers

import (
	"diy-paas/internal/global"
	"github.com/gin-gonic/gin"
)

func RegisterBaseRoutes() {

}

func InitRootRouter() {
	r := gin.Default()
	global.Router = r
	return
}
