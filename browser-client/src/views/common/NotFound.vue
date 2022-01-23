<template>
  <n-card :bordered="false">
    <n-result status="404" title="404 资源不存在" description="生活总归带点荒谬">
      <template #footer>
        <n-button type="primary" @click="router.back()">返回上一页</n-button>
      </template>
    </n-result>
  </n-card>

</template>

<script lang="ts" setup>
import {useRouter} from "vue-router";

const router = useRouter();


</script>

<style scoped>

</style>
