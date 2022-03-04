<script setup lang="ts">
import {darkTheme} from 'naive-ui'
import themes from '@/themes/index'
import {useStore} from 'vuex'
import {ref, computed, onMounted} from 'vue'
import Loading from './components/Loading.vue';

const store = useStore()
const loading = computed(() => store.state.loading)
const themeOverrides = ref(themes[store.state.themes])

import {v4 as uuid} from "uuid"
import {getSystemInfo} from "@/services/system";


onMounted(() => {
  getSystemInfo().then(res => {
    store.commit('setSystemInfo', res.data)
  })
  if (store.state.clientId == '') {
    store.commit('setClientId', uuid())
  }
  if (store.state.login) {
    // TODO: 校验token是否过期，如果过期就退出登录状态
  }
})
</script>

<template>
  <!-- 调整主题变量 -->
  <n-config-provider :theme-overrides="themeOverrides" class="h100">
    <n-message-provider>
      <n-dialog-provider>
        <n-el tag="div" class="h100">
          <router-view/>
        </n-el>
      </n-dialog-provider>
    </n-message-provider>
    <n-back-top :right="30"/>
    <Loading :loading="loading"/>
  </n-config-provider>
</template>

<style lang="scss">
html,
body {
  margin: 0;
  overflow: hidden;
}

::-webkit-scrollbar {
  width: 8px;
  background: transparent;
  padding: 0;
}

::-webkit-scrollbar-thumb {
  &:hover {
    background: #bbbbbbdd;
  }

  background: #bbbbbbaa;
  border-radius: 4px;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  font-size: 16px;
  height: 100vh;
  overflow-x: hidden;
}
</style>
