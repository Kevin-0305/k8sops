import service from '@/utils/request'

// @Tags Service
// @Summary 创建Service
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Service true "创建Service"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /svc/createService [post]
export const createService = (data) => {
     return service({
         url: "/svc/createService",
         method: 'post',
         data
     })
 }


// @Tags Service
// @Summary 删除Service
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Service true "删除Service"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /svc/deleteService [delete]
 export const deleteService = (data) => {
     return service({
         url: "/svc/deleteService",
         method: 'delete',
         data
     })
 }

// @Tags Service
// @Summary 删除Service
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Service"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /svc/deleteService [delete]
 export const deleteServiceByIds = (data) => {
     return service({
         url: "/svc/deleteServiceByIds",
         method: 'delete',
         data
     })
 }

// @Tags Service
// @Summary 更新Service
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Service true "更新Service"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /svc/updateService [put]
 export const updateService = (data) => {
     return service({
         url: "/svc/updateService",
         method: 'put',
         data
     })
 }


// @Tags Service
// @Summary 用id查询Service
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Service true "用id查询Service"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /svc/findService [get]
 export const findService = (params) => {
     return service({
         url: "/svc/findService",
         method: 'get',
         params
     })
 }


// @Tags Service
// @Summary 分页获取Service列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Service列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /svc/getServiceList [get]
 export const getServiceList = (params) => {
     return service({
         url: "/svc/getServiceList",
         method: 'get',
         params
     })
 }