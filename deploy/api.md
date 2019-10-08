## 接口文档
------------------------

* 配置管理平台主要对外提供最新发布版本获取和最新发布内容获取接口。
* 接口支持安全认证，有效期认证等。
* 接口参数秘钥SecretKey获取：项目管理，查看具体项目秘钥。

## 接口使用说明
------------------------

* 公共请求参数

```go
No    string `json:"no" gconv:"no,omitempty"`       // 标识
Name  string `json:"name" gconv:"name,omitempty"`   // 名称
Mac   string `json:"mac" gconv:"mac,omitempty"`     // MAC
```

加密统一采用MD5(Name + No + SecretKey)转小写

* 文档说明

```
提供版本获取和内容获取接口
```

## 接口
------------------------

####  版本获取接口

* 接口说明：发布最新版本获取接口
* 请求方式： **_GET/POST_**
* 请求地址：**_/config/api/version_**
* 请求参数

```
公共参数
```

* 示例

```
http://127.0.0.1:80/config/api/version?name=test&no=20191008132205&mac=2f7c84e4185cb14276e6fb758e668ffb
```

* 返回结果

```json
{
	"code": 0,
	"msg": "success",
	"data": {
		"version": "20191008111411398"
	}
}
```

####  配置内容获取接口

* 接口说明：发布最新内容获取接口
* 请求方式： **_GET/POST_**
* 请求地址：**_/config/api/data_**
* 请求参数

```
公共参数
```

* 示例

```
http://127.0.0.1:80/config/api/data?name=test&no=20191008132206&mac=952f3f234ef21551794ea5731ff537fc
```

* 返回结果

```json
{
	"code": 0,
	"msg": "success",
	"data": {
		"content": [{
			"code": "",
			"key": "sms.username",
			"name": "短信账号",
			"parentKey": "sms",
			"value": "test"
		}, {
			"code": "",
			"key": "sms.type",
			"name": "短信类型",
			"parentKey": "sms",
			"value": "阿里云"
		}, {
			"code": "",
			"key": "sms.passwd",
			"name": "短信密码",
			"parentKey": "sms",
			"value": "111111"
		}, {
			"code": "",
			"key": "system.debug",
			"name": "日志控制配置",
			"parentKey": "system",
			"value": "false"
		}, {
			"code": "",
			"key": "system",
			"name": "系统参数",
			"parentKey": "",
			"value": ""
		}, {
			"code": "",
			"key": "sms",
			"name": "短信配置",
			"parentKey": "",
			"value": ""
		}, {
			"code": "",
			"key": "sex",
			"name": "性别",
			"parentKey": "",
			"value": ""
		}, {
			"code": "1",
			"key": "sex.male",
			"name": "性别男",
			"parentKey": "sex",
			"value": "男"
		}, {
			"code": "2",
			"key": "sex.female",
			"name": "性别女",
			"parentKey": "sex",
			"value": "女"
		}, {
			"code": "3",
			"key": "sex.unknown",
			"name": "性别未知",
			"parentKey": "sex",
			"value": "未知"
		}],
		"version": "20191008111411398"
	}
}
```
