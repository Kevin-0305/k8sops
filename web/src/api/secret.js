import service from '@/utils/request'

// @Tags Secret
// @Summary 创建Secret
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Secret true "创建Secret"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sct/createSecret [post]
export const createSecret = (data) => {
     return service({
         url: "/sct/createSecret",
         method: 'post',
         data
     })
 }


// @Tags Secret
// @Summary 删除Secret
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Secret true "删除Secret"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sct/deleteSecret [delete]
 export const deleteSecret = (data) => {
     return service({
         url: "/sct/deleteSecret",
         method: 'delete',
         data
     })
 }

// @Tags Secret
// @Summary 删除Secret
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Secret"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sct/deleteSecret [delete]
 export const deleteSecretByIds = (data) => {
     return service({
         url: "/sct/deleteSecretByIds",
         method: 'delete',
         data
     })
 }

// @Tags Secret
// @Summary 更新Secret
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Secret true "更新Secret"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sct/updateSecret [put]
 export const updateSecret = (data) => {
     return service({
         url: "/sct/updateSecret",
         method: 'put',
         data
     })
 }


// @Tags Secret
// @Summary 用id查询Secret
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Secret true "用id查询Secret"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sct/findSecret [get]
 export const findSecret = (params) => {
     return service({
         url: "/sct/findSecret",
         method: 'get',
         params
     })
 }


// @Tags Secret
// @Summary 分页获取Secret列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Secret列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sct/getSecretList [get]
 export const getSecretList = (params) => {
     return service({
         url: "/sct/getSecretList",
         method: 'get',
         params
     })
 }