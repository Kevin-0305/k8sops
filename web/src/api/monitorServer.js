import service from '@/utils/request'
// @Tags api
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取用户列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
// {
//  page     int
//	pageSize int
// }
export const getSeverList = (data) => {
    return service({
        url: `/monitorServer/getServerList/${data.page}/${data.pageSize}/`,
        method: 'get',
        data
    })
}


// @Tags Api
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateApiParams true "创建api"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/createApi [post]
export const createServer = (data) => {
    return service({
        url: "/monitorServer/createServer",
        method: 'post',
        data
    })
}

// @Tags menu
// @Summary 根据id获取菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.GetById true "根据id获取菜单"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getApiById [post]
export const getServerById = (data) => {
    return service({
        url: `/monitorServer/getServerInfo/${data.id}/`,
        method: 'get',
        data
    })
}



// @Tags Api
// @Summary 更新api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateApiParams true "更新api"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /api/updateApi [post]
export const updateServer = (data) => {
    console.log(data)
    return service({
        url: `/monitorServer/putServerInfo/${data.ID}/`,
        method: 'put',
        data
    })
}



// @Tags Api
// @Summary 删除指定api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.Api true "删除api"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/deleteApi [post]
export const deleteServer = (data) => {
    return service({
        url: `/monitorServer/deleteServer/${data.ID}/`,
        method: 'delete',
        data
    })
}
