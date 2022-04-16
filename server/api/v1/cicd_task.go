package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags BuildTask
// @Summary 创建BuildTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BuildTask true "创建BuildTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /buildTask/createBuildTask [post]

// func CreateBuildTask(c *gin.Context) {
// 	var buildTask model.BuildTask
// 	_ = c.ShouldBindJSON(&buildTask)
// 	if err := service.CreateBuildTask(buildTask); err != nil {
// 		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
// 		response.FailWithMessage("创建失败", c)
// 	} else {
// 		response.OkWithMessage("创建成功", c)
// 	}
// }

// @Tags BuildTask
// @Summary 删除BuildTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BuildTask true "删除BuildTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /buildTask/deleteBuildTask [delete]
func DeleteBuildTask(c *gin.Context) {
	var buildTask model.BuildTask
	_ = c.ShouldBindJSON(&buildTask)
	if err := service.DeleteBuildTask(buildTask); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags BuildTask
// @Summary 批量删除BuildTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BuildTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /buildTask/deleteBuildTaskByIds [delete]
func DeleteBuildTaskByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteBuildTaskByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags BuildTask
// @Summary 更新BuildTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BuildTask true "更新BuildTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /buildTask/updateBuildTask [put]
func UpdateBuildTask(c *gin.Context) {
	var buildTask model.BuildTask
	_ = c.ShouldBindJSON(&buildTask)
	if err := service.UpdateBuildTask(buildTask); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags BuildTask
// @Summary 用id查询BuildTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BuildTask true "用id查询BuildTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /buildTask/findBuildTask [get]
func FindBuildTask(c *gin.Context) {
	var buildTask model.BuildTask
	_ = c.ShouldBindQuery(&buildTask)
	if err, rebuildTask := service.GetBuildTask(buildTask.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebuildTask": rebuildTask}, c)
	}
}

// @Tags BuildTask
// @Summary 分页获取BuildTask列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.BuildTaskSearch true "分页获取BuildTask列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /buildTask/getBuildTaskList [get]
func GetBuildTaskList(c *gin.Context) {
	var pageInfo request.BuildTaskSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetBuildTaskInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
