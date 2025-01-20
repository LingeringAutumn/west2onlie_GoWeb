package models

import (
	"MyTodoList/config"
	"fmt"
	"gorm.io/gorm"
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
	View      int    `json:"view"` // 访问次数
}

// CreateTodo 创建事务
func CreateTodo(todo *Todo) error {
	return config.DB.Create(todo).Error
}

// GetTodosByUser 查询事务
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

	// 更新访问次数
	for i := range todos {
		config.DB.Model(&todos[i]).UpdateColumn("view", gorm.Expr("view + ?", 1))
	}

	return todos, count, err
}

// UpdateTodoStatus 更新单个待办事项状态
func UpdateTodoStatus(id uint, status int) error {
	result := config.DB.Model(&Todo{}).Where("id = ?", id).Update("status", status)
	if result.Error != nil {
		return fmt.Errorf("更新失败: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("事项不存在")
	}
	return nil
}

// UpdateAllTodosStatus 批量更新事项状态
func UpdateAllTodosStatus(currentStatus, targetStatus int) error {
	result := config.DB.Model(&Todo{}).Where("status = ?", currentStatus).Update("status", targetStatus)
	if result.Error != nil {
		return fmt.Errorf("批量更新失败: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("未找到符合条件的事项")
	}
	return nil
}

// DeleteTodo 删除单个待办事项
func DeleteTodo(id uint) error {
	result := config.DB.Delete(&Todo{}, id)
	if result.Error != nil {
		return fmt.Errorf("删除失败: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("事项不存在")
	}
	return nil
}

// DeleteTodosByStatus 批量删除待办事项
func DeleteTodosByStatus(status int) error {
	var result *gorm.DB
	if status == 2 {
		result = config.DB.Delete(&Todo{})
	} else {
		result = config.DB.Where("status = ?", status).Delete(&Todo{})
	}
	if result.Error != nil {
		return fmt.Errorf("批量删除失败: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("未找到符合条件的事项")
	}
	return nil
}
