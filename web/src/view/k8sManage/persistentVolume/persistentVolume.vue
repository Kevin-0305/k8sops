<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="Name">
          <el-input placeholder="搜索条件" v-model="searchInfo.name"></el-input>
        </el-form-item>
        <!-- <el-select v-model="namespace" placeholder="请选择namespace" @change="namespaceChange">
            <el-option v-for="item in namespacesData" :key="item.metadata.name" :label="item.metadata.name" :value="item.metadata.name"/>
        </el-select> -->
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增PersistentVolume</el-button>
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
    <el-table-column label="Capacity" prop="spec.capacity.storage" width="120"></el-table-column>
    <el-table-column label="ReclaimPolicy" prop="spec.persistentVolumeReclaimPolicy" width="120"></el-table-column>
    <el-table-column label="Status" prop="status.phase" width="120"></el-table-column>
    <el-table-column label="Claim" prop="spec.claimRef.name" width="120"></el-table-column> 
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updatePersistentVolume(scope.row)" size="small" type="primary" icon="el-icon-edit">修改</el-button>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" :title="dialogName">
      <el-form :model="formData" label-position="right" label-width="150px">
        <el-form-item label="Name:">
            <el-input v-model="formData.metadata.name" clearable placeholder="请输入" style="width:200px;"></el-input>
        </el-form-item>

        <el-form-item label="accessModes:">
        <el-checkbox-group v-model="formData.spec.accessModes">
          <el-checkbox label="ReadOnlyMany"></el-checkbox>
          <el-checkbox label="ReadWriteMany"></el-checkbox>
          <el-checkbox label="ReadWriteOnce"></el-checkbox>
        </el-checkbox-group>
        </el-form-item>

        <el-form-item label="回收策略:">
            <el-select v-model="formData.spec.persistentVolumeReclaimPolicy" placeholder="请选择回收策略">
            <el-option v-for="item in reclaimPolicyList" :key="item" :label="item" :value="item"/>
        </el-select>
        </el-form-item>

        <el-form-item label="StorageClass:">
            <el-input v-model="formData.spec.storageClassName" clearable placeholder="请输入" style="width:200px;"></el-input>
        </el-form-item>

         <el-form-item label="存储类型:">
            <el-select v-model="storageType" placeholder="请选择存储类型" @change="storageTypeChange">
            <el-option v-for="item in storageTypeList" :key="item" :label="item" :value="item"/>
        </el-select>
        </el-form-item>
        <el-form-item label="server:" v-if ="nfsVisible">
        <el-input v-model="nfs.server" clearable placeholder="请输入" style="width:200px;" ></el-input>
        </el-form-item>
        <el-form-item label="path:" v-if="nfsVisible">
          <el-input v-model="nfs.path" clearable placeholder="请输入" style="width:200px;"></el-input>
        </el-form-item>
        <el-form-item label="path:" v-if="localVisible">
          <el-input v-model="local.path" clearable placeholder="请输入" style="width:200px;"></el-input>
        </el-form-item>
        <el-form-item label="monitors:" v-if="cephfsVisible">
          <el-input v-model="monitors" type="textarea" rows="2" style="width:200px;" clearable placeholder="地址分行或者以 ; 间隔"></el-input>
        </el-form-item>

         <el-form-item label="path:" v-if="cephfsVisible">
          <el-input v-model="cephfs.path" clearable placeholder="请输入" style="width:200px;"></el-input>
        </el-form-item>
         <el-form-item label="user:" v-if="cephfsVisible">
          <el-input v-model="cephfs.user" clearable placeholder="请输入" style="width:200px;"></el-input>
        </el-form-item>
        <el-form-item label="secretRef:" v-if="cephfsVisible">
          <el-input v-model="cephfs.secretRef.name" clearable placeholder="请输入" style="width:200px;"></el-input>
        </el-form-item>

        <el-form-item label="存储空间大小:">
          <el-input v-model="storage" clearable placeholder="请输入" style="width:100px;"></el-input>
          <!-- <el-select v-model="storageUnit" style="width:80px;" clearable placeholder="单位">
            <el-option v-for="item in storageUnitList" :key="item" :label="item" :value="item"/>
          </el-select> -->

        </el-form-item>
      </el-form>
      
      <div>
        <el-button
          size="small"
          type="primary"
          icon="el-icon-edit"
          @click="addParameter(formData)"
        >新增配置</el-button>
        
        <el-table :data="formData.parameters" stripe style="width: 100%">
          <el-table-column  prop="key" label="key" width="180">
          
            <template slot-scope="scope">
              <div>
                <el-input v-model="scope.row.key" :placeholder="scope.key"></el-input>
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="value" label="值" width="800">
            <template slot-scope="scope">
              <div>
                <el-input v-model.trim="scope.row.value" type="textarea" rows="1"></el-input>
                <!-- <el-input v-model="scope.row.value"></el-input> -->
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
                  @click="deleteParameter(formData.parameters,scope.$index)"
                >删除</el-button>
              </div>
            </template>
          </el-table-column>
        </el-table>
      </div>

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
    createPersistentVolume,
    deletePersistentVolume,
    deletePersistentVolumeByIds,
    updatePersistentVolume,
    findPersistentVolume,
    getPersistentVolumeList
} from "@/api/persistentVolume";  //  此处请自行替换地址


import Menus from "@/view/superAdmin/authority/components/menus";
import Apis from "@/view/superAdmin/authority/components/apis";
import Datas from "@/view/superAdmin/authority/components/datas";
import k8sDrawer from '@/components/k8sDrawer/index.vue';

import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";

import yaml from 'json2yaml';

export default {
  name: "PersistentVolume",
  mixins: [infoList],
  data() {
    return {
      listApi: getPersistentVolumeList,
      namespace: "",
      dialogName: "",
      dialogFormVisible: false,
      drawer: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      accessModes: [],
      reclaimPolicyList:["Retain","Recycle","Delete"],
      formData: {
            apiVersion: "v1",
            kind: "PersistentVolume",
            metadata:{},
            spec:{},
            status:{},
      },
      storage:"",
      storageUnit:"",
      storageUnitList:["Ti","Gi","Mi"],
      storageTypeList: ["NFS","CephFS","Local"],
      storageType:"",
      cephfsVisible: false,
      cephfs:{
        monotors:[],
        user:"",
        secretRef:{
          name:"",
        },
        readOnly:"false",
      },
      monitors:"",
      nfsVisible: false,
      nfs:{
        path:"",
        server:"",
      },
      localVisible: false,
      local:{
        path:"",
      },
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
           this.deletePersistentVolume(row);
        });
      },
      addParameter(formData) {
      if (!formData.parameters) {
        this.$set(formData, "parameters", []);
      }
      formData.parameters.push({
        key: "",
        value: ""
      });
      },
      deleteParameter(parameters, index) {
        parameters.splice(index, 1);
      },
      showConfig(row){
        let data = JSON.parse(JSON.stringify(this.formData));
        console.log(row)
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
        data.data = {}
        for (var i in this.formData.parameters ){
          data.data[this.formData.parameters[i].key] =this.formData.parameters[i].value
        }
        delete data["parameters"]
        this.formDataJson = JSON.stringify(data,null, 4);
        this.formDataYaml = yaml.stringify(data);
        console.log(this.accessModes);
        this.drawer = true;
      },
      storageTypeChange(event){
        if (event == "CephFS") {
          this.cephfsVisible = true;
          this.nfsVisible = false;
          this.localVisible = false;
        }else if (event == "NFS"){
          this.cephfsVisible = false;
          this.nfsVisible = true;
          this.localVisible = false;
        }else if (event == "Local"){
          this.cephfsVisible = false;
          this.nfsVisible = false;
          this.localVisible = true;
        }
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
        const res = await deletePersistentVolumeByIds({ ids })
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
    async updatePersistentVolume(row) {
      const res = await findPersistentVolume({ namespace: row.metadata.namespace,name: row.metadata.name });
      
      if (res.code != 0){
          this.$message({
            type: 'error',
            message: '请求失败'
          });
          return;
      }
      console.log(res)
      this.type = "update";
      delete res.data.reData.metadata["managedFields"]
      this.formData.metadata = res.data.reData.metadata;
      this.formData.spec = res.data.reData.spec;
      this.formData.state =res.data.reData.state;
      this.storage = this.formData.spec.capacity.storage
      if (this.formData.spec.hasOwnProperty("nfs")){
        this.nfs=this.formData.spec.nfs
        this.storageType="nfs"
        this.storageTypeChange("nfs")
      }else if (this.formData.spec.hasOwnProperty("cephfs")){
        this.cephfs=this.formData.spec.cephfs
        for (let i of this.formData.spec.cephfs.monitors){
        this.monitors =   this.monitors + i+"/n"
        }
        this.storageType="cephfs"
        this.storageTypeChange("cephfs")
      }
      else if (this.formData.spec.hasOwnProperty("local")){
        this.local=this.formData.spec.local
        this.storageType="local"
        this.storageTypeChange("local")
      }

      // let storage =  res.data.reData.spec.capacity.storage;
      // this.storageUnit = storage.substring(storage.length - 2)
      // this.formData.spec.storage = storage.substring(0,res.data.reData.spec.storage.length - 2)
      // console.log(storage);
      this.dialogName = "修改"
      this.dialogFormVisible = true;
    },
    resetFormData(){
      this.formData = {
            apiVersion: "v1",
            kind: "PersistentVolume",
            metadata:{},
            spec:{},
            status:{},
      },
      this.cephfs = {
        monotors:[],
        user:"",
        secretRef:{
          name:"",
        },
        readOnly:"false",
      },
      this.nfs={
        path:"",
        server:"",
      },
      this.local={
        path:"",
      },
      this.monitors = ""
    },
    closeDialog() {
      this.resetFormData();
      this.dialogFormVisible = false;
    },
    namespaceChange(){
      this.searchInfo = {
      "namespace":this.namespace,
    }
      this.getTableData();
    },
    async deletePersistentVolume(row) {
      const res = await deletePersistentVolume({namespace: row.metadata.namespace,name: row.metadata.name });
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
      for (var i in data.parameters){
        data.data[this.formData.parameters[i].key] =this.formData.parameters[i].value
      }
      delete data["parameters"]
      this.reqData.name = this.formData.metadata.name;
      this.reqData.namespace = this.formData.metadata.namespace;
      this.reqData.configData = JSON.stringify(data);
      switch (this.type) {
        case "create":
          res = await createPersistentVolume(this.reqData);
          break;
        case "update":
          res = await updatePersistentVolume(this.reqData);
          break;
        default:
          res = await createPersistentVolume(this.reqData);
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



