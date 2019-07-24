package config

import (
	"gcs/module/constants"
	"gcs/module/system"
	"gcs/utils/base"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/database/gdb"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/util/gconv"
)

type TbProject struct {
	// columns START
	Id        int    `json:"id" gconv:"id,omitempty"`                // 主键
	Name      string `json:"name" gconv:"name,omitempty"`            // 项目名称
	SecretKey string `json:"secretKey" gconv:"secret_key,omitempty"` // 秘钥
	Type      string `json:"type" gconv:"type,omitempty"`            // 类型
	Sort      int    `json:"sort" gconv:"sort,omitempty"`            // 排序
	Remark    string `json:"remark" gconv:"remark,omitempty"`        // 备注
	// columns END

	base.BaseModel
}

func (model TbProject) Get() TbProject {
	if model.Id <= 0 {
		glog.Error(model.TableName() + " get id error")
		return TbProject{}
	}

	var resData TbProject
	err := model.dbModel("t").Where(" id = ?", model.Id).Fields(model.columns()).Struct(&resData)
	if err != nil {
		glog.Error(model.TableName()+" get one error", err)
		return TbProject{}
	}

	return resData
}

func (model TbProject) GetOne(form *base.BaseForm) TbProject {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["id"] != "" {
		where += " and id = ? "
		params = append(params, gconv.Int(form.Params["id"]))
	}
	if form.Params != nil && form.Params["name"] != "" {
		where += " and name = ? "
		params = append(params, form.Params["name"])

		where += " and enable = ? "
		params = append(params, constants.EnableYes)

		form.OrderBy = "id desc"
	}

	var resData TbProject
	err := model.dbModel("t").Where(where, params).Fields(
		model.columns()).Limit(1).Struct(&resData)
	if err != nil {
		glog.Error(model.TableName()+" get one error", err)
		return TbProject{}
	}

	return resData
}

func (model TbProject) ListProject(userId int, userType int) []TbProject {
	if userType == constants.UserTypeAdmin {
		return model.List(&base.BaseForm{})
	}

	var resData []TbProject
	err := model.dbModel("t").Fields(model.columns()).LeftJoin(
		"sys_role_project rp", "rp.project_id = t.id ").LeftJoin(
		"sys_user_role ur", "ur.role_id = rp.role_id ").Where(
		"ur.user_id = ? ", userId).Structs(&resData)
	if err != nil {
		glog.Error(model.TableName()+" list error", err)
		return []TbProject{}
	}

	return resData
}

func (model TbProject) List(form *base.BaseForm) []TbProject {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["name"] != "" {
		where += " and name like ? "
		params = append(params, "%"+form.Params["name"]+"%")
	}

	var resData []TbProject
	err := model.dbModel("t").Fields(
		model.columns()).Where(where, params...).OrderBy(form.OrderBy).Structs(&resData)
	if err != nil {
		glog.Error(model.TableName()+" list error", err)
		return []TbProject{}
	}

	return resData
}

func (model TbProject) Page(form *base.BaseForm) []TbProject {
	if form.Page <= 0 || form.Rows <= 0 {
		glog.Error(model.TableName()+" Page Param error", form.Page, form.Rows)
		return []TbProject{}
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
		return []TbProject{}
	}

	var resData []TbProject
	pageNum, pageSize := (form.Page-1)*form.Rows, form.Rows
	dbModel := model.dbModel("t").Fields(model.columns() + ",su1.real_name as updateName,su2.real_name as createName")
	dbModel = dbModel.LeftJoin("sys_user su1", " t.update_id = su1.id ")
	dbModel = dbModel.LeftJoin("sys_user su2", " t.update_id = su2.id ")
	err = dbModel.Where(where, params...).Limit(pageNum, pageSize).OrderBy(form.OrderBy).Structs(&resData)
	if err != nil {
		glog.Error(model.TableName()+" Page list error", err)
		return []TbProject{}
	}

	return resData
}

func (model TbProject) Delete() int64 {
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

func (model TbProject) Update() int64 {
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

func (model *TbProject) Insert() int64 {
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

// TODO 需要编写复制方法
func (model TbProject) copyProjects(userId int, srcProjectId int, destProjectId int) {

	//srcProject := TbProject{Id: srcProjectId}.Get()
	//destProject := TbProject{Id: destProjectId}.Get()
	//
	//srcConfigList := system.SysConfig{}.ListByProjectId(srcProjectId, true)
	//desConfigList := system.SysConfig{}.ListByProjectId(destProjectId, true)
	//
	//if len(srcConfigList) == 0 {
	//	glog.Error(model.TableName() + " srcConfigList no data")
	//	return
	//}
	//
	//srcSet := gset.New()
	//for _, config := range srcConfigList {
	//	srcSet.Add(config.Key)
	//}

}

func (model TbProject) dbModel(alias ...string) *gdb.Model {
	var tmpAlias string
	if len(alias) > 0 {
		tmpAlias = " " + alias[0]
	}
	tableModel := g.DB().Table(model.TableName() + tmpAlias).Safe()
	return tableModel
}

func (model TbProject) PkVal() int {
	return model.Id
}

func (model TbProject) TableName() string {
	return "tb_project"
}

func (model TbProject) columns() string {
	sqlColumns := "t.id,t.name,t.secret_key as secretKey,t.type,t.sort,t.remark,t.enable,t.update_time as updateTime,t.update_id as updateId,t.create_time as createTime,t.create_id as createId"
	return sqlColumns
}
