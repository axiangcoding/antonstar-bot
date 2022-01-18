<template>
  <n-card>
    <template #header>
      <n-h3 type="info" prefix="bar" align-text>
        <n-text type="info">
          快照记录
        </n-text>
      </n-h3>
    </template>
    <template #header-extra>
      <n-button type="warning" size="tiny" secondary @click="refresh">
        <template #icon>
          <SunRegular/>
        </template>
        数据过时了？点我刷新
      </n-button>
    </template>
    <n-space vertical>
      <n-button v-if="data"
                size="tiny" type="info" :dashed="item.query_id!==active" :loading="btnLoading"
                :disabled="item.status==='running'"
                v-for="item in data.gaijin" :key="item.query_id" @click="searchQuery(item.query_id)">
        <div v-if="item.status!=='running'">
          查看于 {{ parseLocalTime(item.updated_at) }} 获得的快照
        </div>
        <div v-else>
          创建于 {{ parseLocalTime(item.created_at) }} 的快照正在查询中...
        </div>
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
  btnLoading.value = true
  emit("searchQuery", queryId)
  btnLoading.value = false
}

const refresh = () => {
  emit("refresh")
}

const emit = defineEmits(["searchQuery", "refresh"])

</script>

<style scoped>

</style>
