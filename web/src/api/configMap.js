import service from '@/utils/request'

// @Tags ConfigMap
// @Summary 创建ConfigMap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ConfigMap true "创建ConfigMap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cm/createConfigMap [post]
export const createConfigMap = (data) => {
     return service({
         url: "/cm/createConfigMap",
         method: 'post',
         data
     })
 }


// @Tags ConfigMap
// @Summary 删除ConfigMap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ConfigMap true "删除ConfigMap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cm/deleteConfigMap [delete]
 export const deleteConfigMap = (data) => {
     return service({
         url: "/cm/deleteConfigMap",
         method: 'delete',
         data
     })
 }

// @Tags ConfigMap
// @Summary 删除ConfigMap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ConfigMap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cm/deleteConfigMap [delete]
 export const deleteConfigMapByIds = (data) => {
     return service({
         url: "/cm/deleteConfigMapByIds",
         method: 'delete',
         data
     })
 }

// @Tags ConfigMap
// @Summary 更新ConfigMap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ConfigMap true "更新ConfigMap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cm/updateConfigMap [put]
 export const updateConfigMap = (data) => {
     return service({
         url: "/cm/updateConfigMap",
         method: 'put',
         data
     })
 }


// @Tags ConfigMap
// @Summary 用id查询ConfigMap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ConfigMap true "用id查询ConfigMap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cm/findConfigMap [get]
 export const findConfigMap = (params) => {
     return service({
         url: "/cm/findConfigMap",
         method: 'get',
         params
     })
 }


// @Tags ConfigMap
// @Summary 分页获取ConfigMap列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取ConfigMap列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cm/getConfigMapList [get]
 export const getConfigMapList = (params) => {
     return service({
         url: "/cm/getConfigMapList",
         method: 'get',
         params
     })
 }