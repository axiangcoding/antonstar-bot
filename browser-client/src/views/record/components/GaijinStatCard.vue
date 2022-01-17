<template>
  <n-card
      :content-style="{padding:0}"
  >
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
      <n-descriptions-item label="胜利场次">
        {{ data['胜利场次'] }}
      </n-descriptions-item>
      <n-descriptions-item label="任务总数">
        {{ data['任务总数'] }}
      </n-descriptions-item>
      <n-descriptions-item label="作战胜率">
        {{ data['作战胜率'] }}
      </n-descriptions-item>
      <n-descriptions-item label="阵亡数">
        {{ data['阵亡数'] }}
      </n-descriptions-item>
      <n-descriptions-item label="银狮获得数">
        {{ data['银狮获得数'] }}
      </n-descriptions-item>
      <n-descriptions-item label="游戏时间">
        {{ data['游戏时间'] }}
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
