# GoSword-Admin
GO语言通用后端框架

# 项目简介
**GoSword**使用Gin、Gorm、Casbin、JWToken、Redis、VUE技术栈，开发的前后端分离后台管理系统。  
**默认管理员账号密码： admin / 123456**

# 项目源码  
|  |  后端代码   | 前端代码 |
|  ----  |  ----  | ----  |
| GitHub | https://github.com/sanyueruanjian/go-sword-admin  | https://github.com/sanyueruanjian/go-sword-admin-web |

# 系统功能
- 用户管理：提供用户的相关配置，新增用户后，默认密码为123456  
- 角色管理：对权限与菜单进行分配，可根据部门设置角色的数据权限  
- 菜单管理：已实现菜单动态路由，后端可配置化，支持多级菜单  
- 部门管理：可配置系统组织架构，树形表格展示  
- 岗位管理：配置各个部门的职位  

# 系统监控
- 在线用户：记录登陆系统的用户 TODO
- 操作日志：记录用户的操作情况 TODO
- 异常日志：记录用户的异常操作情况 TODO

# 项目结构 
```
- app            应用模块
    admin        后台服务 
- common         公共API  
    database     数据库服务
    global       全局服务配置
    logger       日志服务
    middleware   中间件
    router       全局路由注册
    run          运行初始化
- docs           swagger生成文件
- logs           日志目录
- pkg            工具模块
- settings       配置文件目录
- static         静态文件目录
- utils          工具包
- go_sword.sql   数据库文件
- main.go        服务启动文件
- Dockerfile     TODO
```

# 使用指南
 [参考文档](https://docs.qq.com/doc/DRVdKU2FZVnFia0pu)
 
 # 参与开发
 - [@三月软件提供技术支持](http://www.marchsoft.cn/)
 - 由[@Ymq](https://github.com/KAILINYmq)、[@ChenGuangLan](https://github.com/sFFbLL)、[@LiJiaKun](https://github.com/lijiakun123)、[@LinBoLun](https://github.com/linbolun-525)、[@WuYanSong](https://github.com/Enlightemmm)、[@LiuZhiChao](https://github.com/wodeshijie1)参与初版的开发维护
