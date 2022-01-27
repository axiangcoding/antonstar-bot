<template>
  <n-card :bordered="false">
    <template #header>
      <n-h1>硬核狠人（还在开发，放出来图一乐）</n-h1>
    </template>
    <n-grid :x-gap="12" :y-gap="8" cols="1 768:3 1200:3 1920:3">
      <n-gi>
        <CommonRankCard title="封神榜" description="凡人之躯，比肩神明"></CommonRankCard>
      </n-gi>
      <n-gi>
        <CommonRankCard title="狠人榜" description="我来，我见，我征服"></CommonRankCard>
      </n-gi>
      <n-gi>
        <CommonRankCard title="恶人榜" description="哥谭噩梦，哈哈哈哈哈哈哈"></CommonRankCard>
      </n-gi>
    </n-grid>
  </n-card>
</template>

<script lang="ts" setup>
import CommonRankCard from "./components/CommonRankCard.vue";
</script>


<style scoped>

</style>
