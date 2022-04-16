    <template>
      <div>
        <el-form :inline="true">
        <el-link target="_blank" type="primary">{{comName}}</el-link>
        <el-button
          size="small"
          type="primary"
          icon="el-icon-edit"
          @click="addParameter(dataList)"
          style="float:right"
        >新增</el-button>
        </el-form>
        
        <el-table :data="dataList" stripe style="width: 100%">
          <el-table-column  prop="key" label="key" width="180" >
            <template slot-scope="scope">
              <div>
                <el-input v-model="scope.row.key" :placeholder="scope.key"></el-input>
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="value" label="值" width="650" >
            <template slot-scope="scope">
              <div>
                <el-input v-model.trim="scope.row.value" type="textarea" rows="1" @change="inputChange"></el-input>
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
                  @click="deleteParameter(dataList,scope.$index)"
                >删除</el-button>
              </div>
            </template>
          </el-table-column>
        </el-table>
        </div>
</template>

<script>
  export default {
    name: 'dictTable',
    props: {
          dataList: {
            type: Array,
            default: {}
          },
          comName: {
            type: String,
            default: ''
          },
    },
    data (){
      return {
      };
      
    },
    methods: {
      inputChange(){
        this.$emit('update:dataList',this.dataList)
      },
      addParameter(dataList){
        dataList.push({
          key:"",
          value:"",
        })
      },
      deleteParameter(dataList,index){
        dataList.splice(index, 1);
      }
    },
  }
</script>

<style>
</style>