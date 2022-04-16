package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
	"path"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateBuildTask
//@description: 创建BuildTask记录
//@param: buildTask model.BuildTask
//@return: err error

func CreateBuildTask(pro model.Project) (buildTask model.BuildTask, err error) {
	buildTask = model.BuildTask{
		ProjectId: int(pro.ID),
	}
	err = global.GVA_DB.Create(&buildTask).Error
	return buildTask, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteBuildTask
//@description: 删除BuildTask记录
//@param: buildTask model.BuildTask
//@return: err error

func DeleteBuildTask(buildTask model.BuildTask) (err error) {
	err = global.GVA_DB.Delete(&buildTask).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteBuildTaskByIds
//@description: 批量删除BuildTask记录
//@param: ids request.IdsReq
//@return: err error

func DeleteBuildTaskByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.BuildTask{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateBuildTask
//@description: 更新BuildTask记录
//@param: buildTask *model.BuildTask
//@return: err error

func UpdateBuildTask(buildTask model.BuildTask) (err error) {
	err = global.GVA_DB.Save(&buildTask).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBuildTask
//@description: 根据id获取BuildTask记录
//@param: id uint
//@return: err error, buildTask model.BuildTask

func GetBuildTask(id uint) (err error, buildTask model.BuildTask) {
	err = global.GVA_DB.Where("id = ?", id).First(&buildTask).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBuildTaskInfoList
//@description: 分页获取BuildTask记录
//@param: info request.BuildTaskSearch
//@return: err error, list interface{}, total int64

func GetBuildTaskInfoList(info request.BuildTaskSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.BuildTask{})
	var buildTasks []model.BuildTask
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ProjectId != 0 {
		db = db.Where("`project_id` = ?", info.ProjectId)
	}
	if info.Deploy != 0 {
		db = db.Where("`deploy` = ?", info.Deploy)
	}
	if info.BuildWay != 0 {
		db = db.Where("`build_way` = ?", info.BuildWay)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&buildTasks).Error
	return err, buildTasks, total
}

//@author: [kevin-0305](https://github.com/kevin-0305)
//@function: RunBulidTask
//@description: 执行构建任务
//@param: pro model.Project
//@return: err error, isSuccess bool
func RunBuildTask(pro model.Project, task model.BuildTask) {
	proWorkSpace := path.Join(global.GVA_PWD, "workspace", pro.Name)
	var logs string
	var logCh = make(chan string)
	var quitCh = make(chan bool)
	go utils.OutPutChan(logCh, quitCh)
	if ok, _ := utils.PathExists(proWorkSpace); !ok {
		outPut, err := utils.GitClone(pro, proWorkSpace, logCh)
		if err != nil {
			logs = logs + err.Error()
			panic(err)
		}
		logs = logs + outPut
	}
	outPut, err := utils.GitPull(pro, proWorkSpace, logCh)
	if err != nil {
		logs = logs + err.Error()
		panic(err)

	}
	logs = logs + outPut
	outPut, err = utils.CmdSync(pro.BuildScript, logCh)
	if err != nil {

		logs = logs + err.Error()
		panic(err)
	}
	quitCh <- true
	for i := range logCh {
		fmt.Println("log", i)
	}
}
