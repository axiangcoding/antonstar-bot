<template>
  <n-card size="small">
    <template #header>
      <n-h3 prefix="bar" align-text>
        <n-text type="primary">
          Thunder Skill效率
        </n-text>
        <n-tooltip trigger="hover" placement="top">
          <template #trigger>
            <n-icon size="16">
              <QuestionCircleRegular/>
            </n-icon>
          </template>
          效率值仅供参考，由 Thunder Skill计算得到
        </n-tooltip>

      </n-h3>
    </template>
    <template #header-extra>
      <n-button dashed type="info" size="small" @click="showWhy">为什么我没有？</n-button>
    </template>
    <n-space vertical v-if="data && data.found">
      <n-tag type="success"> 该记录由Thunder skill在 {{ parseLocalTime(data.created_at) }} 生成</n-tag>
      <n-tag type="info"> 数据统计的时间为 {{ parseLocalTime(data.last_stat) }}</n-tag>
      <n-grid :x-gap="12" :y-gap="8" cols="1 768:3 1200:2 1920:3">
        <n-gi>
          <n-space vertical align="center">
            <n-progress type="circle" :color="generateColor(data.a.kpd)"
                        :rail-color="changeColor(generateColor(data.a.kpd), { alpha: 0.2 })"
                        :indicator-text-color="generateColor(data.a.kpd)" :percentage="data.a.kpd">
            </n-progress>
            <n-gradient-text type="primary" size="20">街机娱乐</n-gradient-text>
          </n-space>
        </n-gi>
        <n-gi>
          <n-space vertical align="center">
            <n-progress type="circle" :color="generateColor(data.r.kpd)"
                        :rail-color="changeColor(generateColor(data.r.kpd), { alpha: 0.2 })"
                        :indicator-text-color="generateColor(data.r.kpd)" :percentage="data.r.kpd">
            </n-progress>
            <n-gradient-text type="primary" size="20">历史性能</n-gradient-text>
          </n-space>
        </n-gi>
        <n-gi>
          <n-space vertical align="center">
            <n-progress type="circle" :color="generateColor(data.s.kpd)"
                        :rail-color="changeColor(generateColor(data.s.kpd), { alpha: 0.2 })"
                        :indicator-text-color="generateColor(data.s.kpd)" :percentage="data.s.kpd">
            </n-progress>
            <n-gradient-text size="20">全真模拟</n-gradient-text>
          </n-space>
        </n-gi>
      </n-grid>
    </n-space>
    <n-space vertical v-else>
      <n-result
          status="info"
          title="友情提示"
          description="对不起，好像没有你的数据哦"
      >
      </n-result>
    </n-space>
  </n-card>
</template>

<script lang="ts" setup>
import {useDialog, useThemeVars} from 'naive-ui'
import {ref} from "vue";
import {changeColor} from 'seemly'
import {parseLocalTime} from '@/util/time'
import {QuestionCircleRegular} from "@vicons/fa";

const customColors = [
  {color: '#ffffff', percentage: 0, text: '在？'},
  {color: '#5cb87a', percentage: 10, text: '战雷之拉'},
  {color: '#1989fa', percentage: 20, text: '战雷之屑'},
  {color: '#6f7ad3', percentage: 50, text: '战雷之虫'},
  {color: '#e4a13d', percentage: 70, text: '战雷之人'},
  {color: '#f56c6c', percentage: 80, text: '战雷之王'},
  {color: '#8B0000', percentage: 90, text: '战雷皇帝'},
  {color: '#000000', percentage: 100, text: '神中神'},
]

const themeVars = ref(useThemeVars())
defineProps({
  data: Object
})

const generateColor = (percentage: number) => {
  for (let customColor of customColors) {
    if (percentage < customColor.percentage) {
      return customColor.color
    }
  }
  return customColors[customColors.length - 1].color
}

const generateText = (percentage: number) => {
  for (let customColor of customColors) {
    if (percentage < customColor.percentage) {
      return customColor.text
    }
  }
  return customColors[customColors.length - 1].text
}

const dialog = useDialog();
const showWhy = () => {
  dialog.info({
    title: '为什么没有Thunder Skill（TS）的数据',
    content: 'TS不会主动计算一个玩家的数据，如果需要请到TS上手动生成'
  })
}
</script>

<style scoped>

</style>
