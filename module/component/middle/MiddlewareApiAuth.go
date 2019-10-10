package middle

import (
	"gcs/module/api"
	"gcs/module/config"
	"gcs/module/constants"
	"gcs/utils/base"
	"gcs/utils/cache"
	"gcs/utils/resp"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"math"
)

func MiddlewareApiAuth(r *ghttp.Request) {
	form := base.NewForm(r.GetMap())
	bean := api.ConfigBean{
		Name: form.Params["name"],
		Mac:  form.Params["mac"],
		No:   form.Params["no"],
	}
	resp := Auth(r, bean)

	if !resp.Success() {
		r.Response.WriteJson(resp)
		r.ExitAll()
	} else {
		r.Middleware.Next()
	}

}

func Auth(r *ghttp.Request, bean api.ConfigBean) resp.Resp {
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
	// 1分钟
	var diffTime float64 = 60
	if math.Abs(gconv.Float64(gtime.Now().Second()-date.Second())) > diffTime {
		return resp.Fail("time gt " + gconv.String(diffTime/60) + " minute (" + bean.No + ")")
	}

	project := getPublicCache(r)
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

func getPublicCache(r *ghttp.Request) config.TbProject {
	resp := cache.GetMap(constants.CacheProjectInfoKey)
	var project config.TbProject
	if resp.Success() {
		project = config.TbProject{}
		gconv.Struct(resp.Data, &project)
	} else {
		form := base.NewForm(r.GetMap())
		project = config.TbProject{}.GetOne(&form)
		cache.SetexMap(constants.CacheProjectInfoKey+project.Name, gconv.Map(project), constants.CacheTimeOut)
	}

	return project
}
