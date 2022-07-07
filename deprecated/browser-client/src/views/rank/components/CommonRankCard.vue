<template>
  <n-card size="small" class="card-rank">
    <template #header>
      <n-h3 prefix="bar" align-text>
        <n-text type="primary">
          {{ title }}
        </n-text>
      </n-h3>
    </template>
    <template #header-extra>
      <n-text underline italic>
        {{description}}</n-text>
    </template>
    <n-list>
      <n-list-item v-for="i in 10">
        <n-space>
          <n-button strong size="small" quaternary type="primary">[StGGA] WT_GodFather</n-button>
          <n-tag type="error" size="small">玩MI28NM</n-tag>
          <n-tag type="error" size="small">玩KA52</n-tag>
        </n-space>
      </n-list-item>

    </n-list>
  </n-card>
</template>

<script lang="ts" setup>
defineProps({
  title: String,
  description: String
})
</script>

<style scoped>
.card-rank {
  text-align: left;
  min-height: 400px;
}
</style>
