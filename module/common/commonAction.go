package common

import (
	"gcs/module/constants"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

// Login 登录页面
func Welcome(r *ghttp.Request) {
	err := r.Response.WriteTpl("pages/welcome.html", g.Map{})

	if err != nil {
		glog.Error(err)
	}
}

// Login 登录页面
func Debug(r *ghttp.Request) {
	if constants.DEBUG {
		constants.DEBUG = false
		g.DB().SetDebug(false)
		r.Response.Write("debug close ~!~ ")
	} else {
		constants.DEBUG = true
		g.DB().SetDebug(true)
		r.Response.Write("debug open ~!~ ")
	}
}
