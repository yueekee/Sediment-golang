package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TestTable struct {
	ID        int       `json:"id"`
	StartDate time.Time `json:"start_date"`
}

func (TestTable) TableName() string {
	return "test_table"
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/test_data?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	var t TestTable
	if err = db.Find(&t).Error; err != nil {
		log.Println("find TestTable err:", err)
	}
	fmt.Printf("t:%+v\n", t)
}
