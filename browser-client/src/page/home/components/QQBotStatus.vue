<template>
  <n-card hoverable title="QBot机器人实时状态">
    <n-space v-show="showTags">
      <n-tag :type="status.app_enabled?'success':'error'">
        {{ status.app_enabled ? "程序可用" : "程序不可用" }}
        <template #icon>
          <n-icon>
            <check-circle-filled v-if="status.app_enabled"/>
            <error-filled v-else/>
          </n-icon>
        </template>
      </n-tag>
      <n-tag :type="status.plugins_good==null?'warning':status.plugins_good?'success':'error'">
        {{ status.plugins_good == null ? '插件状态未知' : status.plugins_good ? '插件状态正常' : '插件状态不正常' }}
        <template #icon>
          <n-icon>
            <check-circle-filled v-if="status.app_enabled"/>
            <error-filled v-else/>
          </n-icon>
        </template>
      </n-tag>
      <n-tag :type="status.app_good?'success':'error'">
        {{ status.app_good ? "程序正常" : "程序不正常" }}
        <template #icon>
          <n-icon>
            <check-circle-filled v-if="status.app_good"/>
            <error-filled v-else/>
          </n-icon>
        </template>
      </n-tag>
      <n-tag :type="status.online?'success':'error'">
        {{ status.online ? "账号在线" : "账号不在线" }}
        <template #icon>
          <n-icon>
            <check-circle-filled v-if="status.online"/>
            <error-filled v-else/>
          </n-icon>
        </template>
      </n-tag>
    </n-space>
  </n-card>
</template>

<script setup lang="ts">
import {CheckCircleFilled, ErrorFilled} from '@vicons/material'
import {onBeforeMount, onMounted, ref} from "vue";
import api from "@/api";

const showTags = ref(false)

const defaultStatus = {
  app_enabled: false,
  app_good: false,
  online: false,
  plugins_good: null
}

const status = ref(defaultStatus)

let timer: NodeJS.Timer

onMounted(() => {
  getStatus()

  timer = setInterval(() => {
    getStatus()
  }, 5000);
})

onBeforeMount(() => {
  clearInterval(Number(timer))
})

function getStatus() {
  api.get('/api/v1/cqhttp/status').then((resp) => {
    if (resp.data.code === 0) {
      status.value = resp.data.data
    } else {
      status.value = defaultStatus
    }
  }).catch((e) => {
    status.value = defaultStatus
  })
  showTags.value = true
}
</script>

<style scoped>

</style>
