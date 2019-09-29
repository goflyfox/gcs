package system

import (
	"gcs/utils/base"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
)

type SysRoleProject struct {
	// columns START
	Id        int `json:"id" gconv:"id,omitempty"`                // 主键
	RoleId    int `json:"roleId" gconv:"role_id,omitempty"`       // 角色id
	ProjectId int `json:"projectId" gconv:"project_id,omitempty"` // 项目id
	// columns END

	base.BaseModel
}

func (model SysRoleProject) Get() SysRoleProject {
	if model.Id <= 0 {
		glog.Error(model.TableName() + " get id error")
		return SysRoleProject{}
	}

	var resData SysRoleProject
	err := model.dbModel("t").Where(" id = ?", model.Id).Fields(model.columns()).Struct(&resData)
	if err != nil {
		glog.Error(model.TableName()+" get one error", err)
		return SysRoleProject{}
	}

	return resData
}

func (model SysRoleProject) GetOne(form *base.BaseForm) SysRoleProject {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["id"] != "" {
		where += " and id = ? "
		params = append(params, gconv.Int(form.Params["id"]))
	}

	var resData SysRoleProject
	err := model.dbModel("t").Where(where, params).Fields(model.columns()).Struct(&resData)
	if err != nil {
		glog.Error(model.TableName()+" get one error", err)
		return SysRoleProject{}
	}

	return resData
}

func (model SysRoleProject) List(form *base.BaseForm) []SysRoleProject {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["name"] != "" {
		where += " and name like ? "
		params = append(params, "%"+form.Params["name"]+"%")
	}

	var resData []SysRoleProject
	err := model.dbModel("t").Fields(
		model.columns()).Where(where, params).OrderBy(form.OrderBy).Structs(&resData)
	if err != nil {
		glog.Error(model.TableName()+" list error", err)
		return []SysRoleProject{}
	}

	return resData
}

func (model SysRoleProject) Page(form *base.BaseForm) []SysRoleProject {
	if form.Page <= 0 || form.Rows <= 0 {
		glog.Error(model.TableName()+" Page Param error", form.Page, form.Rows)
		return []SysRoleProject{}
	}

	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["name"] != "" {
		where += " and name like ? "
		params = append(params, "%"+form.Params["name"]+"%")
	}

	num, err := model.dbModel("t").Where(where, params).Count()
	form.TotalSize = num
	form.TotalPage = num / form.Rows

	// 没有数据直接返回
	if num == 0 {
		form.TotalPage = 0
		form.TotalSize = 0
		return []SysRoleProject{}
	}

	var resData []SysRoleProject
	pageNum, pageSize := (form.Page-1)*form.Rows, form.Rows
	dbModel := model.dbModel("t").Fields(model.columns() + ",su1.real_name as updateName,su2.real_name as createName")
	dbModel = dbModel.LeftJoin("sys_user su1", " t.update_id = su1.id ")
	dbModel = dbModel.LeftJoin("sys_user su2", " t.update_id = su2.id ")
	err = dbModel.Where(where, params).Limit(pageNum, pageSize).OrderBy(form.OrderBy).Structs(&resData)
	if err != nil {
		glog.Error(model.TableName()+" Page list error", err)
		return []SysRoleProject{}
	}

	return resData
}

func (model SysRoleProject) Delete() int64 {
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

	LogSave(model, DELETE)
	return res
}

func (model SysRoleProject) Update() int64 {
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

	LogSave(model, UPDATE)
	return res
}

func (model *SysRoleProject) Insert() int64 {
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

	LogSave(model, INSERT)
	return res
}

func (model SysRoleProject) dbModel(alias ...string) *gdb.Model {
	var tmpAlias string
	if len(alias) > 0 {
		tmpAlias = " " + alias[0]
	}
	tableModel := g.DB().Table(model.TableName() + tmpAlias).Safe()
	return tableModel
}

func (model SysRoleProject) PkVal() int {
	return model.Id
}

func (model SysRoleProject) TableName() string {
	return "sys_role_project"
}

func (model SysRoleProject) columns() string {
	sqlColumns := "t.id,t.role_id as roleId,t.project_id as projectId"
	return sqlColumns
}
