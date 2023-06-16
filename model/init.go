package model

import (
	"log"
	"os"
	"time"

	tlog "TikTokServer/pkg/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func InitDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Warn, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
	dsn := "niku:123@tcp(127.0.0.1:3307)/TikTokDB?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		tlog.Error(err.Error())
	}

	if err := DB.AutoMigrate(); err != nil {
		tlog.Error(err.Error())
	}

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := DB.DB()
	if err != nil {
		tlog.Error(err.Error())
	}
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	tlog.Info("DB init success", tlog.String("username", "niku"), tlog.String("password", "123"))
}

func GetDB() *gorm.DB {
	return DB
}

// gorm 1.20 版本之后使用连接池，无需显示关闭连接
// func CloseDB() {
// 	DB.Close()
// }
