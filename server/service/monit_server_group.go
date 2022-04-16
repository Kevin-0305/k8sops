package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

//@author: [kevin]](https://github.com/kevin-0305)
//@function: CreateServerGroup
//@description: 新增主机分组
//@param: s model.ServerGroup
//@return: err error, serverInter model.ServerGroup
func CreateServerGroup(s model.ServerGroup) (err error, serverGroupInter model.ServerGroup) {
	err = global.GVA_DB.Create(&s).Error
	return err, s
}

//@author: [kevin]](https://github.com/kevin-0305)
//@function: GetServerGroupList
//@description: 获取主机分组列表
//@param: null
//@return: err error, list interface{},
func GetServerGroupList() (err error, list interface{}) {
	db := global.GVA_DB.Model(&model.ServerGroup{})
	var serverGroupList []model.ServerGroup
	err = db.Find(&serverGroupList).Error
	return err, serverGroupList
}

//@author: [kevin]](https://github.com/kevin-0305)
//@function: FindServerGroupById
//@description: 通过id获取列表
//@param: id int
//@return: err error, serverGroup *model.ServerGroup
func FindServerGroupById(id int) (err error, serverGroup model.ServerGroup) {
	err = global.GVA_DB.Where("id = ?", id).First(&serverGroup).Error
	return
}

//@author: [kevin]](https://github.com/kevin-0305)
//@function: SetServerInfo
//@description: 修改服务器分组信息
//@param: ServerGroup model.ServerGroup
//@return: err error, ServerGroup model.ServerGroup
func SetServerGroupInfo(id int, reqServerGroup model.ServerGroup) (err error, serverGroup model.ServerGroup) {
	err = global.GVA_DB.Where("id=?", id).Updates(&reqServerGroup).Error
	return err, reqServerGroup
}

//@author: [kevin]](https://github.com/kevin-0305)
//@function: DeleteServerGroup
//@description: 删除服务器分组
//@param: id float64
//@return: err error
func DeleteServerGroup(id int) (err error) {
	var serverGroup model.ServerGroup
	err = global.GVA_DB.Where("id=?", id).Delete(&serverGroup).Error
	return err
}
