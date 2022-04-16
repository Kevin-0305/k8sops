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
          <el-button @click="openDialog" type="primary">新增PersistentVolumeClaim</el-button>
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
          <el-button class="table-button" @click="updatePersistentVolumeClaim(scope.row)" size="small" type="primary" icon="el-icon-edit">修改</el-button>
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

        <el-form-item label="Namespace:">
          <el-select v-model="formData.metadata.namespace" placeholder="请选择namespace" @change="namespaceChange" :disabled="namespaceButton">
            <el-option v-for="item in namespacesData" :key="item.metadata.name" :label="item.metadata.name" :value="item.metadata.name"/>
          </el-select>
        </el-form-item>

        <el-form-item label="pv:" v-if="pvVisible">
          <el-select v-model="formData.spec.volumeMode" placeholder="请选择" @change="pvChange">
            <el-option v-for="item in pvList" :key="item" :label="item" :value="item"/>
          </el-select>
        </el-form-item>

        <el-form-item label="accessModes:">
        <el-checkbox-group v-model="formData.spec.accessModes">
          <el-checkbox label="ReadOnlyMany"></el-checkbox>
          <el-checkbox label="ReadWriteMany"></el-checkbox>
          <el-checkbox label="ReadWriteOnce"></el-checkbox>
        </el-checkbox-group>
        </el-form-item>

        <dictTable :dataList.sync="labels" comName="Labels"></dictTable>

        <el-form-item label="StorageClass:">
            <el-input v-model="formData.spec.storageClassName" clearable placeholder="请输入" style="width:200px;"></el-input>
        </el-form-item>

        <el-form-item label="存储空间大小:">
          <el-input v-model="storage" clearable placeholder="请输入" style="width:100px;"></el-input>
        </el-form-item>
        <el-form-item label="volumeMode:">
        <el-select v-model="formData.spec.volumeMode" placeholder="请选择">
            <el-option v-for="item in volumeModeList" :key="item" :label="item" :value="item"/>
        </el-select>
        </el-form-item>
      </el-form>
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
    createPersistentVolumeClaim,
    deletePersistentVolumeClaim,
    deletePersistentVolumeClaimByIds,
    updatePersistentVolumeClaim,
    findPersistentVolumeClaim,
    getPersistentVolumeClaimList
} from "@/api/persistentVolumeClaim";  //  此处请自行替换地址

import { getPersistentVolumeList } from "@/api/persistentVolume"; 

import Menus from "@/view/superAdmin/authority/components/menus";
import Apis from "@/view/superAdmin/authority/components/apis";
import Datas from "@/view/superAdmin/authority/components/datas";
import k8sDrawer from '@/components/k8sDrawer/index.vue';
import dictTable from '@/components/dictTable/index.vue';

import { formatTimeToStr } from "@/utils/date";
import {listToDict} from "@/utils/dictToList";
import {dictToList} from "@/utils/dictToList";

import infoList from "@/mixins/infoList";

import yaml from 'json2yaml';

export default {
  name: "PersistentVolumeClaim",
  mixins: [infoList],
  data() {
    return {
      listApi: getPersistentVolumeClaimList,
      namespace: "",
      namespaceButton:true,
      dialogName: "",
      dialogFormVisible: false,
      drawer: false,
      type: "",
      deleteVisible: false,
      accessModes: [],
      labels:[],
      pvList:[],
      pvVisible:false,
      formData: {
            apiVersion: "v1",
            kind: "PersistentVolumeClaim",
            metadata:{},
            spec:{},
            status:{},
      },
      storage:"",
      reqData: {
        name:"",
        namespace:"",
        data: {},
      },
      volumeModeList:["Filesystem","Block"],
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
           this.deletePersistentVolumeClaim(row);
        });
      },
      showConfig(row){
        let data = JSON.parse(JSON.stringify(this.formData));
        data.metadata = row.metadata;
        data.spec = row.spec;
        data.status = row.status;
        delete data.metadata["managedFields"]
        if (this.labels.length> 0){
          data.metadata.labels = listToDict(this.labels);
        }
        
        this.formDataJson = JSON.stringify(data,null, 4);
        this.formDataYaml = yaml.stringify(data);
        this.drawer = true;
      },
      openDrawer(){
        let data = JSON.parse(JSON.stringify(this.formData))
        data.data = {}
        if (this.labels.length> 0){
          data.metadata.labels = listToDict(this.labels);
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
        const res = await deletePersistentVolumeClaimByIds({ ids })
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
    async updatePersistentVolumeClaim(row) {
      const res = await findPersistentVolumeClaim({ namespace: row.metadata.namespace,name: row.metadata.name });
      if (res.code != 0){
          this.$message({
            type: 'error',
            message: '请求失败'
          });
          return;
      }
      this.type = "update";
      delete res.data.reData.metadata["managedFields"]
      this.formData.metadata = res.data.reData.metadata;
      this.formData.spec = res.data.reData.spec;
      this.formData.state =res.data.reData.state;
      this.storage = this.formData.spec.resources.requests.storage
      this.labels = dictToList(this.formData.selector.matchLabels);
      this.dialogName = "修改"
      this.namespaceButton = true;
      this.pvVisible=false;
      this.dialogFormVisible = true;
    },
    resetFormData(){
      this.formData = {
            apiVersion: "v1",
            kind: "PersistentVolumeClaim",
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
      this.pvList = []
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
    async deletePersistentVolumeClaim(row) {
      const res = await deletePersistentVolumeClaim({namespace: row.metadata.namespace,name: row.metadata.name });
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
      this.reqData.name = this.formData.metadata.name;
      this.reqData.namespace = this.formData.metadata.namespace;
      this.reqData.configData = JSON.stringify(data);
      switch (this.type) {
        case "create":
          res = await createPersistentVolumeClaim(this.reqData);
          break;
        case "update":
          res = await updatePersistentVolumeClaim(this.reqData);
          break;
        default:
          res = await createPersistentVolumeClaim(this.reqData);
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
    async openDialog() {
        this.type = "create";
        this.dialogName = "新增"
        this.formData.namespace = this.namespace;
        const pvRes = await getPersistentVolumeList({page:1,pageSize:999});
        for (let i of pvRes.data.list){
          this.pvList.push(i.metadata.name)
        }
        this.pvVisible=true;
        this.namespaceButton = true;
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



