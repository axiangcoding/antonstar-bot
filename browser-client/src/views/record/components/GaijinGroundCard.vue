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
      <n-descriptions-item :label-style="{'width':'50%'}" label="参战次数(陆战)">
        {{ data['参战次数(陆战)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="参战次数(地面载具)">
        {{ data['参战次数(地面载具)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="参战次数(坦克歼击车)">
        {{ data['参战次数(坦克歼击车)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="参战次数(重型坦克)">
        {{ data['参战次数(重型坦克)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="参战次数(自行防空炮)">
        {{ data['参战次数(自行防空炮)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="游戏时长(陆战)">
        {{ data['游戏时长(陆战)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="游戏时长(地面单位)">
        {{ data['游戏时长(地面单位)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="游戏时长(坦克歼击车)">
        {{ data['游戏时长(坦克歼击车)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="游戏时长(重型坦克)">
        {{ data['游戏时长(重型坦克)'] }}
      </n-descriptions-item>
      <n-descriptions-item label="游戏时长(自行防空炮)">
        {{ data['游戏时长(自行防空炮)'] }}
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

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
