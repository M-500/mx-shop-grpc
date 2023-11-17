package utils

// Package utils
// Date        : 2023/2/15 21:18
// Version     : 1.0.0
// Author      : 代码小学生王木木
// Email       : 18574945291@163.com
// Description :

import (
	"database/sql"
	"github.com/glebarez/sqlite" // 纯 Go 实现的 SQLite 驱动, 详情参考： https://github.com/glebarez/sqlite
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	db    *gorm.DB
	sqlDB *sql.DB
)

func OpenDB(dbType, dsn string, config *gorm.Config, maxIdleConns, maxOpenConns int, models ...interface{}) (err error) {
	if config == nil {
		config = &gorm.Config{}
	}

	if config.NamingStrategy == nil {
		config.NamingStrategy = schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		}
	}
	if dbType == "SQLite" {
		if db, err = gorm.Open(sqlite.Open(dsn), config); err != nil {
			print("opens database failed: %s", err.Error())
			return
		}
	} else {
		if db, err = gorm.Open(mysql.Open(dsn), config); err != nil {
			print("opens database failed: %s", err.Error())
			return
		}
	}

	if sqlDB, err = db.DB(); err == nil {
		sqlDB.SetMaxIdleConns(maxIdleConns)
		sqlDB.SetMaxOpenConns(maxOpenConns)
	} else {
		print(err)
	}

	if err = db.AutoMigrate(models...); nil != err {
		print("auto migrate tables failed: %s", err.Error())
	}
	return
}

// 获取数据库链接
func DB() *gorm.DB {
	return db
}

// 关闭连接
func CloseDB() {
	if sqlDB == nil {
		return
	}
	if err := sqlDB.Close(); nil != err {
		print("Disconnect from database failed: %s", err.Error())
	}
}
