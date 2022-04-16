package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [kevin]](https://github.com/kevin-0305)
//@function: CreateServer
//@description: 新增主机
//@param: u model.ServerHost
//@return: err error, serverInter model.ServerHost

func CreateServer(s model.MonitorServer) (err error, serverInter model.MonitorServer) {
	// var server model.MonitorServer
	// if !errors.Is(global.GVA_DB.Where("pubIpAddress = ? AND ServerType = ?", s.PubIpAddress, s.ServerType).First(&server).Error, gorm.ErrRecordNotFound) {
	// 	return errors.New("此服务器以存在"), serverInter
	// }
	//s.Password = string(utils.RSA_Encrypt([]byte(s.Password), global.GVA_CONFIG.Rsa.RsaPublicKeyFile))
	err = global.GVA_DB.Create(&s).Error
	return err, s
}

//@author: [kevin]](https://github.com/kevin-0305)
//@function: GetServerList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64
func GetServerList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.MonitorServer{})
	var serverList []model.MonitorServer
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&serverList).Error
	return err, serverList, total
}

//@author: [kevin]](https://github.com/kevin-0305)
//@function: GetAllServerList
//@description: 获取所有服务器信息
//@param:
//@return: err error, list interface{}, total int64
func GetAllServerList() (err error, serverList []model.MonitorServer) {
	db := global.GVA_DB.Model(&model.MonitorServer{})
	err = db.Find(&serverList).Error
	return err, serverList
}

//@author: [kevin]](https://github.com/kevin-0305)
//@function: FindServerById
//@description: 通过id获取服务器信息
//@param: id int
//@return: err error, server *model.MonitorServer
func FindServerById(id int) (err error, server model.MonitorServer) {

	err = global.GVA_DB.Where("id = ?", id).First(&server).Error
	// if err = global.GVA_DB.Where("id = ?", id).First(&s).Error; err != nil {
	// 	return errors.New("查询服务器不存在"), &s
	// }
	return
}

//@author: [kevin]](https://github.com/kevin-0305)
//@function: SetServerInfo
//@description: 设置服务器信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func SetServerInfo(id int, reqServer model.MonitorServer) (err error, server model.MonitorServer) {
	err = global.GVA_DB.Where("id=?", id).Updates(&reqServer).Error
	return err, reqServer
}

//@author: [kevin]](https://github.com/kevin-0305)
//@function: DeleteServer
//@description: 删除服务器
//@param: id float64
//@return: err error

func DeleteServer(id int) (err error) {
	var server model.MonitorServer
	err = global.GVA_DB.Where("id = ?", id).Delete(&server).Error
	return err
}
