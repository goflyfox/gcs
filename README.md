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
4. `go run main.go` ; 编译使用 `go build`
5. 访问http://localhost即可看到登录页面，账号/密码：admin/123456

#### 其他配置

1. 如果使用集群，可以开启redis token模式；配置如下：
```toml
# 缓存模式 1 gcache 2 gredis
cache-mode = 2

# Redis数据库配置
[redis]
  default = "127.0.0.1:16379,0,soccer"
  cache   = "127.0.0.1:16379,1,soccer?idleTimeout=600"
```

#### 功能模块

1. 登录、认证、登出
2. 项目管理
3. 参数配置
4. 配置发布
5. 支持菜单权限及项目权限配置
6. 其他：组织机构管理、用户管理、角色管理、菜单管理、日志管理

#### 对外接口
1. [配置平台接口](https://github.com/goflyfox/gcs/blob/master/deploy/api.md "配置平台接口")
2. [基于GF客户端实现](https://github.com/goflyfox/gcsc)
```
github地址：https://github.com/goflyfox/gcsc
gitee地址：https://gitee.com/goflyfox/gcsc
```

#### 平台截图

登录：
![image](https://raw.githubusercontent.com/goflyfox/gcs/master/deploy/image/1.png)

配置管理：
![image](https://raw.githubusercontent.com/goflyfox/gcs/master/deploy/image/g1.png)

配置发布：
![image](https://raw.githubusercontent.com/goflyfox/gcs/master/deploy/image/g2.png)

配置对比：
![image](https://raw.githubusercontent.com/goflyfox/gcs/master/deploy/image/g3.png)

组织机构：
![image](https://raw.githubusercontent.com/goflyfox/gcs/master/deploy/image/2.png)

用户管理：
![image](https://raw.githubusercontent.com/goflyfox/gcs/master/deploy/image/3.png)

日志管理：
![image](https://raw.githubusercontent.com/goflyfox/gcs/master/deploy/image/4.png)

#### 感谢

1. gf框架 [https://github.com/gogf/gf](https://github.com/gogf/gf) 

#### 项目支持

- 项目的发展，离不开大家得支持~！~

- [阿里云最新活动：双11最新活动，低至1折；还有新人礼包；请点击这里](https://www.aliyun.com/1111/2019/home?spm=5176.11533457.1089570.70.4fe277e3TKVLoB&userCode=c4hsn0gc)
- 1核2G1M40G盘，86元/1年， 
- 2核4G3M40G盘，799元/3年，
- 2核8G5M40G盘，1399元/3年。

- [阿里云：ECS云服务器2折起；请点击这里](https://www.aliyun.com/acts/limit-buy?spm=5176.11544616.khv0c5cu5.1.1d8e23e8XHvEIq&userCode=c4hsn0gc)
- [阿里云：ECS云服务器新人优惠券；请点击这里](https://promotion.aliyun.com/ntms/yunparter/invite.html?userCode=c4hsn0gc)

- 也可以请作者喝一杯咖啡:)

![jflyfox](https://raw.githubusercontent.com/jflyfox/jfinal_cms/master/doc/pay01.jpg "Open source support")

