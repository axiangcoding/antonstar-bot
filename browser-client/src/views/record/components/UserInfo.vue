<template>
  <n-card embedded :bordered="true">
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
          <CrawlerInfo :data="queryList" :active="activeQuery" @searchQuery="getInfo"></CrawlerInfo>
        </n-gi>
      </n-grid>
      <n-grid cols="1 768:3 1200:2 1920:3" :x-gap="12" :y-gap="8">
        <n-gi v-for="(index,key) in gaijinData.user_stat" :key="key">
          <GaijinStatCard :data="index" :title="key"></GaijinStatCard>
        </n-gi>
      </n-grid>
      <n-tabs type="line">
        <n-tab-pane name="aviation" tab="空军">
          <n-grid cols="1 768:3 1200:2 1920:3" :x-gap="12" :y-gap="8">
            <n-gi>
              <!--<CommonInfo></CommonInfo>-->
            </n-gi>
            <n-gi>
              <!--<CommonInfo></CommonInfo>-->
            </n-gi>
            <n-gi>
              <!--<CommonInfo></CommonInfo>-->
            </n-gi>
          </n-grid>
        </n-tab-pane>
        <n-tab-pane name="陆军" tab="陆军">
          <n-grid cols="1 768:3 1200:2 1920:3" :x-gap="12" :y-gap="8">
            <n-gi>
              <!--<CommonInfo></CommonInfo>-->
            </n-gi>
            <n-gi>
              <!--<CommonInfo></CommonInfo>-->
            </n-gi>
            <n-gi>
              <!--<CommonInfo></CommonInfo>-->
            </n-gi>
          </n-grid>
        </n-tab-pane>
        <n-tab-pane name="海军" tab="陆军">
          <n-grid cols="1 768:3 1200:2 1920:3" :x-gap="12" :y-gap="8">
            <n-gi>
              <!--<CommonInfo></CommonInfo>-->
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
    <n-space vertical v-else-if="displayStatus==='nothing'">
      <n-empty>大佬，未找到，是否搜索这小子？</n-empty>
    </n-space>
    <n-space vertical v-else>
      <n-skeleton text :repeat="5" :animated="false"></n-skeleton>
    </n-space>
  </n-card>
</template>

<script lang="ts" setup>
import {NCard, NGrid, NGi, NSpace, NTabs, NTabPane, NDivider} from "naive-ui";
import CommonInfo from "@/views/record/components/CommonInfo.vue";
import GaijinStatCard from "@/views/record/components/GaijinStatCard.vue";
import CrawlerInfo from "@/views/record/components/CrawlerInfo.vue";
import {ref, watch} from "vue";
import http from "@/services/request";

const props = defineProps({
  queryList: Object
});

const activeQuery = ref()
const gaijinData = ref({})
const thunderskillData = ref({})

const displayStatus = ref('none')
watch(props, (newVal, oldVal) => {
  if (props.queryList !== undefined && props.queryList.gaijin !== undefined) {
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
      displayStatus.value = 'nothing'
    }
  } else {
    displayStatus.value = 'nothing'
  }
})

const getInfo = async (queryId: string) => {
  activeQuery.value = queryId
  await http.get('v1/war_thunder/userinfo',
      {
        params: {
          "query_id": queryId
        }
      }).then(res => {
    gaijinData.value = res.data['gaijin']
    thunderskillData.value = res.data['thunder_skill']
  })
}


</script>

<script lang="ts">
export default {
  name: "UserInfo",
}
</script>

<style scoped>

</style>
