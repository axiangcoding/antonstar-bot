<template>
  <n-card :bordered="false" size="small">
    <template #header>
      <n-h2 prefix="bar" align-text type="primary">数据统计</n-h2>
    </template>
    <n-row>
      <n-col :span="12">
        <n-statistic label="今日页面访问次数">
          <n-number-animation
              ref="numberAnimationInstRef"
              show-separator
              :from="0"
              :to="todayVisit"
              :active="true"
          />
        </n-statistic>
      </n-col>
      <n-col :span="12">
        <n-statistic label="累计页面访问次数">
          <n-number-animation
              ref="numberAnimationInstRef"
              show-separator
              :from="0"
              :to="totalVisit"
              :active="true"
          />
        </n-statistic>
      </n-col>
      <n-col :span="12">
        <n-statistic label="今日玩家查询次数">
          <n-number-animation
              ref="numberAnimationInstRef"
              show-separator
              :from="0"
              :to="todayQuery"
              :active="true"
          />
        </n-statistic>
      </n-col>
      <n-col :span="12">
        <n-statistic label="累计玩家查询次数">
          <n-number-animation
              ref="numberAnimationInstRef"
              show-separator
              :from="0"
              :to="totalQuery"
              :active="true"
          />
        </n-statistic>
      </n-col>
    </n-row>
  </n-card>
</template>

<script lang="ts" setup>
import {onMounted, ref} from "vue";
import http from "@/services/request";
import {today, toISOTime} from "@/util/time";

const todayVisit = ref()
const totalVisit = ref()

const todayQuery = ref()
const totalQuery = ref()

const queryTodayVisitCount = () => {
  http.get("v1/visits/count", {
    params: {
      'timestamp': toISOTime(today())
    }
  }).then(res => {
    todayVisit.value = res.data.count
  })
}

const queryAllVisitCount = () => {
  http.get("v1/visits/count").then(res => {
    totalVisit.value = res.data.count
  })
}

const queryTodayQueryCount = () => {
  http.get("v1/war_thunder/userinfo/query/count", {
    params: {
      'timestamp': toISOTime(today())
    }
  }).then(res => {
    todayQuery.value = res.data.count
  })
}

const queryAllQueryCount = () => {
  http.get("v1/war_thunder/userinfo/query/count").then(res => {
    totalQuery.value = res.data.count
  })
}

onMounted(() => {
  queryTodayVisitCount()
  queryAllVisitCount()
  queryTodayQueryCount()
  queryAllQueryCount()
})
</script>

<style scoped>

</style>
