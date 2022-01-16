<template>
  <n-grid item-responsive>
    <!-- 左侧栅栏 -->
    <n-grid-item span="0 768:1 1200:2 1920:4">
      <div class="logo" @click="$router.push({name:'home'})">
        <n-image
            height="46"
            :src="LOGO_URL"
            preview-disabled
        />
      </div>
    </n-grid-item>
    <n-grid-item span="24 768:22 1200:20 1920:16" class="header-content">
      <n-menu inverted v-model:value="activeKey" mode="horizontal" :options="menuOptions"/>
    </n-grid-item>
    <!-- 右侧栅栏 -->
    <n-grid-item span="0 768:1 1200:2 1920:4">
      <n-avatar
          size="large"
          src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"
      />
    </n-grid-item>
  </n-grid>
</template>

<script setup lang="ts">
import logo from '@/assets/logo/logo.png'
import {computed, h, ref} from 'vue'
import {NIcon, NAvatar, NMenu} from 'naive-ui'
import {
  Award,
  PersonBooth as PersonIcon,
  WineGlass as WineIcon,
  Home as HomeIcon
} from '@vicons/fa'
import {RouterLink, useRoute} from "vue-router";

const LOGO_URL = ref(logo)

function renderIcon(icon: any) {
  return () => h(NIcon, null, {default: () => h(icon)})
}

const route = useRoute()
const activeKey = computed(() => {
  return route.name
})

const menuOptions = [
  {
    key: 'record',
    icon: renderIcon(Award),
    label: () =>
        h(
            RouterLink,
            {
              to: {
                name: 'record',
                params: {
                  nick: 'zh-CN'
                }
              }
            },
            {default: () => '战绩查询'}
        ),
  },
  {
    key: 'about',
    icon: renderIcon(Award),
    label: () =>
        h(
            RouterLink,
            {
              to: {
                name: 'record',
                params: {
                  nick: 'zh-CN'
                }
              }
            },
            {default: () => '战绩查询'}
        ),
  },

]
</script>

<style lang="scss" scoped>
.logo:hover {
  cursor: pointer;
}


</style>
