<template>
  <n-card embedded :bordered="true">
    <n-spin :show="reloading">
      <n-space v-if="displayStatus==='find'" vertical>
        <div>
          <n-button quaternary type="info" size="small" disabled>这是我</n-button>
          <n-divider vertical/>
          <n-button quaternary type="primary" size="small" disabled>给大佬点赞</n-button>
          <n-divider vertical/>
          <n-button quaternary type="error" size="small" disabled>站内举报</n-button>
          <n-divider vertical/>
        </div>
        <n-grid cols="1 768:2 1200:2 1920:2" :x-gap="12" :y-gap="8">
          <n-gi>
            <CommonInfo :data="gaijinData"></CommonInfo>
          </n-gi>
          <n-gi>
            <CrawlerInfo :data="queryList" :active="activeQuery" @searchQuery="getInfo"
                         @refresh="refreshInfoQueries(route.params.nick)"></CrawlerInfo>
          </n-gi>
        </n-grid>
        <n-grid cols="1 768:3 1200:2 1920:3" :x-gap="12" :y-gap="8">
          <n-gi v-for="(index,key) in gaijinData.user_stat" :key="key">
            <GaijinStatCard :data="index" :title="key"></GaijinStatCard>
          </n-gi>
        </n-grid>
        <n-tabs type="line" v-if="gaijinData.user_rate">
          <n-tab-pane name="aviation" tab="空军">
            <n-grid cols="1 768:3 1200:2 1920:3" :x-gap="12" :y-gap="8">
              <n-gi v-for="(index,key) in gaijinData.user_rate.aviation" :key="key">
                <GaijinAviationCard :data="index" :title="key"></GaijinAviationCard>
              </n-gi>
            </n-grid>
          </n-tab-pane>
          <n-tab-pane name="陆军" tab="陆军">
            <n-grid cols="1 768:3 1200:2 1920:3" :x-gap="12" :y-gap="8">
              <n-gi v-for="(index,key) in gaijinData.user_rate.ground_vehicles" :key="key">
                <GaijinGroundCard :data="index" :title="key"></GaijinGroundCard>
              </n-gi>
            </n-grid>
          </n-tab-pane>
          <n-tab-pane name="海军" tab="海军">
            <n-grid cols="1 768:3 1200:2 1920:3" :x-gap="12" :y-gap="8">
              <n-gi v-for="(index,key) in gaijinData.user_rate.fleet" :key="key">
                <GaijinFleetCard :data="index" :title="key"></GaijinFleetCard>
              </n-gi>
              <n-gi>
                <!--<CommonInfo></CommonInfo>-->
              </n-gi>
              <n-gi>
                <!--<CommonInfo></CommonInfo>-->
              </n-gi>
            </n-grid>
          </n-tab-pane>
        </n-tabs>
      </n-space>
      <n-space align="center" vertical v-else-if="displayStatus==='nothing'">
        <n-empty>未在本站中找到该用户记录</n-empty>
        <n-button type="primary" @click="refreshInfoQueries(route.params.nick)">点我向官网发起查询</n-button>
      </n-space>
      <n-space align="center" vertical v-else-if="displayStatus==='running'">
        <n-empty>正在向官网查询中，请稍等</n-empty>
      </n-space>
      <n-space vertical v-else>
        {{ displayStatus }}
        <n-skeleton text :repeat="5" :animated="false"></n-skeleton>
      </n-space>
    </n-spin>
  </n-card>
</template>

<script lang="ts" setup>
import {NCard, NGrid, NGi, NSpace, NTabs, NTabPane, NDivider, useMessage} from "naive-ui";
import CommonInfo from "@/views/record/components/CommonInfo.vue";
import GaijinStatCard from "@/views/record/components/GaijinStatCard.vue";
import CrawlerInfo from "@/views/record/components/CrawlerInfo.vue";
import {ref, watch} from "vue";
import http from "@/services/request";
import {useRoute, useRouter} from "vue-router";
import GaijinAviationCard from "@/views/record/components/GaijinAviationCard.vue";
import GaijinGroundCard from "@/views/record/components/GaijinGroundCard.vue";
import GaijinFleetCard from "@/views/record/components/GaijinFleetCard.vue";

const props = defineProps({
  queryList: Object,
});

const route = useRoute();
const activeQuery = ref()
const gaijinData = ref({})
const thunderskillData = ref({})

const displayStatus = ref('none')
watch(props, (newVal, oldVal) => {

  if (props.queryList !== undefined && props.queryList.gaijin !== undefined) {
    // 找到最新的一条记录
    let found = false
    let queryId = props.queryList.gaijin[0].query_id
    for (let item of props.queryList.gaijin) {
      if (item.found) {
        found = true
        queryId = item.query_id
        break
      }
    }
    if (found) {
      displayStatus.value = 'find'
      getInfo(queryId)
    } else {
      displayStatus.value = 'running'
    }
  } else {
    displayStatus.value = 'nothing'
  }
})

const reloading = ref(false)
const getInfo = async (queryId: string) => {
  reloading.value = true
  activeQuery.value = queryId
  try {
    await http.get('v1/war_thunder/userinfo',
        {
          params: {
            "query_id": queryId
          }
        }).then(res => {
      gaijinData.value = res.data['gaijin']
      thunderskillData.value = res.data['thunder_skill']
    })
  } finally {
    reloading.value = false
  }
}

const message = useMessage();
const refreshInfoQueries = (nick: any) => {
  http.post('v1/war_thunder/userinfo/refresh',
      {}, {
        params: {
          "nickname": nick
        }
      }).then(res => {
    if (res.data['refresh'] === true) {
      message.success("正在刷新，请稍后")
    } else {
      message.warning('同一个玩家24小时内仅能查询一次！')
    }
  })
}
</script>

<style scoped>

</style>
