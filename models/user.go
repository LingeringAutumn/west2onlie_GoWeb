package models

import "MyTodoList/config"

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
}

// CreateUser 创建用户
func CreateUser(user *User) error {
	return config.DB.Create(user).Error
}

// GetUserByUsername 按用户获取信息
func GetUserByUsername(username string) (*User, error) {
	var user User
	err := config.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}
