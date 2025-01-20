package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB 初始化数据库
func InitDB() {
	// 加载.env文件
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("加载 .env 文件失败")
	}

	// 从环境变量获取数据库配置
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// 连接到Mysql服务器（不指定数据库）
	dsnWithoutDB := fmt.Sprintf("%s:%s@tcp(%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost)

	tempDB, err := gorm.Open(mysql.Open(dsnWithoutDB), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接 MySQL 服务器失败: %v", err)
	}

	// 创建数据库（如果不存在）
	tempDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))

	// 连接到创建后的数据库
	dsnWithDB := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbName)

	DB, err = gorm.Open(mysql.Open(dsnWithDB), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	fmt.Println("数据库连接成功")
}
