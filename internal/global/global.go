package global

import (
	"github.com/emicklei/go-restful/v3"
	"gorm.io/gorm"
)

var (
	Mysql            *gorm.DB
	DefaultContainer *restful.Container
)
