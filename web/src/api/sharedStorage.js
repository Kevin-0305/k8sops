import service from '@/utils/request'

// @Tags SharedStorage
// @Summary 创建SharedStorage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SharedStorage true "创建SharedStorage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ShrSto/createSharedStorage [post]
export const createSharedStorage = (data) => {
     return service({
         url: "/ShrSto/createSharedStorage",
         method: 'post',
         data
     })
 }


// @Tags SharedStorage
// @Summary 删除SharedStorage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SharedStorage true "删除SharedStorage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ShrSto/deleteSharedStorage [delete]
 export const deleteSharedStorage = (data) => {
     return service({
         url: "/ShrSto/deleteSharedStorage",
         method: 'delete',
         data
     })
 }

// @Tags SharedStorage
// @Summary 删除SharedStorage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SharedStorage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ShrSto/deleteSharedStorage [delete]
 export const deleteSharedStorageByIds = (data) => {
     return service({
         url: "/ShrSto/deleteSharedStorageByIds",
         method: 'delete',
         data
     })
 }

// @Tags SharedStorage
// @Summary 更新SharedStorage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SharedStorage true "更新SharedStorage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ShrSto/updateSharedStorage [put]
 export const updateSharedStorage = (data) => {
     return service({
         url: "/ShrSto/updateSharedStorage",
         method: 'put',
         data
     })
 }


// @Tags SharedStorage
// @Summary 用id查询SharedStorage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SharedStorage true "用id查询SharedStorage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ShrSto/findSharedStorage [get]
 export const findSharedStorage = (params) => {
     return service({
         url: "/ShrSto/findSharedStorage",
         method: 'get',
         params
     })
 }


// @Tags SharedStorage
// @Summary 分页获取SharedStorage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SharedStorage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ShrSto/getSharedStorageList [get]
 export const getSharedStorageList = (params) => {
     return service({
         url: "/ShrSto/getSharedStorageList",
         method: 'get',
         params
     })
 }