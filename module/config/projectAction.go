package config

import (
	"gcs/module/system"
	"gcs/utils"
	"gcs/utils/base"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/crypto/gmd5"
	"github.com/gogf/gf/g/encoding/gjson"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/os/gtime"
	"github.com/gogf/gf/g/util/gconv"
)

type ProjectAction struct {
	base.BaseRouter
}

var (
	actionNameProject = "ProjectAction"
)

// path: /index
func (action *ProjectAction) Index(r *ghttp.Request) {
	tplFile := "pages/config/project/project_index.html"
	err := r.Response.WriteTpl(tplFile, g.Map{
		"now": gtime.Datetime(),
	})

	if err != nil {
		glog.Error(err)
	}
}

// path: /copy/index
func (action *ProjectAction) Copyindex(r *ghttp.Request) {
	tplFile := "pages/config/projectcopy/project_index.html"
	err := r.Response.WriteTpl(tplFile, g.Map{
		"now": gtime.Datetime(),
	})

	if err != nil {
		glog.Error(err)
	}
}

// path: /get/{id}
func (action *ProjectAction) Get(r *ghttp.Request) {
	id := r.GetInt("id")
	model := TbProject{Id: id}.Get()
	if model.Id <= 0 {
		base.Fail(r, actionNameProject+" get fail")
	}

	base.Succ(r, model)
}

// path: /delete/{id}
func (action *ProjectAction) Delete(r *ghttp.Request) {
	id := r.GetInt("id")

	model := TbProject{Id: id}
	model.UpdateId = base.GetUser(r).Id
	model.UpdateTime = utils.GetNow()

	num := model.Delete()
	if num <= 0 {
		base.Fail(r, actionNameProject+" delete fail")
	}

	base.Succ(r, "")
}

// path: /save
func (action *ProjectAction) Save(r *ghttp.Request) {
	model := TbProject{}
	err := gconv.StructDeep(r.GetPostMap(), &model)
	if err != nil {
		glog.Error(actionNameProject+" save struct error", err)
		base.Error(r, "save error")
	}

	userId := base.GetUser(r).Id

	model.UpdateId = userId
	model.UpdateTime = utils.GetNow()

	var num int64
	if model.Id <= 0 {
		model.SecretKey, _ = gmd5.Encrypt(utils.GetNow())
		model.CreateId = userId
		model.CreateTime = utils.GetNow()
		num = model.Insert()
	} else {
		num = model.Update()
	}

	if num <= 0 {
		base.Fail(r, actionNameProject+" save fail")
	}

	base.Succ(r, "")
}

// path: /list
func (action *ProjectAction) List(r *ghttp.Request) {
	form := base.NewForm(r.GetPostMap())
	model := TbProject{}

	list := model.List(&form)
	base.Succ(r, list)
}

// path: /page
func (action *ProjectAction) Page(r *ghttp.Request) {
	form := base.NewForm(r.GetPostMap())
	model := TbProject{}

	page := model.Page(&form)
	base.Succ(r, g.Map{"list": page, "form": form})
}

// path: /jqgrid
func (action *ProjectAction) Jqgrid(r *ghttp.Request) {
	form := base.NewForm(r.GetPostMap())
	model := TbProject{}

	page := model.Page(&form)
	r.Response.WriteJson(g.Map{
		"page":    form.Page,
		"rows":    page,
		"total":   form.TotalPage,
		"records": form.TotalSize,
	})
}

// path: /data
func (action *ProjectAction) Data(r *ghttp.Request) {
	userId := base.GetUser(r).Id
	model := system.SysUser{Id: userId}.Get()
	if model.Id <= 0 {
		base.Fail(r, "登录异常")
	}

	list := TbProject{}.ListProject(model.Id, model.UserType)
	base.Succ(r, list)
}

// path: /copy
func (action *ProjectAction) Copy(r *ghttp.Request) {
	userId := base.GetUser(r).Id
	model := system.SysUser{Id: userId}.Get()
	if model.Id <= 0 {
		base.Fail(r, "登录异常")
	}

	srcProjectId := model.ProjectId
	destProjectId := r.GetInt("destProjectId")
	TbProject{}.copyProjects(userId, srcProjectId, destProjectId)

	base.Succ(r, "")
}

// path: /srcProject
func (action *ProjectAction) SrcProject(r *ghttp.Request) {
	userId := base.GetUser(r).Id
	model := system.SysUser{Id: userId}.Get()
	if model.Id <= 0 {
		base.Fail(r, "登录异常")
	}

	srcProjectId := model.ProjectId
	project := TbProject{Id: srcProjectId}.Get()
	if project.Id <= 0 {
		base.Fail(r, actionNameProject+" get fail")
	}

	base.Succ(r, project)
}

// path: /srcDestProjectInfo
func (action *ProjectAction) SrcDestProject(r *ghttp.Request) {
	userId := base.GetUser(r).Id
	model := system.SysUser{Id: userId}.Get()
	if model.Id <= 0 {
		base.Fail(r, "登录异常")
	}

	srcProjectId := model.ProjectId
	srcProject := TbProject{Id: srcProjectId}.Get()
	if srcProject.Id <= 0 {
		base.Fail(r, actionNameProject+" get srcProject fail")
	}

	destProjectId := r.GetInt("destProjectId")
	desProject := TbProject{Id: destProjectId}.Get()
	if desProject.Id <= 0 {
		base.Fail(r, actionNameProject+" get desProject fail")
	}
	srcConfigList := system.SysConfig{}.ListByProjectId(srcProjectId, true)
	desConfigList := system.SysConfig{}.ListByProjectId(destProjectId, true)
	for _, data := range srcConfigList {
		data.Id = 0
		data.ParentId = 0
		data.ProjectId = 0
	}
	for _, data := range desConfigList {
		data.Id = 0
		data.ParentId = 0
		data.ProjectId = 0
	}

	srcConfigStr, _ := gjson.Encode(srcConfigList)
	desConfigStr, _ := gjson.Encode(desConfigList)
	base.Succ(r, g.Map{
		"srcProjectConfig":  string(srcConfigStr),
		"destProjectConfig": string(desConfigStr),
	})
}
