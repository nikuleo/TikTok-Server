package model

import (
	"fmt"
	"log"
	"os"
	"time"

	"TikTokServer/pkg/config"
	"TikTokServer/pkg/tlog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB
)

func InitDB() {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Millisecond * 100, // Slow SQL threshold
			LogLevel:                  logger.Info,            // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,                   // Don't include params in the SQL log
			// Colorful:                  false,                  // Disable color
		},
	)

	cfg := config.GetConfig("dbConfig")
	viper := cfg.Viper
	// dsn := "niku:123@tcp(127.0.0.1:3307)/TikTokDB?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
		viper.GetString("mysql.charset"),
		viper.GetBool("mysql.parseTime"),
		viper.GetString("mysql.loc"),
	)
	// 注意使用 = 赋值，而不是 := 赋值，否则会覆盖全局变量 db，全局 db 仍然为 nil
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		tlog.Error(err.Error())
	}

	// 自动迁移, 若模型在数据库中不存在，则会自动创建对应的表，若已存在，则会检查字段是否发生变化，若发生变化，则会修改表结构
	if err := db.AutoMigrate(&User{}, &Video{}, &Comment{}, &Relation{}, &Message{}); err != nil {
		tlog.Error(err.Error())
	}

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()
	if err != nil {
		tlog.Error(err.Error())
	}
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(viper.GetInt("mysql.maxIdleConns"))

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(viper.GetInt("mysql.maxOpenConns"))

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	tlog.Infof("DB init success, user: %s", viper.GetString("mysql.user"))
}

// func GetDB() *gorm.DB {
// 	return db
// }

// gorm 1.20 版本之后使用连接池，无需显示关闭连接
// func CloseDB() {
// 	DB.Close()
// }
