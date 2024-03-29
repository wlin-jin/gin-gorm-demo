## 项目简介

技术栈： go  gin  gorm

## api接口开发模板

- 发送邮件模块

- 文章列表增删改查功能

- 登录注册功能

## 之后完成的功能

- gorm各种操作方法应用

- gin路由配置，发布方法

- gin项目结构的搭建

目录结构,本项目结构，分成控制器，模型，数据库操作分离，这样项目体积增大后可以不至于太乱，变量容易污染。

# 目录结构


- controller  控制器
    - admin.go   登录控制
    - passages.go  文章管理
    - region.go    机房管理
- database 数据库配置
- models  模型，处理数据库逻辑
- app.go  入口文件
- databaseModel.go  数据库结构体
- router.go  路由文件
- email.go  发送邮件服务

# 项目说明

> controller/region.go

api 透传，将要查找的数据，直接传递key,value  这样可以不需要判断，可以通过查询数据库，轻松的完成查询任务。


# 优化方向：

- 一键打包，部署，用makefile管理
- gorm 自动迁移数据库
- golang 项目可扩展性，通过配置可以连接多个mysql库，多个mysql机器。多个mongodb 机器和表
- 负载均衡，高可用，模块话, 高并发
- 日志管理
- 项目中get post调用
- 企业微信调用
- 数据库操作和api调用分离
- 公用函数分离
- 微服务化，建立api-getway
- 开发中不同情景的使用
    + xlsx 读写
    + email 服务的调用
    + ufile的调用， 阿里云，腾讯云
- 单元测试
- k8s引入


