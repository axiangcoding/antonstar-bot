<script setup lang="ts">

const themeOverrides = {
  common: {
    fontWeightStrong: '600'
  }
}
</script>

<template>
  <n-config-provider :theme-overrides="themeOverrides">
    <router-view/>
  </n-config-provider>
</template>

<style scoped>
</style>
