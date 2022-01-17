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
      <n-descriptions-item :label-style="{'width':'50%'}" label="参战次数(空战)">
        {{ data['参战次数(空战)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="参战次数(战斗机)">
        {{ data['参战次数(战斗机)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="参战次数(轰炸机)">
        {{ data['参战次数(轰炸机)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="参战次数(攻击机)">
        {{ data['参战次数(攻击机)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="游戏时长(空战)">
        {{ data['游戏时长(空战)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="游戏时长(战斗机)">
        {{ data['游戏时长(战斗机)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="游戏时长(轰炸机)">
        {{ data['游戏时长(轰炸机)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="游戏时长(攻击机)">
        {{ data['游戏时长(攻击机)'] }}
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
