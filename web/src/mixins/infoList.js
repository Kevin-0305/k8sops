import { getDict } from "@/utils/dictionary";
import { getNamespacesList } from "@/api/namespaces"

export default {
    data() {
        return {
            page: 1,
            total: 10,
            pageSize: 10,
            tableData: [],
            searchInfo: {},
            namespace: "default",
            namespacesData: [],
            listNamespaces: getNamespacesList,
        }
    },
    methods: {
        filterDict(value, type) {
            const rowLabel = this[type + "Options"] && this[type + "Options"].filter(item => item.value == value)
            return rowLabel && rowLabel[0] && rowLabel[0].label
        },
        async getDict(type) {
            const dicts = await getDict(type)
            this[type + "Options"] = dicts
            return dicts
        },
        handleSizeChange(val) {
            this.pageSize = val
            this.getTableData()
        },
        handleCurrentChange(val) {
            this.page = val
            this.getTableData()
        },
        async getTableData(page = this.page, pageSize = this.pageSize) {
            const table = await this.listApi({ page, pageSize, ...this.searchInfo })
            if (table.code == 0) {
                this.tableData = table.data.list
                this.total = table.data.total
                this.page = table.data.page
                this.pageSize = table.data.pageSize
            }
        },
        async getNamespaceData(page = 1, pageSize = 999,clusterId=0){
            this.searchInfo = {
                "clusterId":clusterId
            }
            const table = await this.listNamespaces({ page, pageSize,...this.searchInfo})
            if (table.code == 0){
                this.namespacesData = table.data.list
            }
        }
    }
}