<template>
  <n-layout-footer>
    <n-space vertical>
      <n-text strong>安东星 v0.2.0</n-text>
      <div>
        <n-button quaternary @click="router.push({name:'about'})">关于我们</n-button>
        <n-divider vertical/>
        <n-button quaternary>问题反馈</n-button>
      </div>
      <n-text italic>安东星开发小组 All Rights Reserved.</n-text>
    </n-space>

  </n-layout-footer>
</template>

<script lang="ts" setup>
import {useRouter} from "vue-router";

const router = useRouter()
</script>


<style lang="scss" scoped>
.n-layout-footer {
  height: var(--footer-height);
  overflow: hidden;
}
</style>
