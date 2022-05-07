## 项目介绍


## 部署

## 开发指南

### 数据库

各位小伙伴可以使用我的云数据库

```plain
# 服务器IP
mysql 8.0版本

host: 47.96.30.253:3306
# 数据库账号密码
username: test01  password: Douyintest01.
username: test02  password: Douyintest02.
```

可以使用 Goland 自带的 `Database` 工具 直接连接，可视化操作也很方便，大可不必再用其他的工具

### 依赖说明

> 简单介绍此项目所需的一些第三方库，以及选型原因

#### gorm

[官网地址](https://gorm.io/zh_CN/docs/index.html)
目前主流 GO 开源ORM 框架

#### gin

[官网地址](https://gin-gonic.com/zh-cn/docs/introduction/)
主流 Web 服务框架，简单易用，功能强大

#### validator

[项目地址](https://github.com/go-playground/validator)
具体用法可以参考官方给的demo

#### gini

[官网地址](https://ini.unknwon.io/docs)
解析项目运行的配置参数,简单易用

