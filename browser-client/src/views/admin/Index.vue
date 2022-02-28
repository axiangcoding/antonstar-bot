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
import {Book, ChartArea, Home, MicrophoneAlt, PersonBooth, WineBottle} from "@vicons/fa";

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
    label: '1973年的弹珠玩具',
    key: 'pinball-1973',
    icon: renderIcon(Book),
    disabled: true,
    children: [
      {
        label: '鼠',
        key: 'rat'
      }
    ]
  },
  {
    label: '舞，舞，舞',
    key: 'dance-dance-dance',
    icon: renderIcon(Book),
    children: [
      {
        type: 'group',
        label: '人物',
        key: 'people',
        children: [
          {
            label: '叙事者',
            key: 'narrator',
            icon: renderIcon(PersonBooth)
          },
          {
            label: '羊男',
            key: 'sheep-man',
            icon: renderIcon(PersonBooth)
          }
        ]
      },
      {
        label: '饮品',
        key: 'beverage',
        icon: renderIcon(WineBottle),
        children: [
          {
            label: '威士忌',
            key: 'whisky'
          }
        ]
      },
      {
        label: '食物',
        key: 'food',
        children: [
          {
            label: '三明治',
            key: 'sandwich'
          }
        ]
      },
      {
        label: '过去增多，未来减少',
        key: 'the-past-increases-the-future-recedes'
      }
    ]
  }
]

const handleUpdateValue = () => {

}
</script>

<style scoped>

</style>
