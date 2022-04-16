<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">     
        <el-form-item label="部署环境">
          <el-input placeholder="搜索条件" v-model="searchInfo.deploy"></el-input>
        </el-form-item>    
        <el-form-item label="构建方式">
          <el-input placeholder="搜索条件" v-model="searchInfo.buildWay"></el-input>
        </el-form-item>          
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
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
    <el-table-column label="日期" width="180">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>
    
    <el-table-column label="工程ID" prop="projectId" width="180"></el-table-column> 
    
    <el-table-column label="描述" prop="description" width="180"></el-table-column> 
    
    <el-table-column label="部署环境" prop="deploy" width="180"></el-table-column> 
    
    <el-table-column label="构建方式" prop="buildWay" width="180"></el-table-column> 
    
    <el-table-column label="状态" prop="state" width="180"></el-table-column> 
    
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateBuildTask(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
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
         <el-form-item label="工程ID:"><el-input v-model.number="formData.projectId" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="描述:">
            <el-input v-model="formData.description" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="部署环境:"><el-input v-model.number="formData.deploy" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="构建方式:"><el-input v-model.number="formData.buildWay" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="状态:"><el-input v-model.number="formData.state" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="构建日志:">
            <el-input v-model="formData.buildLog" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="测试日志:">
            <el-input v-model="formData.testLog" clearable placeholder="请输入" ></el-input>
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
    createBuildTask,
    deleteBuildTask,
    deleteBuildTaskByIds,
    updateBuildTask,
    findBuildTask,
    getBuildTaskList
} from "@/api/cicd_task";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";

export default {
  name: "BuildTask",
  mixins: [infoList],
  data() {
    return {
      listApi: getBuildTaskList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            projectId:0,
            description:"",
            deploy:0,
            buildWay:0,
            state:0,
            buildLog:"",
            testLog:"",
            
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
           this.deleteBuildTask(row);
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
        const res = await deleteBuildTaskByIds({ ids })
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
    async updateBuildTask(row) {
      const res = await findBuildTask({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.rebuildTask;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          projectId:0,
          description:"",
          deploy:0,
          buildWay:0,
          state:0,
          buildLog:"",
          testLog:"",
          
      };
    },
    async deleteBuildTask(row) {
      const res = await deleteBuildTask({ ID: row.ID });
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
          res = await createBuildTask(this.formData);
          break;
        case "update":
          res = await updateBuildTask(this.formData);
          break;
        default:
          res = await createBuildTask(this.formData);
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
    this.searchInfo.projectId=this.$route.query.projectId
    this.getTableData();
}
};
</script>

<style>
</style>
