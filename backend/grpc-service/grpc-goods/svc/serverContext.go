package svc

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/17 12:58
//
import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"grpc-user/config"
	"grpc-user/model"
	"grpc-user/pkg/utils"
	"log"
	"os"
	"time"
)

var svcContext *ServiceContext

type ServiceContext struct {
	Config    *config.Config
	MysqlConn *gorm.DB
}

func NewSrvCtx(path string) *ServiceContext {
	cfg := config.NewConfig(path)
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, // 慢查询阈值
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false, // 禁止彩色打印
		},
	)
	gormConf := &gorm.Config{
		Logger: gormLogger,
	}
	err := utils.OpenDB("", cfg.MySQl.Datasource, gormConf, 10, 20, model.ModelList...)
	if err != nil {
		panic(any(err))
		return nil
	}
	svcContext = &ServiceContext{
		Config:    cfg,
		MysqlConn: utils.DB(),
	}
	return svcContext
}

func GetSrvCtx() *ServiceContext {
	return svcContext
}
