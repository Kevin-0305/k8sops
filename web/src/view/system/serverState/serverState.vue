<template>
  <div>
    <el-row :gutter="15" class="system_state">
      <el-col 
      :span="12" 
      v-for="item in serverStats"
      :key = item.ip
      >
        <el-card
          class="card_item"
          :body-style="{ height: '250px', 'overflow-y': 'scroll' }"
        >
          <div slot="header">{{item.hostName}}</div>

          <div>
            <el-row :gutter="10">
              <el-col :span="12">服务器地址:</el-col>
              <el-col :span="12" v-text="item.ip"> </el-col>
            </el-row>

            <el-row :key="item.ip" :gutter="10">
            <el-col :span="10">Load</el-col>

            <el-col :span="12"
                ><el-progress
                type="line"
                :percentage="item.CPU*100|rounding"
                :color="colors"
                ></el-progress
            ></el-col>
            </el-row>

            <el-row :key="item.ip" :gutter="10">
            <el-col :span="6">Ram</el-col>
            <el-col :span="4">已使用 {{item.ramUse/1024/1024/1024|rounding}}G</el-col>
            <el-col :span="12"
                ><el-progress
                type="line"
                :percentage="item.ram*100|rounding"
                :color="colors"
                ></el-progress
            ></el-col>
            <el-col :span="2">总 {{item.ramUse|integer}}G</el-col>
            </el-row>

            <el-row :key="item.ip" :gutter="10">
            <el-col :span="10">Flow</el-col>
            <el-col :span="12"
                >
                累计收发流量 {{item.flow|integer}}G
                <!-- <el-progress
                type="line"
                :percentage="item.flow/1024/1024"
                :color="colors"
                ></el-progress
            > -->
            </el-col>
            </el-row>

            <el-row :key="item.ip" :gutter="10">
            <el-col :span="6">Disk</el-col>
            <el-col :span="4">已使用 {{item.DiskUse|integer}}G</el-col>
            <el-col :span="12"
                ><el-progress
                type="line"
                :percentage="item.disk*100|rounding"
                :color="colors"
                ></el-progress
            ></el-col>
            <el-col :span="2">总 {{item.DiskTotal|integer}}G</el-col>
            </el-row>

          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { getServerStatList } from "@/api/monitorStat.js";
export default {
  name: "State",
  data() {
    return {
      timer:null,
      state: {},
      serverStats:[],
      colors: [
        { color: "#5cb87a", percentage: 20 },
        { color: "#e6a23c", percentage: 40 },
        { color: "#f56c6c", percentage: 80 },
      ],
    };
  },
  created() { 
    this.reload();
    this.timer = setInterval(() => {
      this.reload();
    }, 1000*10);
  },
  beforeDestroy(){
    clearInterval(this.timer)
    this.timer = null
  },
  methods: {
    async reload() {
      const { data } = await getServerStatList();
      this.serverStats = data.list;
    },
  },
  filters: {
  rounding (value) {
    return value.toFixed(2)
  },
  integer (value) {
    return Math.ceil(value/1024/1024/1024)
  }
}
};
</script>

<style>
.system_state {
  padding: 10px;
}

.card_item {
  height: 280px;
}
</style>
