<script setup lang="ts">

import AppFooter from "@/components/AppFooter.vue";

const themeOverrides = {
  common: {
    fontWeightStrong: '600'
  }
}
</script>

<template>

  <n-config-provider  :theme-overrides="themeOverrides">
    <n-layout>
      <n-layout-content>
        <router-view/>
      </n-layout-content>

      <n-layout-footer>
      <app-footer/>
      </n-layout-footer>
    </n-layout>

  </n-config-provider>
</template>

<style scoped>


.n-layout-footer {
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f5f5f5;
  margin: auto auto;
  height: 70px;
}

</style>
