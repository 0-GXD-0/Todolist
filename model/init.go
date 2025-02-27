package model

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Database(connstring string) {
	db, err := gorm.Open("mysql", connstring)
	if err != nil {
		fmt.Println(err)
		panic("mysql数据库连接错误")
	}
	fmt.Println("数据库连接成功")

	db.LogMode(true)

	if gin.Mode() == "release" {
		db.LogMode(false)
	}

	db.SingularTable(true)                       //表名不加s
	db.DB().SetMaxIdleConns(20)                  //设置连接池
	db.DB().SetMaxOpenConns(100)                 //最大连接数
	db.DB().SetConnMaxLifetime(time.Second * 30) //连接时间

	DB = db
	migration()
}
