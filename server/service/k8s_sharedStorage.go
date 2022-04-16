package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateSharedStorage
//@description: 创建SharedStorage记录
//@param: ShrSto model.SharedStorage
//@return: err error

func CreateSharedStorage(ShrSto model.SharedStorage) (err error) {
	err = global.GVA_DB.Create(&ShrSto).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSharedStorage
//@description: 删除SharedStorage记录
//@param: ShrSto model.SharedStorage
//@return: err error

func DeleteSharedStorage(ShrSto model.SharedStorage) (err error) {
	err = global.GVA_DB.Delete(&ShrSto).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSharedStorageByIds
//@description: 批量删除SharedStorage记录
//@param: ids request.IdsReq
//@return: err error

func DeleteSharedStorageByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SharedStorage{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateSharedStorage
//@description: 更新SharedStorage记录
//@param: ShrSto *model.SharedStorage
//@return: err error

func UpdateSharedStorage(ShrSto model.SharedStorage) (err error) {
	err = global.GVA_DB.Save(&ShrSto).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSharedStorage
//@description: 根据id获取SharedStorage记录
//@param: id uint
//@return: err error, ShrSto model.SharedStorage

func GetSharedStorage(id uint) (err error, ShrSto model.SharedStorage) {
	err = global.GVA_DB.Where("id = ?", id).First(&ShrSto).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSharedStorageInfoList
//@description: 分页获取SharedStorage记录
//@param: info request.SharedStorageSearch
//@return: err error, list interface{}, total int64

func GetSharedStorageInfoList(info request.SharedStorageSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.SharedStorage{})
    var ShrStos []model.SharedStorage
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.ServerIP != "" {
        db = db.Where("`server_ip` = ?",info.ServerIP)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&ShrStos).Error
	return err, ShrStos, total
}