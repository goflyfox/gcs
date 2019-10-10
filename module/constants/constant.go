package constants

// 调试日志配置
var DEBUG = true

const (
	CacheUserTokenPre   = "API353_USER_TOKEN_"
	EnableYes           = 1
	UserTypeAdmin       = 1
	CachePublicDataKey  = "API:PUBLIC_DATA"   // 接口发布数据Key
	CacheProjectInfoKey = "API:PROJECT_"      // 项目信息Key
	CacheTimeOut        = 24 * 60 * 60 * 1000 // 缓存默认超时24个小时
)
