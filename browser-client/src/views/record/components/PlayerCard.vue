<template>
  <n-card size="small" hoverable embedded="">
    <n-thing>
      <template #avatar>
        <n-avatar
            round
            object-fit="scale-down"
        >{{ data.nick.substring(0, 1).toUpperCase() }}
        </n-avatar>
      </template>
      <template #header>
        {{ data.clan }}
        <n-divider vertical/>
        {{ data.nick }}
      </template>
      <template #header-extra>
        <n-tag v-if="data.banned" size="small" type="error">
          已被封禁！
          <template #avatar>
            <n-icon>
              <Ban/>
            </n-icon>
          </template>
        </n-tag>

      </template>
      <template #description>
        <n-space>
          <n-tag type="primary" size="small" round>{{ data.title ? data.title : '-' }}</n-tag>
          <n-tag type="info" size="small" round>{{ data.level }}级</n-tag>
        </n-space>
      </template>
      <n-descriptions :column="1" label-placement="left" bordered>
        <n-descriptions-item :label-style="{'width':'40%'}" label="击杀死亡比">
          暂未提供
        </n-descriptions-item>
        <n-descriptions-item label="街机效率">
          <n-progress
              type="line"
              :percentage="data.ts_ab_rate"
          />
        </n-descriptions-item>

        <n-descriptions-item label="历史效率" style="width: 100%">
          <n-progress
              type="line"
              :percentage="data.ts_rb_rate"
          />
        </n-descriptions-item>
        <n-descriptions-item label="全真效率">
          <n-progress
              type="line"
              :percentage="data.ts_sb_rate"
          />
        </n-descriptions-item>
      </n-descriptions>


      <template #footer>
        <n-tag size="small" type="info">更新时间：{{ parseLocalTime(data.lastUpdate) }}</n-tag>
      </template>
      <template #action>
        <n-space justify="end">
          <n-button type="info" size="small" disabled>战绩对比</n-button>
          <n-button type="primary" size="small" @click="router.push({name:'player', params:{nick:data.nick}})">
            查看详情
          </n-button>
        </n-space>
      </template>
    </n-thing>
  </n-card>

</template>

<script lang="ts" setup>
import {parseLocalTime} from "@/util/time";
import {Ban} from "@vicons/fa";
import {useRouter} from "vue-router";

const router = useRouter();
const props = defineProps({
  data: Object
});


</script>

<style scoped>

</style>