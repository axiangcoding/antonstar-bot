<template>
  <n-layout embedded>
    <n-layout-header>
      <Header/>
    </n-layout-header>
    <n-layout-content content-style="min-height:1800px">
      <router-view/>
    </n-layout-content>
    <n-layout-footer>成府路</n-layout-footer>
  </n-layout>
</template>

<script lang="ts">
import Header from '@/components/Header.vue'

export default {
  components: {Header}
}
</script>

<style scoped>
.n-layout-header,
.n-layout-footer {
  background: rgba(128, 128, 128, 0.2);
  padding: 24px;
}

.n-layout-sider {
  background: rgba(128, 128, 128, 0.3);
}

.n-layout-content {
  background: rgba(128, 128, 128, 0.4);
}
</style>