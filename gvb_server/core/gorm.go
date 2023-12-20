package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvb_server/global"
	"log"
	"time"
)

func InitGrom() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		//global.LOG.Warn("Mysql is not configured, cancel grom connect")
		log.Println("未配置mysql, 取消gorm连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()

	var mysqlLogger logger.Interface
	if global.Config.System.Env == "dev" {
		// dev : all sql log
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}
	//global.MysqlLog = logger.Default.LogMode(logger.Info)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		log.Fatalf(fmt.Sprintf("[%s] Failed to connect to mysql", dsn))
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)               // max idle connections 10
	sqlDB.SetMaxOpenConns(100)              // max open connections 100
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 连接最大复用时间, 不能超过MySQL的wait_timeout
	return db
}
