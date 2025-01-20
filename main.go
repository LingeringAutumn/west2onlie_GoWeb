package main

import (
	"MyTodoList/config"
	"MyTodoList/models"
	"MyTodoList/routes"
	"fmt"
)

// 主入口
func main() {
	// 初始化数据库（自动创建数据库）
	config.InitDB()
	// 自动迁移表
	err := config.DB.AutoMigrate(&models.User{}, &models.Todo{})
	if err != nil {
		fmt.Println("数据库迁移失败:", err)
		return
	}

	fmt.Println("数据库迁移完成！")
	r := routes.SetupRouter()
	r.Run(":8080")
}
