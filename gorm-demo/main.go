package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	Age  uint8
}

func main() {
	//连接MySQL
	dsn := "root:Huawei12#$@tcp(152.136.176.118:3306)/go-demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 创建一条记录
	db.Create(&User{Name: "zhanshan", Age: 20})
	// 查询
	var user []User
	db.Where("name = ?", "zhanshan").Find(&user)
	fmt.Printf("%v\n", user)

}
