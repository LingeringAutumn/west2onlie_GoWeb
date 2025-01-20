# 全局公共参数

**全局Header参数**

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| 暂无参数 |

**全局Query参数**

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| 暂无参数 |

**全局Body参数**

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| 暂无参数 |

**全局认证方式**

> 无需认证

# 状态码说明

| 状态码 | 中文描述 |
| --- | ---- |
| 暂无参数 |

# 注册测试

> 创建人: 比亚

> 更新人: 比亚

> 创建时间: 2025-01-20 18:15:26

> 更新时间: 2025-01-20 18:15:53

```text
暂无描述
```

**接口状态**

> 开发中

**接口URL**

> http://localhost:8080/api/user/register 

| 环境  | URL |
| --- | --- |


**请求方式**

> POST

**Content-Type**

> json

**请求Body参数**

```javascript
{
  "username": "admin1",
  "password": "admin1123456"
}

```

**认证方式**

> 继承父级

**响应示例**

* 成功(200)

```javascript
暂无数据
```

* 失败(404)

```javascript
暂无数据
```

# 成功登录测试

> 创建人: 比亚

> 更新人: 比亚

> 创建时间: 2025-01-20 18:16:56

> 更新时间: 2025-01-20 22:28:16

```text
暂无描述
```

**接口状态**

> 开发中

**接口URL**

> http://localhost:8080/todolist/user/login 

| 环境  | URL |
| --- | --- |


**请求方式**

> POST

**Content-Type**

> json

**请求Body参数**

```javascript
{
  "username": "admin4",
  "password": "admin4123456"
}

```

**认证方式**

> 继承父级

**响应示例**

* 成功(200)

```javascript
暂无数据
```

* 失败(404)

```javascript
暂无数据
```

# 创建事务

> 创建人: 比亚

> 更新人: 比亚

> 创建时间: 2025-01-20 18:17:45

> 更新时间: 2025-01-20 22:28:32

```text
暂无描述
```

**接口状态**

> 开发中

**接口URL**

> http://localhost:8080/todolist/todos/ 

| 环境  | URL |
| --- | --- |


**请求方式**

> POST

**Content-Type**

> json

**请求Header参数**

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| Authorization | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzc0NjQxMDYsInVzZXJJRCI6NH0.xkqVbzbrwnE196f8jt01iE13nXcZLa8gZV86BUTbnyg | String | 是 | - |

**请求Body参数**

```javascript
{
  "title": "事务测试标题14",
  "content": "事务测试内容14",
  "status": 0,
  "createdAt": 1700000000,
  "startTime": 1700000001,
  "endTime": 1700009999
}

```

**认证方式**

> 继承父级

**响应示例**

* 成功(200)

```javascript
暂无数据
```

* 失败(404)

```javascript
暂无数据
```

# 按关键字查询单个事务

> 创建人: 比亚

> 更新人: 比亚

> 创建时间: 2025-01-20 18:18:41

> 更新时间: 2025-01-20 22:28:36

```text
暂无描述
```

**接口状态**

> 开发中

**接口URL**

> http://localhost:8080/todolist/todos/get

| 环境  | URL |
| --- | --- |


**请求方式**

> POST

**Content-Type**

> json

**请求Header参数**

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| Authorization | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzc0NjQxMDYsInVzZXJJRCI6NH0.xkqVbzbrwnE196f8jt01iE13nXcZLa8gZV86BUTbnyg | String | 是 | - |

**请求Body参数**

```javascript
{
  "status": 0,
  "keyword": "事务",
  "page": 1,
  "limit": 5
}

```

**认证方式**

> 继承父级

**响应示例**

* 成功(200)

```javascript
暂无数据
```

* 失败(404)

```javascript
暂无数据
```

# 修改单个事务为已完成

> 创建人: 比亚

> 更新人: 比亚

> 创建时间: 2025-01-20 19:49:31

> 更新时间: 2025-01-20 22:28:40

```text
暂无描述
```

**接口状态**

> 开发中

**接口URL**

> http://localhost:8080/todolist/todos/updateStatus

| 环境  | URL |
| --- | --- |


**请求方式**

> PUT

**Content-Type**

> json

**请求Header参数**

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| Authorization | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzc0NjQxMDYsInVzZXJJRCI6NH0.xkqVbzbrwnE196f8jt01iE13nXcZLa8gZV86BUTbnyg | String | 是 | - |
| Content-Type | application/json | String | 是 | - |

**请求Body参数**

```javascript
{
  "id": 15,
  "status": 1
}

```

**认证方式**

> 继承父级

**响应示例**

* 成功(200)

```javascript
暂无数据
```

* 失败(404)

```javascript
暂无数据
```

# 删除单个任务

> 创建人: 比亚

> 更新人: 比亚

> 创建时间: 2025-01-20 19:55:42

> 更新时间: 2025-01-20 22:28:44

```text
暂无描述
```

**接口状态**

> 开发中

**接口URL**

> http://localhost:8080/todolist/todos/delete

| 环境  | URL |
| --- | --- |


**请求方式**

> DELETE

**Content-Type**

> json

**请求Header参数**

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| Authorization | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzc0NjAyODUsInVzZXJJRCI6MX0.XlSHsaInEBXuH1ucgI8E8ZL_WzRjH9ccWSf7FSMUpFs | String | 是 | - |

**请求Body参数**

```javascript
{
  "id": 999
}

```

**认证方式**

> 继承父级

**响应示例**

* 成功(200)

```javascript
暂无数据
```

* 失败(404)

```javascript
暂无数据
```

# 修改所有事务为已完成

> 创建人: 比亚

> 更新人: 比亚

> 创建时间: 2025-01-20 22:25:38

> 更新时间: 2025-01-20 22:28:47

```text
暂无描述
```

**接口状态**

> 开发中

**接口URL**

> http://localhost:8080/todolist/todos/updateAllStatus

| 环境  | URL |
| --- | --- |


**请求方式**

> PUT

**Content-Type**

> json

**请求Header参数**

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| Authorization | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzc0NjkzNjMsInVzZXJJRCI6NH0.IucBuozYs-Y-5i4tSL36qk6UuszkdwEAvbOqFk3sgXU | String | 是 | - |
| Content-Type | - | String | 是 | - |

**请求Body参数**

```javascript
{
  "currentStatus": 0,
  "targetStatus": 1
}

```

**认证方式**

> 继承父级

**响应示例**

* 成功(200)

```javascript
暂无数据
```

* 失败(404)

```javascript
暂无数据
```

# 删除所有任务

> 创建人: 比亚

> 更新人: 比亚

> 创建时间: 2025-01-20 22:27:52

> 更新时间: 2025-01-20 22:28:52

```text
暂无描述
```

**接口状态**

> 开发中

**接口URL**

> http://localhost:8080/todolist/todos/deleteByStatus

| 环境  | URL |
| --- | --- |


**请求方式**

> DELETE

**Content-Type**

> json

**请求Header参数**

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| Authorization | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzc0NjkzNjMsInVzZXJJRCI6NH0.IucBuozYs-Y-5i4tSL36qk6UuszkdwEAvbOqFk3sgXU | String | 是 | - |

**请求Body参数**

```javascript
{
  "status": 1
}

```

**认证方式**

> 继承父级

**响应示例**

* 成功(200)

```javascript
暂无数据
```

* 失败(404)

```javascript
暂无数据
```
