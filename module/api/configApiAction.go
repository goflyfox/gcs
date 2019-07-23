package api

import (
	"gcs/module/config"
	"gcs/utils/base"
	"gcs/utils/resp"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/crypto/gmd5"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/os/gtime"
	"github.com/gogf/gf/g/text/gstr"
	"github.com/gogf/gf/g/util/gconv"
	"math"
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
	form := base.NewForm(r.GetMap())

	model := config.TbConfigPublic{}.GetOne(&form)
	if model.Id <= 0 {
		base.Fail(r, " get version fail")
	}

	base.Succ(r, g.Map{
		"version": model.Version,
	})
}

// path: /data
func (action *ConfigApiAction) Data(r *ghttp.Request) {
	form := base.NewForm(r.GetMap())

	model := config.TbConfigPublic{}.GetOne(&form)
	if model.Id <= 0 {
		base.Fail(r, " get version fail")
	}

	base.Succ(r, g.Map{
		"version": model.Version,
		"content": model.Content,
	})
}

// path: /auth
func (action *ConfigApiAction) Auth(r *ghttp.Request) {
	form := base.NewForm(r.GetMap())
	bean := ConfigBean{
		Name: form.Params["name"],
		Mac:  form.Params["mac"],
		No:   form.Params["no"],
	}
	resp := Auth(r, bean)

	r.Response.WriteJson(resp)
	r.ExitAll()
}

func Auth(r *ghttp.Request, bean ConfigBean) resp.Resp {
	isApiValidFlag := true
	if !isApiValidFlag {
		return resp.Succ("")
	}

	if bean.Name == "" {
		return resp.Fail("name is null")
	}
	if len(bean.Name) < 1 || len(bean.Name) > 50 {
		return resp.Fail("name is illegal")
	}
	if bean.Mac == "" {
		return resp.Fail("mac is null")
	}
	if len(bean.Mac) != 32 {
		return resp.Fail("mac is illegal")
	}
	if bean.No == "" {
		return resp.Fail("no is null")
	}
	// yyyyMMddHHmmss
	if len(bean.No) < 14 {
		return resp.Fail("no length small")
	}
	if len(bean.No) > 20 {
		return resp.Fail("no length big")
	}
	date := gtime.NewFromStrFormat(bean.No, "YmdHis")
	if date == nil {
		return resp.Fail("no is illegal")
	}
	if math.Abs(gconv.Float64(gtime.Now().Millisecond()-date.Millisecond())) > 10*60*1000 {
		return resp.Fail("time gt 10 minute")
	}

	form := base.NewForm(r.GetMap())

	project := config.TbProject{}.GetOne(&form)
	if project.Id <= 0 {
		return resp.Fail("project is not exist")
	}

	keyStr, err := gmd5.Encrypt(bean.Name + bean.No + project.SecretKey)
	if err != nil {
		glog.Error("mac encrypt error", err)
		return resp.Fail("mac encrypt error")
	}

	if bean.Mac != gstr.ToLower(keyStr) {
		return resp.Fail("mac is illegal")
	}

	return resp.Succ("")
}
