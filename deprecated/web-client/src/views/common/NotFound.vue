<template>
  <el-card class="card-main">
    <el-empty
        :image="notFoundUrl"
        :image-size="1000" description="页面未找到，是不是还在开发中？">
      <el-button  type="info" @click="router.back()">返回上一页</el-button>
    </el-empty>
  </el-card>

</template>

<script lang="ts">
export default {
  name: "not_found"
}
</script>

<script lang="ts" setup>

import {useRouter} from "vue-router";
import url from "../../assets/image/404.jpg"
import {ref} from "vue";


const router = useRouter();
const notFoundUrl = ref(url)

</script>

<style lang="scss" scoped>
.card-main {
  min-height: 85vh;
}
</style>
