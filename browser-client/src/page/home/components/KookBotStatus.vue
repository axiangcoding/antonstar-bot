<template>
  <n-card title="Kook机器人实时状态">
    <n-space>
      <n-tag type="info">
        暂未接入
        <template #icon>
          <n-icon>
            <error-filled/>
          </n-icon>
        </template>
      </n-tag>
    </n-space>
  </n-card>
</template>

<script setup lang="ts">
import {ErrorFilled} from '@vicons/material'
</script>

<style scoped>

</style>