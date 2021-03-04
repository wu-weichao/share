package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"share/configs"
)

func GetDB(c *configs.DatabaseConfig) (db *gorm.DB, err error) {
	cfg := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
	switch c.Type {
	case "mysql":
		// user:pass@tcp(ip:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			c.User,
			c.Password,
			c.Host,
			c.Port,
			c.DBName,
			c.Charset)
		db, err = gorm.Open(mysql.Open(dsn), cfg)
	case "pgsql":
		// host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			c.Host,
			c.User,
			c.Password,
			c.DBName,
			c.Port)
		db, err = gorm.Open(postgres.Open(dsn), cfg)
	}
	return
}
