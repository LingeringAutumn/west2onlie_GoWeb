package models

import (
	"MyTodoList/config"
	"fmt"
)

type Todo struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Status    int    `json:"status"` // 0: 未完成, 1: 已完成
	CreatedAt int64  `json:"createdAt"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
	UserID    uint   `json:"userId"`
}

func CreateTodo(todo *Todo) error {
	return config.DB.Create(todo).Error
}

func GetTodosByUser(userID uint, status int, keyword string, offset int, limit int) ([]Todo, int64, error) {
	var todos []Todo
	var count int64
	query := config.DB.Where("user_id = ?", userID)
	if status >= 0 {
		query = query.Where("status = ?", status)
	}
	if keyword != "" {
		query = query.Where("title LIKE ?", "%"+keyword+"%")
	}
	err := query.Offset(offset).Limit(limit).Find(&todos).Error
	query.Count(&count)
	return todos, count, err
}

func UpdateTodoStatus(id uint, status int) error {
	return config.DB.Model(&Todo{}).Where("id = ?", id).Update("status", status).Error
}

/*func DeleteTodo(id uint) error {
	return config.DB.Where("id = ?", id).Delete(&Todo{}).Error
}*/

func DeleteTodo(id uint) error {
	result := config.DB.Where("id = ?", id).Delete(&Todo{})
	if result.RowsAffected == 0 {
		return fmt.Errorf("事项不存在")
	}
	return result.Error
}
#