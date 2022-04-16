import service from '@/utils/request'

// @Tags PersistentVolume
// @Summary 创建PersistentVolume
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PersistentVolume true "创建PersistentVolume"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /svc/createPersistentVolume [post]
export const createPersistentVolume = (data) => {
     return service({
         url: "/pv/createPersistentVolume",
         method: 'post',
         data
     })
 }


// @Tags PersistentVolume
// @Summary 删除PersistentVolume
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PersistentVolume true "删除PersistentVolume"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /svc/deletePersistentVolume [delete]
 export const deletePersistentVolume = (data) => {
     return service({
         url: "/pv/deletePersistentVolume",
         method: 'delete',
         data
     })
 }

// @Tags PersistentVolume
// @Summary 删除PersistentVolume
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PersistentVolume"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /svc/deletePersistentVolume [delete]
 export const deletePersistentVolumeByIds = (data) => {
     return service({
         url: "/pv/deletePersistentVolumeByIds",
         method: 'delete',
         data
     })
 }

// @Tags PersistentVolume
// @Summary 更新PersistentVolume
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PersistentVolume true "更新PersistentVolume"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /svc/updatePersistentVolume [put]
 export const updatePersistentVolume = (data) => {
     return service({
         url: "/pv/updatePersistentVolume",
         method: 'put',
         data
     })
 }


// @Tags PersistentVolume
// @Summary 用id查询PersistentVolume
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PersistentVolume true "用id查询PersistentVolume"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /svc/findPersistentVolume [get]
 export const findPersistentVolume = (params) => {
     return service({
         url: "/pv/findPersistentVolume",
         method: 'get',
         params
     })
 }


// @Tags PersistentVolume
// @Summary 分页获取PersistentVolume列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取PersistentVolume列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /svc/getPersistentVolumeList [get]
 export const getPersistentVolumeList = (params) => {
     return service({
         url: "/pv/getPersistentVolumeList",
         method: 'get',
         params
     })
 }