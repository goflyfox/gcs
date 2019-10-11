package config

import (
	"gcs/module/constants"
	"gcs/module/system"
	"gcs/utils/base"
	"gcs/utils/cache"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
)

type TbConfigPublic struct {
	// columns START
	Id          int    `json:"id" gconv:"id,omitempty"`                    // 主键
	ProjectId   int    `json:"projectId" gconv:"project_id,omitempty"`     // 项目ID
	ProjectName string `json:"projectName" gconv:"project_name,omitempty"` // 项目名称
	Version     string `json:"version" gconv:"version,omitempty"`          // 版本
	Content     string `json:"content" gconv:"content,omitempty"`          // 内容
	// columns END
	BeforeContent string `json:"beforeContent" gconv:"before_content,omitempty"` // 之前内容

	base.BaseModel
}

func (model TbConfigPublic) GetBefore() TbConfigPublic {
	if model.Id <= 0 {
		glog.Error(model.TableName() + " get id error")
		return TbConfigPublic{}
	}

	var resData TbConfigPublic
	err := model.dbModel("t").Where(
		" project_id = ?", model.ProjectId).Where(
		" id < ?", model.Id).Fields(model.columns()).OrderBy(
		"id desc").Limit(1).Struct(&resData)
	if err != nil {
		glog.Error(model.TableName()+" get one error", err)
		return TbConfigPublic{}
	}

	return resData
}

func (model TbConfigPublic) Get() TbConfigPublic {
	if model.Id <= 0 {
		glog.Error(model.TableName() + " get id error")
		return TbConfigPublic{}
	}

	var resData TbConfigPublic
	err := model.dbModel("t").Where(" id = ?", model.Id).Fields(model.columns()).Struct(&resData)
	if err != nil {
		glog.Error(model.TableName()+" get one error", err)
		return TbConfigPublic{}
	}

	return resData
}

func (model TbConfigPublic) GetOne(form *base.BaseForm) TbConfigPublic {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["id"] != "" {
		where += " and id = ? "
		params = append(params, gconv.Int(form.Params["id"]))
	}

	if form.Params != nil && form.Params["name"] != "" {
		where += " and project_name = ? "
		params = append(params, form.Params["name"])

		where += " and enable = ? "
		params = append(params, constants.EnableYes)

		form.OrderBy = "id desc"
	}

	var resData TbConfigPublic
	err := model.dbModel("t").Where(where, params).Fields(
		model.columns()).OrderBy(form.OrderBy).Limit(1).Struct(&resData)
	if err != nil {
		glog.Error(model.TableName()+" get one error", err)
		return TbConfigPublic{}
	}

	return resData
}

func (model TbConfigPublic) List(form *base.BaseForm) []TbConfigPublic {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["name"] != "" {
		where += " and name like ? "
		params = append(params, "%"+form.Params["name"]+"%")
	}

	var resData []TbConfigPublic
	err := model.dbModel("t").Fields(
		model.columns()).Where(where, params).OrderBy(form.OrderBy).Structs(&resData)
	if err != nil {
		glog.Error(model.TableName()+" list error", err)
		return []TbConfigPublic{}
	}

	return resData
}

func (model TbConfigPublic) Page(form *base.BaseForm) []TbConfigPublic {
	if form.Page <= 0 || form.Rows <= 0 {
		glog.Error(model.TableName()+" Page Param error", form.Page, form.Rows)
		return []TbConfigPublic{}
	}

	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["name"] != "" {
		where += " and name like ? "
		params = append(params, "%"+form.Params["name"]+"%")
	}

	num, err := model.dbModel("t").Where(where, params...).Count()
	form.TotalSize = num
	form.TotalPage = num / form.Rows

	// 没有数据直接返回
	if num == 0 {
		form.TotalPage = 0
		form.TotalSize = 0
		return []TbConfigPublic{}
	}

	var resData []TbConfigPublic
	pageNum, pageSize := (form.Page-1)*form.Rows, form.Rows
	dbModel := model.dbModel("t").Fields(model.columns() + ",su1.real_name as updateName,su2.real_name as createName")
	dbModel = dbModel.LeftJoin("sys_user su1", " t.update_id = su1.id ")
	dbModel = dbModel.LeftJoin("sys_user su2", " t.update_id = su2.id ")
	err = dbModel.Where(where, params...).Limit(pageNum, pageSize).OrderBy(form.OrderBy).Structs(&resData)
	if err != nil {
		glog.Error(model.TableName()+" Page list error", err)
		return []TbConfigPublic{}
	}

	return resData
}

func (model TbConfigPublic) Delete() int64 {
	if model.Id <= 0 {
		glog.Error(model.TableName() + " delete id error")
		return 0
	}

	r, err := model.dbModel().Where(" id = ?", model.Id).Delete()
	if err != nil {
		glog.Error(model.TableName()+" delete error", err)
		return 0
	}

	res, err2 := r.RowsAffected()
	if err2 != nil {
		glog.Error(model.TableName()+" delete res error", err2)
		return 0
	}

	system.LogSave(model, system.DELETE)
	return res
}

func (model TbConfigPublic) Update() int64 {
	r, err := model.dbModel().Data(model).Where(" id = ?", model.Id).Update()
	if err != nil {
		glog.Error(model.TableName()+" update error", err)
		return 0
	}

	res, err2 := r.RowsAffected()
	if err2 != nil {
		glog.Error(model.TableName()+" update res error", err2)
		return 0
	}

	system.LogSave(model, system.UPDATE)
	return res
}

func (model *TbConfigPublic) Insert() int64 {
	r, err := model.dbModel().Data(model).Insert()
	if err != nil {
		glog.Error(model.TableName()+" insert error", err)
		return 0
	}

	res, err2 := r.RowsAffected()
	if err2 != nil {
		glog.Error(model.TableName()+" insert res error", err2)
		return 0
	} else if res > 0 {
		lastId, err2 := r.LastInsertId()
		if err2 != nil {
			glog.Error(model.TableName()+" LastInsertId res error", err2)
			return 0
		} else {
			model.Id = gconv.Int(lastId)
		}
	}

	system.LogSave(model, system.INSERT)
	return res
}

func (model *TbConfigPublic) GetCacheModel(form *base.BaseForm) TbConfigPublic {
	resp := cache.GetMap(constants.CachePublicDataKey)
	var publicModel TbConfigPublic
	if resp.Success() {
		publicModel = TbConfigPublic{}
		gconv.Struct(resp.Data, &publicModel)
	} else {
		publicModel = TbConfigPublic{}.GetOne(form)
		cache.SetexMap(constants.CachePublicDataKey, gconv.Map(publicModel), constants.CacheTimeOut)
	}

	return publicModel
}

func (model TbConfigPublic) dbModel(alias ...string) *gdb.Model {
	var tmpAlias string
	if len(alias) > 0 {
		tmpAlias = " " + alias[0]
	}
	tableModel := g.DB().Table(model.TableName() + tmpAlias).Safe()
	return tableModel
}

func (model TbConfigPublic) PkVal() int {
	return model.Id
}

func (model TbConfigPublic) TableName() string {
	return "tb_config_public"
}

func (model TbConfigPublic) columns() string {
	sqlColumns := "t.id,t.project_id as projectId,t.project_name as projectName,t.version,t.content,t.enable,t.update_time as updateTime,t.update_id as updateId,t.create_time as createTime,t.create_id as createId"
	return sqlColumns
}
