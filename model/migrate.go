package model

import "log"

func migration() {
	// 自动迁移模式
	err := DB.Set("gorm:table_option", "charset=utf8mb4").
		AutoMigrate(&User{}).
		AutoMigrate(&Task{}).Error
	if err != nil {
		log.Fatalf("自动迁移失败: %v", err)
		return
	}

	log.Println("自动迁移成功")

	// 添加外键约束
	err = DB.Model(&Task{}).AddForeignKey("uid", "User(id)", "CASCADE", "CASCADE").Error
	if err != nil {
		log.Fatalf("添加外键失败: %v", err)
		return
	}

	log.Println("外键添加成功")
}
