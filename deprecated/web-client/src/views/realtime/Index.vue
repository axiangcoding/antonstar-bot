<template>
  <div>
    <el-card class="card-main">
      <!--<iframe style="width: 100%;height: 87vh" :src="localhostUrl" :onerror="handleIframeError">-->
      <!--</iframe>-->
    </el-card>
  </div>
</template>

<script lang="ts">
export default {
  name: "Index"
}
</script>


<script lang="ts" setup>
import {onMounted, onUnmounted, ref} from "vue";
import service from "../../util/request";

const handleIframeError = () => {
  console.log("error")
}


const localhostUrl = ref('http://localhost:8111')


onMounted(() => {

})

onUnmounted(() => {

})
</script>
<style scoped>
.card-main {
  min-height: 87vh;
}
</style>
