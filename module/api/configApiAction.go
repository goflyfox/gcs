package api

import (
	"gcs/module/config"
	"gcs/module/constants"
	"gcs/utils/base"
	"gcs/utils/cache"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
)

type ConfigBean struct {
	// columns START
	No    string `json:"no" gconv:"no,omitempty"`       // 标识
	Name  string `json:"name" gconv:"name,omitempty"`   // 名称
	Key   string `json:"key" gconv:"key,omitempty"`     // 键
	Value string `json:"value" gconv:"value,omitempty"` // 值
	Code  string `json:"code" gconv:"code,omitempty"`   // 编码
	Mac   string `json:"mac" gconv:"mac,omitempty"`     // MAC
	// columns END
}

type ConfigApiAction struct {
	base.BaseRouter
}

// path: /version
func (action *ConfigApiAction) Version(r *ghttp.Request) {
	model := action.getModel(r)
	if model.Id <= 0 {
		base.Fail(r, " get version fail")
	}

	base.Succ(r, g.Map{
		"version": model.Version,
	})
}

// path: /data
func (action *ConfigApiAction) Data(r *ghttp.Request) {
	model := action.getModel(r)
	if model.Id <= 0 {
		base.Fail(r, " get version fail")
	}

	// 转换成对象
	dataList, err := gjson.DecodeToJson(model.Content)
	if err != nil {
		glog.Error("Data error", err)
		base.Error(r, "data error")
	}

	base.Succ(r, g.Map{
		"version": model.Version,
		"content": dataList.ToArray(),
	})
}

func (action *ConfigApiAction) getModel(r *ghttp.Request) config.TbConfigPublic {
	resp := cache.GetMap(constants.CachePublicDataKey)
	var model config.TbConfigPublic
	if resp.Success() {
		model = config.TbConfigPublic{}
		gconv.Struct(resp.Data, &model)
	} else {
		form := base.NewForm(r.GetMap())
		model = config.TbConfigPublic{}.GetOne(&form)
		cache.SetexMap(constants.CachePublicDataKey, gconv.Map(model), constants.CacheTimeOut)
	}

	return model
}
