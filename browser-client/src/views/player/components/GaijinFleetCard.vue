<template>
  <n-card size="small"
          :header-style="{margin:0}"
          :content-style="{padding:0}">
    <template #header>
      <n-h3 prefix="bar" align-text>
        <n-text type="primary">
          {{ title === 'ab' ? '街机娱乐' : title === 'rb' ? '历史性能' : '全真模拟' }}
        </n-text>
      </n-h3>
    </template>
    <n-descriptions
        v-if="data"
        :column="1"
        label-placement="left"
        bordered
    >
      <n-descriptions-item :label-style="{'width':'50%'}" n label="参战次数(海军)">
        {{ data['参战次数(海军)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="参战次数(舰船)">
        {{ data['参战次数(舰船)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="参战次数(鱼雷艇)">
        {{ data['参战次数(鱼雷艇)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="参战次数(炮艇)">
        {{ data['参战次数(炮艇)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="参战次数(鱼雷炮艇)">
        {{ data['参战次数(鱼雷炮艇)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="参战次数(猎潜艇)">
        {{ data['参战次数(猎潜艇)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="参战次数(海军驳渡船)">
        {{ data['参战次数(海军驳渡船)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="游戏时长(海战)">
        {{ data['游戏时长(海战)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="游戏时长(船舰)">
        {{ data['游戏时长(船舰)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="游戏时长(鱼雷艇)">
        {{ data['游戏时长(鱼雷艇)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="游戏时长(炮艇)">
        {{ data['游戏时长(炮艇)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="游戏时长(鱼雷炮艇)">
        {{ data['游戏时长(鱼雷炮艇)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="游戏时长(猎潜艇)">
        {{ data['游戏时长(猎潜艇)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="游戏时长(驱逐舰)">
        {{ data['游戏时长(驱逐舰)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="游戏时长(海军驳渡船)">
        {{ data['游戏时长(海军驳渡船)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="击毁目标总计">
        {{ data['击毁目标总计'] }}
      </n-descriptions-item>
      <n-descriptions-item label="空中单位摧毁数">
        {{ data['空中单位摧毁数'] }}
      </n-descriptions-item>
      <n-descriptions-item label="地面单位摧毁数">
        {{ data['地面单位摧毁数'] }}
      </n-descriptions-item>
      <n-descriptions-item label="水面单位摧毁数">
        {{ data['水面单位摧毁数'] }}
      </n-descriptions-item>
    </n-descriptions>
  </n-card>
</template>

<script lang="ts" setup>
defineProps({
  data: Object,
  title: String,
})
</script>

<style lang="scss" scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
