import service from '@/utils/request'

// @Tags PersistentVolumeClaim
// @Summary 创建PersistentVolumeClaim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PersistentVolumeClaim true "创建PersistentVolumeClaim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /svc/createPersistentVolumeClaim [post]
export const createPersistentVolumeClaim = (data) => {
     return service({
         url: "/pvc/createPersistentVolumeClaim",
         method: 'post',
         data
     })
 }


// @Tags PersistentVolumeClaim
// @Summary 删除PersistentVolumeClaim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PersistentVolumeClaim true "删除PersistentVolumeClaim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /svc/deletePersistentVolumeClaim [delete]
 export const deletePersistentVolumeClaim = (data) => {
     return service({
         url: "/pvc/deletePersistentVolumeClaim",
         method: 'delete',
         data
     })
 }

// @Tags PersistentVolumeClaim
// @Summary 删除PersistentVolumeClaim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PersistentVolumeClaim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /svc/deletePersistentVolumeClaim [delete]
 export const deletePersistentVolumeClaimByIds = (data) => {
     return service({
         url: "/pvc/deletePersistentVolumeClaimByIds",
         method: 'delete',
         data
     })
 }

// @Tags PersistentVolumeClaim
// @Summary 更新PersistentVolumeClaim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PersistentVolumeClaim true "更新PersistentVolumeClaim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /svc/updatePersistentVolumeClaim [put]
 export const updatePersistentVolumeClaim = (data) => {
     return service({
         url: "/pvc/updatePersistentVolumeClaim",
         method: 'put',
         data
     })
 }


// @Tags PersistentVolumeClaim
// @Summary 用id查询PersistentVolumeClaim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PersistentVolumeClaim true "用id查询PersistentVolumeClaim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /svc/findPersistentVolumeClaim [get]
 export const findPersistentVolumeClaim = (params) => {
     return service({
         url: "/pvc/findPersistentVolumeClaim",
         method: 'get',
         params
     })
 }


// @Tags PersistentVolumeClaim
// @Summary 分页获取PersistentVolumeClaim列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取PersistentVolumeClaim列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /svc/getPersistentVolumeClaimList [get]
 export const getPersistentVolumeClaimList = (params) => {
     return service({
         url: "/pvc/getPersistentVolumeClaimList",
         method: 'get',
         params
     })
 }