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

// @Tags Project
// @Summary 创建Project
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Project true "创建Project"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pro/createProject [post]
func CreateProject(c *gin.Context) {
	var pro model.Project
	_ = c.ShouldBindJSON(&pro)
	if err := service.CreateProject(pro); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Project
// @Summary 删除Project
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Project true "删除Project"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pro/deleteProject [delete]
func DeleteProject(c *gin.Context) {
	var pro model.Project
	_ = c.ShouldBindJSON(&pro)
	if err := service.DeleteProject(pro); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Project
// @Summary 批量删除Project
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Project"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /pro/deleteProjectByIds [delete]
func DeleteProjectByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteProjectByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Project
// @Summary 更新Project
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Project true "更新Project"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pro/updateProject [put]
func UpdateProject(c *gin.Context) {
	var pro model.Project
	_ = c.ShouldBindJSON(&pro)
	if err := service.UpdateProject(pro); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Project
// @Summary 用id查询Project
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Project true "用id查询Project"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pro/findProject [get]
func FindProject(c *gin.Context) {
	var pro model.Project
	_ = c.ShouldBindQuery(&pro)
	if err, repro := service.GetProject(pro.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repro": repro}, c)
	}
}

// @Tags Project
// @Summary 分页获取Project列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ProjectSearch true "分页获取Project列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pro/getProjectList [get]
func GetProjectList(c *gin.Context) {
	var pageInfo request.ProjectSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetProjectInfoList(pageInfo); err != nil {
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

// @Tags Project
// @Summary 运行构建
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Project true "用id运行构建任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"运行成功"}"
// @Router /pro/runBuild [get]
func RunBuild(c *gin.Context) {
	var pro model.Project
	_ = c.ShouldBindQuery(&pro)
	if err, repro := service.GetProject(pro.ID); err != nil {
		global.GVA_LOG.Error("构建失败项目不存在!", zap.Any("err", err))
		response.FailWithMessage("构建失败项目不存在", c)
	} else {
		task, err := service.CreateBuildTask(repro)
		if err != nil {
			global.GVA_LOG.Error("构建失败构建任务创建失败!", zap.Any("err", err))
			response.FailWithMessage("构建失败构建任务创建失败", c)
		}
		service.RunBuildTask(repro, task)
		response.OkWithData(gin.H{"repro": repro}, c)
	}
}
