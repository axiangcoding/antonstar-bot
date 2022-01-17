<script setup lang="ts">
import {darkTheme} from 'naive-ui'
import themes from '@/themes/index'
import {useStore} from 'vuex'
import {ref} from 'vue'

const store = useStore()
const themeOverrides = ref(themes[store.state.themes])

</script>

<template>
  <!-- 调整主题变量 -->
  <n-config-provider :theme-overrides="themeOverrides" class="h100">
    <n-message-provider>
      <n-el tag="div">
        <router-view/>
      </n-el>
    </n-message-provider>
    <n-back-top :right="100"/>
  </n-config-provider>
</template>

<style lang="scss">
html,
body {
  margin: 0;
  overflow: hidden;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  font-size: $fontSize;
  height: 100vh;
  overflow-x: hidden;

}
</style>
