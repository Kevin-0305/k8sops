<template>
  <div>
    <el-drawer  :visible.sync="drawer" :with-header="false" size="40%" title="配置文件" :before-close="closeDrawer">
      <el-tabs class="role-box" type="border-card">
        <el-tab-pane label="JSON">
           <!-- <JsonView :json="this.formData"></JsonView> -->
           <el-input v-model.trim="formDataJson" type="textarea" rows="35"></el-input>
        </el-tab-pane>
        <el-tab-pane label="YAML">
          <el-input v-model="formDataYaml" type="textarea" rows="35"></el-input>
          <!-- <textarea >{{this.formDataYaml}}</textarea> -->
        </el-tab-pane>
      </el-tabs>
    </el-drawer>
  </div>
</template>


<script>
  export default {
    name: 'k8sDrawer',
    props: {
          formDataJson: {
            type: String,
            default: ''
          },
          formDataYaml: {
            type: String,
            default: ''
          },
          drawer: {
            type: Boolean,
            default: false
          },
    },
    data (){
      return {
        Json:"",
        Yaml:"",
      };
    },
    methods: {
    closeDrawer(){
      this.$emit('update:drawer',false)
    }
    }
  }

</script>

<style>
.json-box {
  overflow-y: auto;
}
</style>

