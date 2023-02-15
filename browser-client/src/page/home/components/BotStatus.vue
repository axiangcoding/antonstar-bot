<template>
  <n-card title="机器人实时状态">
    <n-descriptions column="1" label-placement="left">
      <n-descriptions-item label="QQ机器人">
        正常
        <n-icon color="18a058">
          <check-circle-filled/>
        </n-icon>
      </n-descriptions-item>
      <n-descriptions-item label="Kook机器人">
        故障
        <n-icon color="d03050">
          <error-filled/>
        </n-icon>
      </n-descriptions-item>
    </n-descriptions>
  </n-card>
</template>

<script setup lang="ts">
import {CheckCircleFilled, ErrorFilled} from '@vicons/material'
import axios from "axios";
import {onMounted} from "vue";

onMounted(() => {
  getStatus()
})

function getStatus() {
  axios.get('/api/v1/system/info').then((resp) => {
    console.log(resp)
  })
}
</script>

<style scoped>

</style>