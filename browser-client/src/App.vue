<script setup lang="ts">
import themes from '@/themes/index'
import {useStore} from 'vuex'
import {ref, computed, onMounted} from 'vue'
import Loading from './components/Loading.vue';

const store = useStore()
const loading = computed(() => store.state.loading)
const themeOverrides = ref(themes[store.state.themes])

import {v4 as uuid} from "uuid"
import {getSystemInfo} from "@/services/system";
import {zhCN, dateZhCN} from "naive-ui";

onMounted(() => {
  console.log("🚀🚀🚀千里之行，始于足下")
  console.log("😎😎😎有兴趣参与安东星的开发吗，私信b站用户【摸鱼的33】获取更多信息，来为社区贡献能力！")
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
  <n-config-provider :theme-overrides="themeOverrides" class="h100" :locale="zhCN" :date-locale="dateZhCN">
    <n-message-provider>
      <n-dialog-provider>
        <n-el tag="div" class="h100">
          <router-view/>
        </n-el>
      </n-dialog-provider>
    </n-message-provider>

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
