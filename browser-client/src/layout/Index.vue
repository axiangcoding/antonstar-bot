<template>
  <n-layout embedded class="layout">
    <Header></Header>
    <n-layout-content class="content">
      <router-view></router-view>
    </n-layout-content>
    <Footer></Footer>
  </n-layout>
</template>

<script lang="ts" setup>
import Header from '@/components/Header.vue'
import Footer from '@/components/Footer.vue'
import {useMessage} from "naive-ui";
// 在组件外的js中调用message，需要挂载到全局变量中，同时需要在n-message-provider的子级组件中挂载
window.$message = useMessage()
</script>

<style lang="scss" scoped>
.content {
  padding-top: $headerHeight;
  min-height: calc(100vh - $footerHeight);
  max-width: 1200px;
  margin: 0 auto;
}

.layout {
  background-image: linear-gradient(135deg, #FFF5C3 10%, #9452A5 100%);
}
</style>
