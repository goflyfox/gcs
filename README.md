# gcs

#### 介绍
gcs(go config server) 配置管理平台,此项目基于gf框架开发，支持项目管理、配置管理、配置发布、项目对比、配置复制，接口获取配置等；

* github地址：https://github.com/goflyfox/gcs
* gitee地址：https://gitee.com/goflyfox/gcs

#### 安装教程

1. 从git下载项目： git clone https://github.com/goflyfox/gcs
2. 安装mysql数据库，创建db，运行deploy下gcs.sql脚本
3. 修改config下config.toml配置文件
```toml
# 数据库配置
[database]
    link = "root:123456@tcp(127.0.0.1:3306)/gcs"
```
4. `go build && go run main.go`
5. 访问http://localhost即可看到登录页面，账号/密码：admin/123456

#### 功能模块

1. 登录、认证、登出
2. 项目管理
3. 参数配置
4. 配置发布
5. 其他：组织机构管理、用户管理、角色管理、菜单管理、日志管理

#### 平台截图

登录：
![image](https://raw.githubusercontent.com/goflyfox/gmanager/master/deploy/image/1.png)

组织机构：
![image](https://raw.githubusercontent.com/goflyfox/gmanager/master/deploy/image/2.png)

用户管理：
![image](https://raw.githubusercontent.com/goflyfox/gmanager/master/deploy/image/3.png)

日志管理：
![image](https://raw.githubusercontent.com/goflyfox/gmanager/master/deploy/image/4.png)

#### 感谢

1. gf框架 [https://github.com/gogf/gf](https://github.com/gogf/gf) 
