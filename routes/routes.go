package routes

import (
	"MyTodoList/controllers"
	"MyTodoList/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 用户相关接口
	userGroup := r.Group("/todolist/user")
	{
		userGroup.POST("/register", controllers.Register)
		userGroup.POST("/login", controllers.Login)
	}

	// 事务相关接口
	todoGroup := r.Group("/todolist/todos").Use(middlewares.AuthMiddleware())
	{
		todoGroup.POST("/", controllers.CreateTodo)  // 创建待办事项
		todoGroup.POST("/get", controllers.GetTodos) // 查询待办事项（POST接收JSON）
		todoGroup.PUT("/updateStatus", controllers.UpdateTodoStatus)
		todoGroup.PUT("/updateAllStatus", controllers.UpdateAllTodoStatus)
		todoGroup.DELETE("/delete", controllers.DeleteTodo)
		todoGroup.DELETE("/deleteByStatus", controllers.DeleteTodosByStatus)
	}

	return r
}
