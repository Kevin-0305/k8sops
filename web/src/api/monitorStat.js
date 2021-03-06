import service from "@/utils/request";


// @Tags system
// @Summary 获取服务器运行状态
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /system/getServerInfo [post]
export const getServerStatList = () => {
    return service({
        url: "/monitorServer/getServerStatList",
        method: "get",
        donNotShowLoading: true
    });
};