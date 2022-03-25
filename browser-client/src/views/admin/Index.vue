<template>
  <n-card :bordered="false" style="min-width: 1200px">
    <n-layout has-sider>
      <n-layout-sider content-style="padding-top: 5px;"
                      collapse-mode="width"
                      :collapsed-width="64"
                      :width="200"
      >
        <n-menu :options="menuOptions" @update:value="handleUpdateValue"/>
      </n-layout-sider>
      <n-layout-content content-style="padding: 24px;">
        <router-view/>
      </n-layout-content>
    </n-layout>
  </n-card>
</template>

<script lang="ts" setup>

import {renderIcon} from "@/util/naive";
import {h} from "vue";
import {RouterLink} from "vue-router";
import {MenuOption} from "naive-ui";
import {Users, ChartArea, Globe, MicrophoneAlt, ExternalLinkAlt} from "@vicons/fa";

const menuOptions: MenuOption[] = [
  {
    label: () =>
        h(RouterLink, {to: {name: 'admin',}}, {default: () => '总览'}),
    key: 'admin',
    icon: renderIcon(ChartArea)
  },
  {
    key: 'divider-1',
    type: 'divider',
    props: {
      style: {
        marginLeft: '32px'
      }
    }
  },
  {
    label: () =>
        h(RouterLink, {to: {name: 'adminNotice',}}, {default: () => '公告管理'}),
    key: 'adminNotice',
    icon: renderIcon(MicrophoneAlt)
  },
  {
    label: () =>
        h(RouterLink, {to: {name: 'adminUser',}}, {default: () => '用户管理'}),
    key: 'adminUser',
    icon: renderIcon(Users)
  },
  {
    label: () =>
        h(RouterLink, {to: {name: 'adminGlobalConfig',}}, {default: () => '全局配置'}),
    key: 'adminGlobalConfig',
    icon: renderIcon(Globe)
  },
]

const handleUpdateValue = () => {

}
</script>

<style scoped>

</style>
