# MyTodoLIst

## 项目依赖：
  ``` 
  # 初始化 Go 模块（如果未初始化）
  go mod init MyTodoList
  # 安装 Gin 框架（Web 框架）
  go get -u github.com/gin-gonic/gin
  # 安装 GORM（ORM框架）
  go get -u gorm.io/gorm
  go get -u gorm.io/driver/mysql
  # 安装 JWT（用于用户身份验证）
  go get -u github.com/golang-jwt/jwt/v5
  # 安装 bcrypt（用于密码加密）
  go get -u golang.org/x/crypto/bcrypt
```

## 项目架构：
- config
    - config.go         # 数据库初始化
    - jwt.go            # JWT 配置和方法
- controllers
    - userController.go # 用户注册和登录控制器
    - todoController.go # 待办事项的增删查改控制器
- middlewares
    - authMiddleware.go # JWT 鉴权中间件
- models
    - user.go           # 用户模型
    - todo.go           # 待办事项模型
- routes
  - routes.go         # 路由配置
- main.go             # 项目入口文件
- go.mod              # Go 模块文件
