<template>
  <div>
    <!-- 搜索组件 -->
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="公网IP">
          <el-input placeholder="公网IP" v-model="searchInfo.path"></el-input>
        </el-form-item>
        <el-form-item label="分组">
          <el-input placeholder="分组" v-model="searchInfo.description"></el-input>
        </el-form-item>
        <el-form-item label="系统">
           <el-select clearable placeholder="系统" v-model="searchInfo.system">
            <el-option
              :key="sysIndex"
              :label="sysVal"
              :value="sysKey"
              v-for="(sysVal,sysKey,sysIndex) in systemOptions"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select clearable placeholder="状态" v-model="searchInfo.state">
            <el-option
              :key="item.value"
              :label="item.label"
              :value="item.value"
              v-for="item in stateOptions"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog('add')" type="primary">新增服务器</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
              </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <!-- 列表表单 -->
    <el-table :data="tableData" @sort-change="sortChange" border stripe @selection-change="handleSelectionChange">
       <el-table-column
        type="selection"
        width="55">
      </el-table-column>
      <el-table-column label="id" min-width="60" prop="ID" sortable="custom"></el-table-column>
      <el-table-column label="公网IP" min-width="150" prop="pubIpAddress" sortable="custom"></el-table-column>
      <el-table-column label="内网IP" min-width="150" prop="pveIpAddress" sortable="custom"></el-table-column>
      <el-table-column label="端口" min-width="150" prop="sshPort" sortable="custom"></el-table-column>
      <el-table-column label="服务器名" min-width="150" prop="hostName" sortable="custom"></el-table-column>
      <el-table-column label="分组" min-width="150" prop="groupId" sortable="custom"></el-table-column>
      <el-table-column label="系统" min-width="150" prop="system" sortable="custom" :formatter="systemFiletr"></el-table-column>
      <el-table-column label="状态" min-width="150" prop="state" sortable="custom">
        <template slot-scope="scope">
          <div>
            <!-- {{scope.row.state|stateFiletr}} -->
            <el-tag
              :key="scope.row.stateFiletr"
              :type="scope.row.state|tagTypeFiletr"
              effect="dark"
              size="mini"
            >{{scope.row.state|stateFiletr}}</el-tag>
            <!-- {{scope.row.method|methodFiletr}} -->
          </div>
        </template>
      </el-table-column>

      <!-- 操作按钮 -->
      <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
          <el-button @click="editServer(scope.row)" size="small" type="primary" icon="el-icon-edit">编辑</el-button>
          <el-button
            @click="deleteServer(scope.row)"
            size="small"
            type="danger"
            icon="el-icon-delete"
          >删除</el-button>
          <el-button @click="openWebSSH(scope.row)" size="small" type="primary" icon="el-icon-edit">SSH</el-button>
        </template>
      </el-table-column>

    </el-table>
    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <!-- 弹窗组件 -->
    <el-dialog :before-close="closeDialog" :title="dialogTitle" :visible.sync="dialogFormVisible">
      <el-form :inline="true" :model="form" :rules="rules" label-width="80px" ref="apiForm">

        <el-form-item label="公网IP" prop="pubIpAddress">
          <el-input autocomplete="off" v-model="form.pubIpAddress"></el-input>
        </el-form-item>

        <el-form-item label="内网IP" prop="pveIpAddress">
          <el-input autocomplete="off" v-model="form.pveIpAddress"></el-input>
        </el-form-item>

        <el-form-item label="IPv6" prop="ipV6Address">
          <el-input autocomplete="off" v-model="form.ipV6Address"></el-input>
        </el-form-item>

        <el-form-item label="SSH端口" prop="sshPort">
          <el-input autocomplete="off" v-model="form.sshPort" type='number'></el-input>
        </el-form-item>

        <el-form-item label="主机名" prop="hostName">
          <el-input autocomplete="off" v-model="form.hostName"></el-input>
        </el-form-item>

        <el-form-item label="登录账号" prop="user">
          <el-input autocomplete="off" v-model="form.user"></el-input>
        </el-form-item>

        <el-form-item label="密码" prop="password">
          <el-input autocomplete="off" v-model="form.password"></el-input>
        </el-form-item>

        <el-form-item label="分组" prop="groupId">
          <el-select placeholder="请选择" v-model="form.groupId">
            <el-option
              :key="index"
              :label="val"
              :value="key"
              v-for="(val,key,index) in systemOptions"
            ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="ssh文件" prop="keyFile">
          <el-input autocomplete="off" v-model="form.keyFile"></el-input>
        </el-form-item>

        <el-form-item label="服务器类型" prop="serverType">
          <el-select placeholder="请选择" v-model="form.serverType">
            <el-option
              :key="index"
              :label="val"
              :value="key"
              v-for="(val,key,index) in systemOptions"
            ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="系统" prop="system">
          <el-select placeholder="请选择" v-model="form.system">
            <el-option
              :key="index"
              :label="val"
              :value="key"
              v-for="(val,key,index) in systemOptions"
            ></el-option>
          </el-select>
        </el-form-item>


      </el-form>
      <div class="warning">添加服务器需要在角色管理内配置权限才可使用</div>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成 条件搜索时候 请把条件安好后台定制的结构体字段 放到 this.searchInfo 中即可实现条件搜索

import {
  deleteApisByIds
} from "@/api/api";
import {
  getSeverList,
  createServer,
  updateServer,
  getServerById,
  deleteServer
  
} from "@/api/monitorServer"
import infoList from "@/mixins/infoList";
import { toSQLLine } from "@/utils/stringFun";
const systemOptions = {
  0:"unknow",
  1:"Centos7",
  2:"Centos8",
  3:"Ubuntu",
  4:"Debian"
  
}
// const systemOptions = [
//   {
//     value: 1,
//     label: "CentOS7",
//   },
//   {
//     value: 2,
//     label: "CentOS8",
//   },
//   {
//     value: 3,
//     label: "Ubuntu",
//   },
//   {
//     value: 4,
//     label: "Debian",
//   }
// ];
const stateOptions = [
  
  {
    value: 1,
    label: "开启",
    colorType: "success"
  },
  {
    value: 2,
    label: "关闭",
    colorType: "danger"
  },
  {
    value: 0,
    label: "未知",
    colorType: "warning"
  }

];

export default {
  name: "Server",
  mixins: [infoList],
  data() {
    return {
      deleteVisible:false,
      listApi: getSeverList,
      dialogFormVisible: false,
      dialogTitle: "添加服务器",
      apis:[],
      form: {
        id:0,
        pubIpAddress:"",
        pveIpAddress:"",
        ipV6Address:"",
        sshPort:0,
        hostName:"",
        user:"",
        password:"",
        groupId:0,
        keyFile:"",
        state:0,
        serverType:0,
        system:0,
        
      },
      stateOptions:stateOptions,
      systemOptions:systemOptions,
      type: "",
      rules: {

        path: [{ required: true, message: "请输入api路径", trigger: "blur" }],
        apiGroup: [
          { required: true, message: "请输入组名称", trigger: "blur" }
        ],
        method: [
          { required: true, message: "请选择请求方式", trigger: "blur" }
        ],
        description: [
          { required: true, message: "请输入api介绍", trigger: "blur" }
        ]
      }
    };
  },
  methods: {
    //  选中api
      handleSelectionChange(val) {
        this.apis = val;
      },
      async onDelete(){
        const ids = this.apis.map(item=>item.ID)
        const res = await deleteApisByIds({ids})
        if(res.code==0){
          this.$message({
            type:"success",
            message:res.msg
          })
         if (this.tableData.length == ids.length) {
              this.page--;
          }
          this.deleteVisible = false
          this.getTableData()
        }
      },
    // 排序
    sortChange({ prop, order }) {
      if (prop) {
        this.searchInfo.orderKey = toSQLLine(prop);
        this.searchInfo.desc = order == "descending";
      }
      this.page = 1;
      this.pageSize = 10;
      this.getTableData();
    },
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1;
      this.pageSize = 10;
      this.getTableData();
    },
    initForm() {
      this.$refs.apiForm.resetFields();
      this.form = {
        path: "",
        apiGroup: "",
        method: "",
        description: ""
      };
    },
    closeDialog() {
      this.initForm();
      this.dialogFormVisible = false;
    },
    openDialog(type) {
      switch (type) {
        case "add":
          this.dialogTitle = "新增服务器";
          break;
        case "edit":
          this.dialogTitle = "修改服务器信息";
          break;
        default:
          break;
      }
      this.type = type;
      this.dialogFormVisible = true;
    },
    async editServer(row) {
      const res = await getServerById({ id: row.ID });
      console.log(res)
      this.form = res.data.server;
      this.openDialog("edit");
    },
    async deleteServer(row) {
      this.$confirm("此操作将删除此条服务器信息, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(async () => {
          const res = await deleteServer(row);
          if (res.code == 0) {
            this.$message({
              type: "success",
              message: "删除成功!"
            });
            if (this.tableData.length == 1) {
              this.page--;
            }
            this.page = 1;
            this.pageSize = 10;
            this.getTableData();
          }
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除"
          });
        });
    },
    async openWebSSH(row){

      let routeUrl = this.$router.resolve({
          path: "/webssh",
          query: {id:row.ID,name:row.hostName}
     });
     window.open(routeUrl .href, '_blank');
    },
    //系统映射函数
    systemFiletr(row){
      var label = ""
      label =  systemOptions[row.system]
      return label;
    },
    async enterDialog() {
      this.$refs.apiForm.validate(async valid => {
        if (valid) {
          switch (this.type) {
            case "add":
              {
                const res = await createServer(this.form);
                if (res.code == 0) {
                  this.$message({
                    type: "success",
                    message: "添加成功",
                    showClose: true
                  });
                }
                this.getTableData();
                this.closeDialog();
              }

              break;
            case "edit":
              {
                this.form.sshPort = parseInt(this.form.sshPort)
                const res = await updateServer(this.form);
                if (res.code == 0) {
                  this.$message({
                    type: "success",
                    message: "编辑成功",
                    showClose: true
                  });
                }
                this.getTableData();
                this.closeDialog();
              }
              break;
            default:
              {
                this.$message({
                  type: "error",
                  message: "未知操作",
                  showClose: true
                });
              }
              break;
          }
        }
      });
    }
  },
  filters: {
    // 状态过滤器
    stateFiletr(value){
      
      const item = stateOptions.filter(item => item.value === value)[0];
      return item.label;

    },
    methodFiletr(value) {
      const target = stateOptions.filter(item => item.value === value)[0];
      // return target && `${target.label}(${target.value})`
      return target && `${target.label}`;
    },
    tagTypeFiletr(value) {
      const target = stateOptions.filter(item => item.value === value)[0];
      return target && `${target.colorType}`;
    }
  },
  created() {
    this.getTableData();
  }
};
</script>
<style scoped lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}
.el-tag--mini {
  margin-left: 5px;
}
.warning {
  color: #dc143c;
}
</style>