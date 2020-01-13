package config

import (
	"gcs/module/constants"
	"gcs/module/system"
	"gcs/utils"
	"gcs/utils/base"
	"gcs/utils/cache"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
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

	// 删除缓存
	cache.Del(constants.CacheProjectInfoKey + model.Name)

	base.Succ(r, "")
}

// path: /save
func (action *ProjectAction) Save(r *ghttp.Request) {
	model := TbProject{}
	err := gconv.StructDeep(r.GetMap(), &model)
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
	// 加入缓存
	cache.SetexMap(constants.CacheProjectInfoKey+model.Name, gconv.Map(model), constants.CacheTimeOut)

	base.Succ(r, "")
}

// path: /list
func (action *ProjectAction) List(r *ghttp.Request) {
	form := base.NewForm(r.GetMap())
	model := TbProject{}

	list := model.List(&form)
	base.Succ(r, list)
}

// path: /page
func (action *ProjectAction) Page(r *ghttp.Request) {
	form := base.NewForm(r.GetMap())
	model := TbProject{}

	page := model.Page(&form)
	base.Succ(r, g.Map{"list": page, "form": form})
}

// path: /jqgrid
func (action *ProjectAction) Jqgrid(r *ghttp.Request) {
	form := base.NewForm(r.GetMap())
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

// 修改当前用户密码
func (action *ProjectAction) User(r *ghttp.Request) {
	userId := base.GetUser(r).Id
	model := system.SysUser{Id: userId}.Get()
	if model.Id <= 0 {
		base.Fail(r, "登录异常")
	}

	projectId := r.GetInt("projectId")
	if projectId == 0 {
		base.Fail(r, "参数错误")
	}

	model.UpdateId = userId
	model.UpdateTime = utils.GetNow()
	model.ProjectId = projectId
	// TODO
	project := TbProject{Id: projectId}.Get()
	model.ProjectName = project.Name

	num := model.Update()

	if num <= 0 {
		base.Fail(r, actionNameProject+" Project fail")
	}

	base.Succ(r, "")
}
