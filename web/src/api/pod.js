import service from '@/utils/request'

// @Tags Pod
// @Summary 创建Pod
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Pod true "创建Pod"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /svc/createPod [post]
export const createPod = (data) => {
     return service({
         url: "/pod/createPod",
         method: 'post',
         data
     })
 }


// @Tags Pod
// @Summary 删除Pod
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Pod true "删除Pod"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /svc/deletePod [delete]
 export const deletePod = (data) => {
     return service({
         url: "/pod/deletePod",
         method: 'delete',
         data
     })
 }

// @Tags Pod
// @Summary 删除Pod
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Pod"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /svc/deletePod [delete]
 export const deletePodByIds = (data) => {
     return service({
         url: "/pod/deletePodByIds",
         method: 'delete',
         data
     })
 }

// @Tags Pod
// @Summary 更新Pod
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Pod true "更新Pod"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /svc/updatePod [put]
 export const updatePod = (data) => {
     return service({
         url: "/pod/updatePod",
         method: 'put',
         data
     })
 }


// @Tags Pod
// @Summary 用id查询Pod
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Pod true "用id查询Pod"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /svc/findPod [get]
 export const findPod = (params) => {
     return service({
         url: "/pod/findPod",
         method: 'get',
         params
     })
 }


// @Tags Pod
// @Summary 分页获取Pod列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Pod列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /svc/getPodList [get]
 export const getPodList = (params) => {
     return service({
         url: "/pod/getPodList",
         method: 'get',
         params
     })
 }