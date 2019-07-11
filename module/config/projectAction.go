package config

import (
	"gcs/utils"
	"gcs/utils/base"
	"github.com/gogf/gf/g"
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
	tplFile := "pages/system/project_index.html"
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
	err := gconv.Struct(r.GetPostMap(), &model)
	if err != nil {
		glog.Error(actionNameProject+" save struct error", err)
		base.Error(r, "save error")
	}

	userId := base.GetUser(r).Id

	model.UpdateId = userId
	model.UpdateTime = utils.GetNow()

	var num int64
	if model.Id <= 0 {
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
