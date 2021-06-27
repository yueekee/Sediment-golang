package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



func main() {
	dsn := "host=pgm-2ze3dh0r3p3g1e66127770.pg.rds.aliyuncs.com user=gemdale password=1q2w#E$R dbname=gemdale port=1921 sslmode=disable TimeZone=Asia/Shanghai"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//如果出错就GameOver了
	if err != nil {
		panic(err)
	}

	//db.Table("t_follow_up_record").Find("sid = 284942812187725824 and del = f and state = 1")
	fmt.Println("ok")
}
