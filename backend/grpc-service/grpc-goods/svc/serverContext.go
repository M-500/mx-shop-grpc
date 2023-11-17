package svc

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/17 12:58
//
import (
	"gorm.io/gorm"
	"grpc-user/config"
)

var svcContext *ServiceContext

type ServiceContext struct {
	Config    *config.Config
	MysqlConn *gorm.DB
}

func NewSrvCfg() *ServiceContext {
	return &ServiceContext{}
}

func GetSrvConfig() *ServiceContext {
	return svcContext
}
