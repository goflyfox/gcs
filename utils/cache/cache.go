package cache

import (
	"gcs/utils/resp"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"time"
)

const (
	ModeCache = 1
	ModeRedis = 2
)

// Set 设置缓存
func SetMap(cacheKey string, data g.Map) resp.Resp {
	return SetexMap(cacheKey, data, 0)
}

// Setex 设置缓存带超时时间
func SetexMap(cacheKey string, data g.Map, timeout int) resp.Resp {
	dataJson, err := gjson.Encode(data)
	if err != nil {
		glog.Error("[Cache]cache json encode error", err)
		return resp.Error("cache json encode error")
	}

	return Setex(cacheKey, gconv.String(dataJson), timeout)
}

// GetMap 获取缓存
func GetMap(cacheKey string) resp.Resp {
	data := Get(cacheKey)
	if !data.Success() {
		return data
	}

	var dataMap g.Map
	err := gjson.DecodeTo(data.Data, &dataMap)
	if err != nil {
		glog.Error("[Cache]cache get json error", err)
		return resp.Error("cache get json error")
	}

	return resp.Succ(dataMap)
}

// Set 设置缓存
func Set(cacheKey string, data string) resp.Resp {
	return Setex(cacheKey, data, 0)
}

// Setex 设置缓存带超时时间
func Setex(cacheKey string, data string, timeout int) resp.Resp {
	cacheMode := getCacheMode()
	switch cacheMode {
	case ModeCache:
		gcache.Set(cacheKey, data, gconv.Duration(timeout)*time.Millisecond)
	case ModeRedis:
		var err error
		// 设置redis缓存
		if timeout <= 0 {
			_, err = g.Redis().Do("SET", cacheKey, data)
		} else {
			_, err = g.Redis().Do("SETEX", cacheKey, timeout, data)
		}

		if err != nil {
			glog.Error("[Cache]cache set error", err)
			return resp.Error("cache set error")
		}
	default:
		return resp.Error("cache model error")
	}

	return resp.Succ(data)
}

// Get 获取缓存
func Get(cacheKey string) resp.Resp {
	cacheMode := getCacheMode()
	var dataStr string

	switch cacheMode {
	case ModeCache:
		data := gcache.Get(cacheKey)
		if data == nil {
			return resp.Fail("data is nil")
		}
		dataStr = gconv.String(data)
	case ModeRedis:
		data, err := g.Redis().Do("GET", cacheKey)
		if err != nil {
			glog.Error("[Cache]cache get error", err)
			return resp.Error("cache get error")
		}
		if data == nil {
			return resp.Fail("data is nil")
		}
		dataStr = gconv.String(data)
	default:
		return resp.Error("cache model error")
	}

	return resp.Succ(dataStr)
}

// Del 删除缓存
func Del(cacheKey string) resp.Resp {
	cacheMode := getCacheMode()
	switch cacheMode {
	case ModeCache:
		gcache.Remove(cacheKey)
	case ModeRedis:
		var err error
		_, err = g.Redis().Do("DEL", cacheKey)
		if err != nil {
			glog.Error("[Cache]cache remove error", err)
			return resp.Error("cache remove error")
		}
	default:
		return resp.Error("cache model error")
	}

	return resp.Succ("")
}

func getCacheMode() int8 {
	cacheMode := g.Config().GetInt8("cache-mode", 1)
	if cacheMode == 0 || cacheMode > ModeRedis {
		cacheMode = ModeCache
	}
	return cacheMode
}
