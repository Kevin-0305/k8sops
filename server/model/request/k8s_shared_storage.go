package request

import "gin-vue-admin/model"

type SharedStorageSearch struct{
    model.SharedStorage
    PageInfo
}