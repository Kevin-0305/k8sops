import service from '@/utils/request'

// @Tags BuildTask
// @Summary 创建BuildTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BuildTask true "创建BuildTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /buildTask/createBuildTask [post]
export const createBuildTask = (data) => {
     return service({
         url: "/buildTask/createBuildTask",
         method: 'post',
         data
     })
 }


// @Tags BuildTask
// @Summary 删除BuildTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BuildTask true "删除BuildTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /buildTask/deleteBuildTask [delete]
 export const deleteBuildTask = (data) => {
     return service({
         url: "/buildTask/deleteBuildTask",
         method: 'delete',
         data
     })
 }

// @Tags BuildTask
// @Summary 删除BuildTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BuildTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /buildTask/deleteBuildTask [delete]
 export const deleteBuildTaskByIds = (data) => {
     return service({
         url: "/buildTask/deleteBuildTaskByIds",
         method: 'delete',
         data
     })
 }

// @Tags BuildTask
// @Summary 更新BuildTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BuildTask true "更新BuildTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /buildTask/updateBuildTask [put]
 export const updateBuildTask = (data) => {
     return service({
         url: "/buildTask/updateBuildTask",
         method: 'put',
         data
     })
 }


// @Tags BuildTask
// @Summary 用id查询BuildTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BuildTask true "用id查询BuildTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /buildTask/findBuildTask [get]
 export const findBuildTask = (params) => {
     return service({
         url: "/buildTask/findBuildTask",
         method: 'get',
         params
     })
 }


// @Tags BuildTask
// @Summary 分页获取BuildTask列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取BuildTask列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /buildTask/getBuildTaskList [get]
 export const getBuildTaskList = (params) => {
     console.log(params)
     return service({
         url: "/buildTask/getBuildTaskList",
         method: 'get',
         params
     })
 }