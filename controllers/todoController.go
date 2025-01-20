package controllers

import (
	"MyTodoList/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateTodo 创建待办事项
func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "无效的请求数据",
		})
		return
	}
	userID, _ := c.Get("userID") // 从jwt中获取用户id
	todo.UserID = userID.(uint)

	if err := models.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"msg":    "创建失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "创建成功",
		"data":   todo,
	})
}

// GetTodos 获取待办事项：支持分页、关键词搜索、状态筛选
func GetTodos(c *gin.Context) {
	var req struct {
		Status  int    `json:"status"`
		Keyword string `json:"keyword"`
		Page    int    `json:"page"`
		Limit   int    `json:"limit"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "请求参数错误",
		})
		return
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 {
		req.Limit = 10
	}

	userID, _ := c.Get("userID")
	todos, total, err := models.GetTodosByUser(userID.(uint), req.Status, req.Keyword, (req.Page-1)*req.Limit, req.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"msg":    "查询失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "查询成功",
		"data": gin.H{
			"items": todos,
			"total": total,
		},
	})
}

// UpdateTodoStatus 更新单个待办事项状态
func UpdateTodoStatus(c *gin.Context) {
	var req struct {
		ID     uint `json:"id"`
		Status int  `json:"status"` // 0: 未完成, 1: 已完成
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "无效的请求数据",
		})
		return
	}

	err := models.UpdateTodoStatus(req.ID, req.Status)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"msg":    "更新失败，事项不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "事项已标记为完成",
	})
}

// UpdateAllTodoStatus 批量更新待办事项状态
func UpdateAllTodoStatus(c *gin.Context) {
	var req struct {
		TargetStatus  int `json:"targetStatus"` // 0: 未完成, 1: 已完成
		CurrentStatus int `json:"currentStatus"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "无效的请求数据",
		})
		return
	}

	err := models.UpdateAllTodosStatus(req.CurrentStatus, req.TargetStatus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"msg":    "批量更新失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "所有事项已更新",
	})
}

// DeleteTodo 删除单个待办事项
func DeleteTodo(c *gin.Context) {
	var req struct {
		ID uint `json:"id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "无效的请求数据",
		})
		return
	}

	err := models.DeleteTodo(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"msg":    "删除失败，事项不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "事项删除成功",
	})
}

// DeleteTodosByStatus 批量删除待办事项
func DeleteTodosByStatus(c *gin.Context) {
	var req struct {
		Status int `json:"status"` // 0: 删除未完成, 1: 删除已完成, 2: 删除所有
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "无效的请求数据",
		})
		return
	}

	err := models.DeleteTodosByStatus(req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"msg":    "批量删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "批量事项已删除",
	})
}
