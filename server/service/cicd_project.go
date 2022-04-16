package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateProject
//@description: 创建Project记录
//@param: pro model.Project
//@return: err error

func CreateProject(pro model.Project) (err error) {
	err = global.GVA_DB.Create(&pro).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteProject
//@description: 删除Project记录
//@param: pro model.Project
//@return: err error

func DeleteProject(pro model.Project) (err error) {
	err = global.GVA_DB.Delete(&pro).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteProjectByIds
//@description: 批量删除Project记录
//@param: ids request.IdsReq
//@return: err error

func DeleteProjectByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Project{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateProject
//@description: 更新Project记录
//@param: pro *model.Project
//@return: err error

func UpdateProject(pro model.Project) (err error) {
	err = global.GVA_DB.Save(&pro).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetProject
//@description: 根据id获取Project记录
//@param: id uint
//@return: err error, pro model.Project

func GetProject(id uint) (err error, pro model.Project) {
	err = global.GVA_DB.Where("id = ?", id).First(&pro).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetProjectInfoList
//@description: 分页获取Project记录
//@param: info request.ProjectSearch
//@return: err error, list interface{}, total int64

func GetProjectInfoList(info request.ProjectSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Project{})
    var pros []model.Project
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
    }
    if info.Language != 0 {
        db = db.Where("`language` = ?",info.Language)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&pros).Error
	return err, pros, total
}