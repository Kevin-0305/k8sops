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

// @Tags SharedStorage
// @Summary 创建SharedStorage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SharedStorage true "创建SharedStorage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ShrSto/createSharedStorage [post]
func CreateSharedStorage(c *gin.Context) {
	var ShrSto model.SharedStorage
	_ = c.ShouldBindJSON(&ShrSto)
	if err := service.CreateSharedStorage(ShrSto); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SharedStorage
// @Summary 删除SharedStorage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SharedStorage true "删除SharedStorage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ShrSto/deleteSharedStorage [delete]
func DeleteSharedStorage(c *gin.Context) {
	var ShrSto model.SharedStorage
	_ = c.ShouldBindJSON(&ShrSto)
	if err := service.DeleteSharedStorage(ShrSto); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SharedStorage
// @Summary 批量删除SharedStorage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SharedStorage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ShrSto/deleteSharedStorageByIds [delete]
func DeleteSharedStorageByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteSharedStorageByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags SharedStorage
// @Summary 更新SharedStorage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SharedStorage true "更新SharedStorage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ShrSto/updateSharedStorage [put]
func UpdateSharedStorage(c *gin.Context) {
	var ShrSto model.SharedStorage
	_ = c.ShouldBindJSON(&ShrSto)
	if err := service.UpdateSharedStorage(ShrSto); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SharedStorage
// @Summary 用id查询SharedStorage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SharedStorage true "用id查询SharedStorage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ShrSto/findSharedStorage [get]
func FindSharedStorage(c *gin.Context) {
	var ShrSto model.SharedStorage
	_ = c.ShouldBindQuery(&ShrSto)
	if err, reShrSto := service.GetSharedStorage(ShrSto.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reShrSto": reShrSto}, c)
	}
}

// @Tags SharedStorage
// @Summary 分页获取SharedStorage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SharedStorageSearch true "分页获取SharedStorage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ShrSto/getSharedStorageList [get]
func GetSharedStorageList(c *gin.Context) {
	var pageInfo request.SharedStorageSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetSharedStorageInfoList(pageInfo); err != nil {
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
