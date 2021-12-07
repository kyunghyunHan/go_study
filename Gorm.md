# Gorm

```
package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Product struct {
	gorm.Model
	Code   string
	Price  uint
	People string
}

type UserU struct {
	gorm.Model
	Code   string
	Price  uint
	People string
}
type Userkkk struct {
	gorm.Model
	Code   string
	Price  uint
	People string
	name   string
}

func main() {
	dsn := "admin:diary1234@tcp(teamdiary.ciszsojgfruj.ap-northeast-2.rds.amazonaws.com)/pl2h?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Db연결에 실패")
	}

	db.AutoMigrate(&Product{})
	db.AutoMigrate(&UserU{})
	db.AutoMigrate(&Userkkk{})
}

```
