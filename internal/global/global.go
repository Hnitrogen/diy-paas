package global

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	Mysql            *gorm.DB
	DefaultContainer *restful.Container
	Router           *gin.Engine
)
