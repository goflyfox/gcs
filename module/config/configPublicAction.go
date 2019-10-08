package config

import (
	"gcs/module/system"
	"gcs/utils"
	"gcs/utils/base"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

type ConfigPublicAction struct {
	base.BaseRouter
}

var (
	actionNameConfigPublic = "ConfigPublicAction"
)

// path: /index
func (action *ConfigPublicAction) Index(r *ghttp.Request) {
	tplFile := "pages/config/configpublic/configpublic_index.html"
	err := r.Response.WriteTpl(tplFile, g.Map{
		"now": gtime.Datetime(),
	})

	if err != nil {
		glog.Error(err)
	}
}

// path: /get/{id}
func (action *ConfigPublicAction) Get(r *ghttp.Request) {
	id := r.GetInt("id")
	model := TbConfigPublic{Id: id}.Get()
	if model.Id <= 0 {
		base.Fail(r, actionNameConfigPublic+" get fail")
	}

	before := TbConfigPublic{Id: id, ProjectId: model.ProjectId}.GetBefore()
	if before.Id <= 0 {
		model.BeforeContent = ""
	} else {
		model.BeforeContent = before.Content
	}

	base.Succ(r, model)
}

// path: /project
func (action *ConfigPublicAction) Project(r *ghttp.Request) {
	userId := base.GetUser(r).Id
	user := system.SysUser{Id: userId}.Get()
	if user.Id <= 0 {
		base.Fail(r, "登录异常")
	}

	project := TbProject{Id: user.ProjectId}.Get()
	if project.Id <= 0 {
		base.Fail(r, actionNameProject+" get project fail")
	}

	form := &base.BaseForm{Params: g.MapStrStr{"name": project.Name}}
	configPublic := TbConfigPublic{}.GetOne(form)
	beforeContent := ""
	if configPublic.Id > 0 {
		beforeContent = configPublic.Content
	}

	srcConfigList := system.SysConfig{}.ListByProjectId(user.ProjectId, true)
	// 去除无用字段
	dataList := g.List{}
	for _, data := range srcConfigList {
		dataList = append(dataList, g.Map{
			"name":      data.Name,
			"key":       data.Key,
			"value":     data.Value,
			"code":      data.Code,
			"parentKey": data.ParentKey,
		})
	}
	srcConfigStr, _ := gjson.Encode(dataList)

	base.Succ(r, g.Map{
		"projectId":     project.Id,
		"projectName":   project.Name,
		"beforeContent": beforeContent,
		"content":       gconv.String(srcConfigStr),
	})
}

// path: /delete/{id}
func (action *ConfigPublicAction) Delete(r *ghttp.Request) {
	id := r.GetInt("id")

	model := TbConfigPublic{Id: id}
	model.UpdateId = base.GetUser(r).Id
	model.UpdateTime = utils.GetNow()

	num := model.Delete()
	if num <= 0 {
		base.Fail(r, actionNameConfigPublic+" delete fail")
	}

	base.Succ(r, "")
}

// path: /save
func (action *ConfigPublicAction) Save(r *ghttp.Request) {
	model := TbConfigPublic{}
	err := gconv.Struct(r.GetPostMap(), &model)
	if err != nil {
		glog.Error(actionNameConfigPublic+" save struct error", err)
		base.Error(r, "save error")
	}

	userId := base.GetUser(r).Id

	model.UpdateId = userId
	model.UpdateTime = utils.GetNow()

	var num int64
	if model.Id <= 0 {
		user := system.SysUser{Id: userId}.Get()
		if user.Id <= 0 {
			base.Fail(r, "登录异常")
		}
		model.Version = gtime.Now().Format("YmdHisu")
		srcConfigList := system.SysConfig{}.ListByProjectId(user.ProjectId, true)
		// 去除无用字段
		dataList := g.List{}
		for _, data := range srcConfigList {
			dataList = append(dataList, g.Map{
				"name":      data.Name,
				"key":       data.Key,
				"value":     data.Value,
				"code":      data.Code,
				"parentKey": data.ParentKey,
			})
		}

		srcConfigStr, _ := gjson.Encode(dataList)
		model.Content = gconv.String(srcConfigStr)
		model.CreateId = userId
		model.CreateTime = utils.GetNow()
		num = model.Insert()
	} else {
		num = model.Update()
	}

	if num <= 0 {
		base.Fail(r, actionNameConfigPublic+" save fail")
	}

	base.Succ(r, "")
}

// path: /list
func (action *ConfigPublicAction) List(r *ghttp.Request) {
	form := base.NewForm(r.GetPostMap())
	model := TbConfigPublic{}

	list := model.List(&form)
	base.Succ(r, list)
}

// path: /page
func (action *ConfigPublicAction) Page(r *ghttp.Request) {
	form := base.NewForm(r.GetPostMap())
	model := TbConfigPublic{}

	page := model.Page(&form)
	base.Succ(r, g.Map{"list": page, "form": form})
}

// path: /jqgrid
func (action *ConfigPublicAction) Jqgrid(r *ghttp.Request) {
	form := base.NewForm(r.GetPostMap())
	model := TbConfigPublic{}

	page := model.Page(&form)
	r.Response.WriteJson(g.Map{
		"page":    form.Page,
		"rows":    page,
		"total":   form.TotalPage,
		"records": form.TotalSize,
	})
}
