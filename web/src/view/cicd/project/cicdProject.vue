<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="项目">
          <el-input placeholder="搜索条件" v-model="searchInfo.name"></el-input>
        </el-form-item>      
        <el-form-item label="语言">
          <el-input placeholder="搜索条件" v-model="searchInfo.language"></el-input>
        </el-form-item>                    
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">添加工程</el-button>
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
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="selection" width="55"></el-table-column>
    <el-table-column label="创建日期" width="250">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>
    
    <el-table-column label="项目名" prop="name" width="250"></el-table-column> 
    
    <el-table-column label="说明" prop="description" width="250"></el-table-column> 
    
      <el-table-column label="语言" prop="language" width="120">
        <template slot-scope="scope">
          {{filterDict(scope.row.language,"int")}}
        </template>
      </el-table-column>
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateProject(scope.row)" size="small" type="primary" icon="el-icon-edit">编辑</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="openBuildTask(scope.row)">构建历史</el-button>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
         <el-form-item label="项目名:">
            <el-input v-model="formData.name" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="说明:">
            <el-input v-model="formData.description" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="语言:">
             <el-select v-model="formData.language" placeholder="请选择" clearable>
                 <el-option 
                 :key="index"
                 :label="val"
                 :value="key"
                 v-for="(val,key,index) in languageOptions">
                 </el-option>
             </el-select>
      </el-form-item>
       
         <el-form-item label="仓库地址:">
            <el-input v-model="formData.wareHouse" clearable placeholder="请输入" ></el-input>
         </el-form-item>
       
         <el-form-item label="仓库账号:">
            <el-input v-model="formData.wareHouseAccout" clearable placeholder="请输入" ></el-input>
         </el-form-item>
       
         <el-form-item label="仓库密码:">
            <el-input v-model="formData.wareHousePassword" clearable placeholder="请输入" ></el-input>
         </el-form-item>

        <el-form-item label="构建脚本">
          <el-input v-model="formData.buildScript" type="textarea" placeholder="构建脚本"
            :autosize="{minRows: 1, maxRows: 2000}" :style="{width: '100%',height: '300%'}"></el-input>
        </el-form-item>
       
        <el-form-item label="测试脚本:">
            <el-input v-model="formData.testScript" clearable placeholder="请输入" ></el-input>
        </el-form-item>
       
        <el-form-item label="服务器组:">
            <el-input v-model="formData.servers" clearable placeholder="请输入" ></el-input>
        </el-form-item>

        <el-form-item label="生产环境中运行的脚本">
          <el-input v-model="formData.runScript" type="textarea" placeholder="写入部署至生产环境中运行的脚本"
           :autosize="{minRows: 1, maxRows: 2000}"
           :style="{width: '100%',height: 300}"></el-input>
        </el-form-item>
       
        <el-form-item label="工程目录:">
            <el-input v-model="formData.deploymentDir" clearable placeholder="请输入" ></el-input>
        </el-form-item>
       </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
    createProject,
    deleteProject,
    deleteProjectByIds,
    updateProject,
    findProject,
    getProjectList
} from "@/api/cicdProject";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "Project",
  mixins: [infoList],
  data() {
    return {
      listApi: getProjectList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      languageOptions:{
        1:"js",
        2:"python",
        3:"go"
      },
      formData: {
        name:"",
        description:"",
        language:0,
        wareHouse:"",
        wareHouseAccout:"",
        wareHousePassword:"",
        buildScript:"",
        testScript:"",
        servers:"",
        deploymentDir:"",
        runScript:"",
        
      }
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "是" :"否";
      } else {
        return "";
      }
    }
  },
  methods: {
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10               
        this.getTableData()
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      deleteRow(row){
        this.$confirm('确定要删除吗?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
           this.deleteProject(row);
        });
      },
      async onDelete() {
        const ids = []
        if(this.multipleSelection.length == 0){
          this.$message({
            type: 'warning',
            message: '请选择要删除的数据'
          })
          return
        }
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
        const res = await deleteProjectByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          if (this.tableData.length == ids.length) {
              this.page--;
          }
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateProject(row) {
      const res = await findProject({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.repro;
        this.dialogFormVisible = true;
      }
    },
    async openBuildTask(row){
      this.$router.push( {path:"buildTask",query:{projectId: row.ID}});
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          name:"",
          description:"",
          language:0,
          wareHouse:"",
          wareHouseAccout:"",
          wareHousePassword:"",
          buildScript:"",
          testScript:"",
          servers:"",
          deploymentDir:"",
          runScript:"",
          
      };
    },
    async deleteProject(row) {
      const res = await deleteProject({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        if (this.tableData.length == 1) {
            this.page--;
        }
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createProject(this.formData);
          break;
        case "update":
          res = await updateProject(this.formData);
          break;
        default:
          res = await createProject(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();
  
    await this.getDict("int");
    
}
};
</script>

<style>
</style>
