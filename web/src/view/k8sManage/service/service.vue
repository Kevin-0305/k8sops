<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="Name">
          <el-input placeholder="搜索条件" v-model="searchInfo.name"></el-input>
        </el-form-item>
        <el-select v-model="namespace" placeholder="请选择namespace" @change="namespaceChange">
            <el-option v-for="item in namespacesData" :key="item.metadata.name" :label="item.metadata.name" :value="item.metadata.name"/>
        </el-select>  
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增Service</el-button>
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
    <el-table-column label="创建日期" width="180">
         <template slot-scope="scope">{{scope.row.metadata.creationTimestamp|formatDate}}</template>
    </el-table-column>
    
    <el-table-column label="Name" prop="metadata.name" width="120"></el-table-column> 
    <el-table-column label="type" prop="spec.type" width="120"></el-table-column>
    <el-table-column label="clusterIP" prop="spec.clusterIP" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateService(scope.row)" size="small" type="primary" icon="el-icon-edit">修改</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
          <el-button type="table-button" size="small" @click="showConfig(scope.row)">查看配置文件</el-button>
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

    
    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" :title="dialogName" width="1600px">
      <el-form :model="formData" label-position="right" label-width="120px">
        <el-form-item label="Name:">
            <el-input v-model="formData.metadata.name" clearable placeholder="请输入" style="width:200px;"></el-input>
        </el-form-item>
        <el-form-item label="Namespace:">
            <el-select v-model="formData.metadata.namespace" placeholder="请选择namespace" @change="namespaceChange" :disabled="namespaceButton">
            <el-option v-for="item in namespacesData" :key="item.metadata.name" :label="item.metadata.name" :value="item.metadata.name"/>
        </el-select>
        </el-form-item>
          <el-form-item label="type:">
            <el-select v-model="formData.spec.type" placeholder="请选择ServiceType" @change="typeChange($event)">
              <el-option v-for="item in typeList" :key="item" :label="item" :value="item"/>
            </el-select>
        </el-form-item>
        <el-form-item label="externalName:" v-show="externalNameInput">
          <el-input v-model="formData.metadata.name" clearable placeholder="请输入" style="width:500px;" ></el-input>
        </el-form-item>
      </el-form>


      <dictTable :dataList.sync="annotations" comName="Annotations"></dictTable>
      <dictTable :dataList.sync="labels" comName="Labels"></dictTable>

      <el-form :inline="true">
        <el-link target="_blank" type="primary">posts</el-link>
        <el-button
          size="small"
          type="primary"
          icon="el-icon-edit"
          @click="addPort(ports)"
          style="float:right"
        >新增</el-button>
      </el-form>
      <el-table :data="ports" stripe style="width: 100%">
          <el-table-column  prop="name" label="name" width="180" >
            <template slot-scope="scope">
              <div>
                <el-input v-model="scope.row.name" :placeholder="scope.row.name"></el-input>
              </div>
            </template>
          </el-table-column>

          <el-table-column  prop="protocol" label="protocol" width="180" >
            <template slot-scope="scope">
              <el-select v-model="scope.row.protocol" placeholder="请选择协议类型" >
                <el-option v-for="item in protocolList" :key="item" :label="item" :value="item"/>
              </el-select>
            </template>
          </el-table-column>

          <el-table-column prop="port" label="port" width="180">
            <template slot-scope="scope">
              <div>
                <el-input-number v-model="scope.row.port" :placeholder="scope.row.port" ></el-input-number>
              </div>
            </template>
          </el-table-column>

          <el-table-column  prop="targetPort" label="targetPort" width="180">
            <template slot-scope="scope">
              <div>
                <el-input-number v-model="scope.row.targetPort" :placeholder="scope.row.targetPort"></el-input-number>
              </div>
            </template>
          </el-table-column>

          <el-table-column  prop="nodePort" label="nodePort" width="180" >
            <template slot-scope="scope">
              <div>
                <el-input-number v-model="scope.row.nodePort" :placeholder="scope.row.nodePort"></el-input-number>
              </div>
            </template>
          </el-table-column>

          <el-table-column>
            <template slot-scope="scope">
              <div>
                <el-button
                  type="danger"
                  size="small"
                  icon="el-icon-delete"
                  @click="deletePort(ports,scope.$index)"
                >删除</el-button>
              </div>
            </template>
          </el-table-column>
        </el-table>

      <dictTable :dataList.sync="selectors" comName="Selectors"></dictTable>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
        <el-button @click="openDrawer" type="primary">查看配置文件</el-button>
      </div>
    </el-dialog>
    <k8sDrawer  :drawer.sync="drawer" :formDataJson="formDataJson" :formDataYaml="formDataYaml" ></k8sDrawer>
  </div>
</template>

<script>
import {
    createService,
    deleteService,
    deleteServiceByIds,
    updateService,
    findService,
    getServiceList
} from "@/api/service";  //  此处请自行替换地址


import Menus from "@/view/superAdmin/authority/components/menus";
import Apis from "@/view/superAdmin/authority/components/apis";
import Datas from "@/view/superAdmin/authority/components/datas";
import k8sDrawer from '@/components/k8sDrawer/index.vue';
import dictTable from '@/components/dictTable/index.vue';

import { formatTimeToStr } from "@/utils/date";
import {listToDict} from "@/utils/dictToList";
import infoList from "@/mixins/infoList";

import yaml from 'json2yaml';

export default {
  name: "Service",
  mixins: [infoList],
  data() {
    return {
      listApi: getServiceList,
      namespace: "",
      dialogName: "",
      namespaceButton: false,
      dialogFormVisible: false,
      drawer: false,
      type: "",
      ports: [],
      deleteVisible: false,
      multipleSelection: [],
      externalNameInput:false,
      formData: {
            apiVersion: "v1",
            kind: "Service",
            metadata:{},
            spec: {
              sessionAffinity: "None",
              type: "ClusterIP"
            },
            status: {},
      },
      annotations:[],
      labels:[],
      selectors:[],
      typeList: ["ClusterIP","NodePort","ExternalName","LoadBalancer"],
      protocolList: ["TCP","UDP","SCTP"],
      reqData: {
        name:"",
        namespace:"",
        data: {},
      },
      formDataJson:"",
      formDataYaml:"",
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
  components: {
    Menus,
    Apis,
    Datas,
    yaml,
    k8sDrawer,
    dictTable,
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
           this.deleteService(row);
        });
      },
      deletePort(parameters, index) {
        parameters.splice(index, 1);
      },
      showConfig(row){
        let data = JSON.parse(JSON.stringify(this.formData));
        data.metadata = row.metadata;
        data.spec = row.spec;
        data.status = row.status;
        delete data.metadata["managedFields"]
        this.formDataJson = JSON.stringify(data,null, 4);
        this.formDataYaml = yaml.stringify(data);
        this.drawer = true;
      },
      openDrawer(){
        let data = JSON.parse(JSON.stringify(this.formData))
        data.metadata.annotations = listToDict(this.annotations)
        data.metadata.labels = listToDict(this.labels);
        data.spec.selector = listToDict(this.selectors)
        for (const i of this.ports){
          if ( i.port ==0 || i.targetPort == 0){
            this.$message.error("请检查port配置")
            return 
          }else{
            if (data.spec.type =="NodePort" && i.nodePort ==0) {
              this.$message.error("请检查port配置")
              return 
            }
          }
        }
        data.spec.ports = JSON.parse(JSON.stringify(this.ports))
        if (data.spec.type !="NodePort"){
          for (const i of data.spec.ports){
            delete i["nodePort"]
          }
        }
        this.formDataJson = JSON.stringify(data,null, 4);
        this.formDataYaml = yaml.stringify(data);
        this.drawer = true;
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
        const res = await deleteServiceByIds({ ids })
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
    async updateService(row) {
      const res = await findService({ namespace: row.metadata.namespace,name: row.metadata.name });
      if (res.code != 0){
          this.$message({
            type: 'error',
            message: '请求失败'
          });
          return;
      }
      this.type = "update";
      this.annotations = [];
      this.labels = [];
      this.selectors = [];
      
      delete res.data.reData.metadata["managedFields"];
      this.formData.metadata = res.data.reData.metadata;
      this.formData.spec = res.data.reData.spec;
      this.ports = res.data.reData.spec.ports
      for (var i in res.data.reData.metadata.annotations){
        this.annotations.push({key:i,value:res.data.reData.metadata.annotations[i]})
      };
      for (var i in res.data.reData.spec.selector){
        this.selectors.push({key:i,value:res.data.reData.spec.selector[i]})
      };
      for (var i in res.data.reData.metadata.labels){
        this.labels.push({key:i,value:res.data.reData.metadata.labels[i]})
      };

      this.namespaceButton = true;
      this.dialogName = "修改";
      this.dialogFormVisible = true;
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData.metadata = {};
      this.ports = [];
    },
    namespaceChange(){
      this.searchInfo = {
      "namespace":this.namespace,
      },
      this.getTableData();
    },
    typeChange(event){
      if (event == "ExternalName") {
      this.externalNameInput =true;
      }else{
        this.externalNameInput =false;
      }
    },
    addPort(portList){
      portList.push({
        name:"",
        protocol:"",
        port:0,
        targetPort:0,
        nodePort:0,
        }
      )
    },
    async deleteService(row) {
      const res = await deleteService({namespace: row.metadata.namespace,name: row.metadata.name });
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
      let data;
      data = JSON.parse(JSON.stringify(this.formData));
      data.spec.ports = this.ports 
      this.reqData.name = this.formData.metadata.name;
      this.reqData.namespace = this.formData.metadata.namespace;
      this.reqData.configData = JSON.stringify(data);
      switch (this.type) {
        case "create":
          res = await createService(this.reqData);
          break;
        case "update":
          res = await updateService(this.reqData);
          break;
        default:
          res = await createService(this.reqData);
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
      this.dialogName = "新增"
      this.formData.namespace = this.namespace;
      this.annotations = [];
      this.labels = [];
      this.selectors = [];
      this.namespaceButton = false;
      this.dialogFormVisible = true;
    }
  },
  async created() {
    this.searchInfo = {
      "namespace":"default"
    }
    await this.getTableData();
    await this.getNamespaceData();
  }
}
</script>

<style>
</style>



