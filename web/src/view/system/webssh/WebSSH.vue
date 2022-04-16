<template>
  <div class="container">
    <my-terminal :terminal="terminal"></my-terminal>
  </div>
</template>

<script>
import Console from './Console'

export default {
  name: "WebSSh",
  data () {
    return {
      terminal: {
        pid: 1,
        serverId:0,
        name: 'terminal',
        cols: 400,
        rows: 400
      }
    }
  },
  created(){
          // 取到路由带过来的参数
          const  serverId = this.$route.query.id
          this.terminal.serverId = serverId
          this.terminal.name = this.$route.query.hostName
          this.terminal.pid = serverId
  },
  
  components: {
    'my-terminal': Console
  },


}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
