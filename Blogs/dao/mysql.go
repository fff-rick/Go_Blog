package dao

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	dataSourceName := fmt.Sprintf("root:5923@tcp(localhost:3306)/goblog?charset=utf8&loc=%s&parseTime=true",
		url.QueryEscape("Asia/Shanghai"))
	
	// 配置 GORM 日志
	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略记录未找到错误
			Colorful:                  true,        // 彩色日志
		},
	)
	
	// 连接数据库
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Println("数据库连接错误：", err)
		panic(err)
	}
	
	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("获取数据库连接池错误：", err)
		panic(err)
	}
	
	//最大空闲连接数
	sqlDB.SetMaxIdleConns(5)
	//最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 连接最大存活时间
	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	//空闲连接最大存活时间
	sqlDB.SetConnMaxIdleTime(time.Minute * 1)
	
	DB = db
	log.Println("数据库连接成功")
}
