package lib

import (
	"fmt"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

func init() {
	DB = initGormDB()
}

func initGormDB() *gorm.DB {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local", "root", "root", "127.0.0.1", 3306, "qujiedan")
	orm, err := gorm.Open(mysql.Open(dns), &gorm.Config{}) //
	if err != nil {
		log.Fatalf("连接mysql数据库失败:%v\n", err)
	}
	mysqlDB, _ := orm.DB()
	mysqlDB.SetConnMaxLifetime(30 * time.Second)
	mysqlDB.SetMaxIdleConns(5)
	mysqlDB.SetMaxOpenConns(10)

	return orm
}
