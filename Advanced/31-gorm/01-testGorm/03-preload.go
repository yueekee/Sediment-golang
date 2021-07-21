package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username  string
	ProjectID int     `json:"project_id"`
	//Orders1    []Order	`gorm:"ForeignKey:Pid;References:ProjectID"`
	Orders1    []Order	`gorm:"References:ProjectID"`
}
// foreignKey	指定当前模型的列作为连接表的外键
// references	指定引用表的列名，其将被映射为连接表外键

//joinForeignKey	指定连接表的外键列名，其将被映射到当前表
//joinReferences	指定连接表的外键列名，其将被映射到引用表

type Order struct {
	PID    int `json:"pid" gorm:"joinForeignKey:Pid"`
	UserID uint
	Price  float64
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/test_data?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	var users []User
	db.Debug().Preload("Orders1").Find(&users)
	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL
	// SELECT * FROM `orders` WHERE `orders`.`user_id` IN (1,2)

	//db.Debug().Table("users").Preload("Orders.ProjectID",  func(db *gorm.DB) *gorm.DB {
	//	return db.Where("users.project_id = orders.pid")
	//}).Scan(&users)

	marshal, _ := json.Marshal(users)
	fmt.Printf("%s\n", string(marshal))
}
