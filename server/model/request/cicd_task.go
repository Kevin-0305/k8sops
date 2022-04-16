package request

import "gin-vue-admin/model"

type BuildTaskSearch struct{
    model.BuildTask
    PageInfo
}