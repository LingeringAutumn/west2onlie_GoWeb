package controllers

import (
	"MyTodoList/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 创建待办事项
func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "无效的请求数据",
		})
		return
	}
	userID, _ := c.Get("userID") // 从JWT中获取用户ID
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

// 获取待办事项（支持分页、关键词搜索、状态筛选）
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

// 更新待办事项状态
func UpdateTodoStatus(c *gin.Context) {
	var req struct {
		ID     uint `json:"id"`
		Status int  `json:"status"`
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

// 删除待办事项
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
