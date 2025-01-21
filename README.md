# MyTodoList

## 项目依赖：
  ``` 
  # 初始化go模块
  go mod init MyTodoList
  # 安装gin框架
  go get -u github.com/gin-gonic/gin
  # 安装gorm
  go get -u gorm.io/gorm
  go get -u gorm.io/driver/mysql
  # 安装jwt用于用户身份验证
  go get -u github.com/golang-jwt/jwt/v5
  # 安装bcrypt用于密码加密
  go get -u golang.org/x/crypto/bcrypt
```

## 项目架构：
`````
- config
    - config.go         # 数据库初始化
    - jwt.go            # jwt配置和方法
- controllers
    - userController.go # 用户注册和登录控制器
    - todoController.go # 待办事项的增删查改控制器
- middlewares
    - authMiddleware.go # jwt鉴权中间件
- models
    - user.go           # 用户模型
    - todo.go           # 待办事项模型
- routes
  - routes.go         # 路由配置
- utils
  - hash.go           # 哈希校验
  - jwt.go            # jwt鉴权
- main.go             # 项目入口
- go.mod              # 模块文件
```

# 注意大小写、命名规范、apifox、数据库查看软件、虚拟机、docker
# navicat