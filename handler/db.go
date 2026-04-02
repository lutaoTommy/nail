package handler

import (
	"fmt"
	"sync"
	"time"
	"nail/config"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
)

/*mysql*/
var mysqlSession *gorm.DB
var mysqlOnce sync.Once

/*获取单例数据库对象*/
func getMysqlConn() *gorm.DB {
	mysqlOnce.Do(createMysql)
	return mysqlSession
}

/*创建数据库连接*/
func createMysql() {
	var err error
	// charset=utf8mb4 支持完整 Unicode；timeout=10s 连接超时；parseTime 解析时间为 time.Time；loc=Local 时区
	dsn := fmt.Sprintf("%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s", config.GetMysqlUrl())
	mysqlSession, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := mysqlSession.DB()
	if err != nil {
		panic(err)
	}
	// 连接池：最大打开连接数，避免占满 MySQL max_connections
	sqlDB.SetMaxOpenConns(25)
	// 最大空闲连接数，减少建连开销
	sqlDB.SetMaxIdleConns(10)
	// 空闲连接最大存活时间，避免长时间空闲连接被服务端关闭
	sqlDB.SetConnMaxLifetime(time.Hour)
}