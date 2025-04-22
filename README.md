# 定义接口

### 接口：创建用户

##### 1. 接口描述
- 该接口用于创建新用户账号。

##### 2. 请求方法 & URL
- Method: POST
- URL: /api/v1/login

##### 3. 请求头 (Headers)
- Content-Type: application/json

##### 4. 请求参数
###### 4.1 Body 参数
| 字段名    | 类型   | 必填 | 说明                    | 示例值              |
|-----------|-------|------|-------------------------|--------------------|
| username  | String| 是   | 用户名                  | john_doe           |
| password  | String| 是   | 密码                    | 123456             |
| email     | String| 否   | 邮箱                    | john@example.com   |
| phone     | String| 否   | 手机号                  | 13700000000        |

**请求示例：**
```json
{
  "username": "john_doe",
  "password": "123456",
  "email": "john@example.com",
  "phone": "13700000000"
}
```
##### 5. 返回结果
| 字段名     | 类型     | 说明             |
|---------|--------|----------------|
| code    | Number | 业务状态码；200 表示成功 |
| message | String | 提示信息           |
| data    | Object | 创建成功后返回的用户信息   |
**返回示例**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 101,
    "username": "john_doe",
    "email": "john@example.com",
    "phone": "13700000000",
    "createdAt": "2025-01-01T00:00:00Z",
    "updatedAt": "2025-01-01T00:00:00Z"
  }
}
```
##### 6. 异常返回
| 状态码 | 说明          | 响应示例                                         |
|-----|-------------|----------------------------------------------|
| 400 | 请求参数不合法     | {"code":400,"message":"Invalid params"}      |
| 401 | 未授权/Token无效 | {"code":401,"message":"Unauthorized"}        |
| 409 | 用户已存在       | {"code":409,"message":"User already exists"} |
| 500 | 服务器内部错误     | {"code":500,"message":"Internal Error"}      |
##### 7. 业务说明
- 调用本接口前，需确保用户未被注册，否则返回 code=409。
- 如果需要执行其他相关初始化操作，可以在这里补充说明。

### 接口：用户登陆

##### 1. 接口描述
- 该接口用于创建新用户账号。

##### 2. 请求方法 & URL
- Method: POST
- URL: /api/v1/users

##### 3. 请求头 (Headers)
- Content-Type: application/json
- Authorization: Bearer <token>  (可选，用于鉴权)

##### 4. 请求参数
###### 4.1 Body 参数
| 字段名    | 类型   | 必填 | 说明                    | 示例值              |
|-----------|-------|------|-------------------------|--------------------|
| username  | String| 是   | 用户名                  | john_doe           |
| password  | String| 是   | 密码                    | 123456             |
| email     | String| 否   | 邮箱                    | john@example.com   |
| phone     | String| 否   | 手机号                  | 13700000000        |
| role      | String| 否   | 用户角色 (admin / user) | user               |

**请求示例：**
```json
{
  "username": "john_doe",
  "password": "123456",
  "email": "john@example.com",
  "phone": "13700000000",
  "role": "user"
}
```
##### 5. 返回结果
| 字段名     | 类型     | 说明           |
|---------|--------|--------------|
| code    | Number | 业务状态码；0 表示成功 |
| message | String | 提示信息         |
| data    | Object | 创建成功后返回的用户信息 |
**返回示例**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 101,
    "username": "john_doe",
    "email": "john@example.com",
    "phone": "13700000000",
    "role": "user",
    "createdAt": "2025-01-01T00:00:00Z",
    "updatedAt": "2025-01-01T00:00:00Z"
  }
}
```
##### 6. 异常返回
| 状态码 | 说明          | 响应示例                                         |
|-----|-------------|----------------------------------------------|
| 400 | 请求参数不合法     | {"code":400,"message":"Invalid params"}      |
| 401 | 未授权/Token无效 | {"code":401,"message":"Unauthorized"}        |
| 409 | 用户已存在       | {"code":409,"message":"User already exists"} |
| 500 | 服务器内部错误     | {"code":500,"message":"Internal Error"}      |
##### 7. 业务说明
- 调用本接口前，需确保用户未被注册，否则返回 code=409。
- 如果需要执行其他相关初始化操作，可以在这里补充说明。
