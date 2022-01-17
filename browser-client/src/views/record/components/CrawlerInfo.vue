<template>
  <n-card>
    <template #header>
      <n-h3 type="info" prefix="bar" align-text>
        <n-text type="info">
          获取记录
        </n-text>
      </n-h3>
    </template>
    <template #header-extra>
      <n-button type="warning" size="small" secondary>
        <template #icon>
          <SunRegular/>
        </template>
        数据过时了？点我刷新
      </n-button>
    </template>
    <n-space vertical>
      <n-button v-if="data" size="tiny" type="info" :dashed="item.query_id!==active" :loading="btnLoading"
                v-for="item in data.gaijin" :key="item.query_id" @click="searchQuery(item.query_id)">
        查看于 {{ parseLocalTime(item.updated_at) }} 获得的快照 【{{ item.found ? '找到' : '未找到' }} {{ item.status }}】
      </n-button>
    </n-space>
  </n-card>
</template>

<script lang="ts" setup>
import {SunRegular} from "@vicons/fa";
import {parseLocalTime} from "@/util/time";
import {ref} from "vue";

const props = defineProps({
  data: Object,
  active: String
})

const btnLoading = ref(false)
const searchQuery = (queryId: string) => {
  btnLoading.value=true
  emit("searchQuery", queryId)
  btnLoading.value=false
}

const emit = defineEmits(["searchQuery"])

</script>

<style scoped>

</style>
