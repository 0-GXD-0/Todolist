package main

import (
	"demo1/conf"
	"demo1/router"
	"log"
)

func main() {
	conf.Init()
	r := router.NewRouter()

	// 启动服务器
	if err := r.Run(":3000"); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
