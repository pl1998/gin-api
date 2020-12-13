## 基于gin 框架待建的一个go api 脚手架

### 初始化操作

```
go clone https://github.com/pl1998/gin-api.git
```

### 进入项目根目录

```
cp .env.example .env

go mod download

air

```
## 功能集成以及使用
  * air 热加载
  * gorm 操作数据库
  * jwt 身份认证
  * redis 缓存
  * gin框架的一些路由、控制器、以及中间件的分离和使用

## 开发日志🎉

| 日期   | 更新级别 | 更新内容      | 贡献者 | 当前状态 |
| ------| -------- | --------- | ---- | ---- |
| 2020-12-10|   feat   | 项目结构待建   | [pl1998](https://github.com/pl1998)  | 已合并到master分支     |
| 2020-12-11|   feat   | 路由、中间件、控制器、验证器编写   | [pl1998](https://github.com/pl1998)  | 已合并到master分支     |
| 2020-12-12|   fix 、feat   | 新增jwt、用户注册接口、目录结构优化   | [pl1998](https://github.com/pl1998)  | 已合并到master分支     |
| 2020-12-13|   feat   | 用户登录、token颁发、校验中间件编写   | [pl1998](https://github.com/pl1998)  | 已合并到master分支     |

## 当前支持api
  * 用户注册 `api/register` 
  * 用户登录 `api/login`    
  * 用户信息 `api/users`    
  
 
