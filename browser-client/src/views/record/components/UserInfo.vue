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
        <n-result size="small" status="404" title="未在本站中找到该用户" description="未在本站中找到记录并不代表该用户真的不存在，点击下面的按钮查询试试？">
          <template #footer>
            <n-button type="info" secondary @click="refreshInfoQueries(route.params.nick)">点我向官网发起查询</n-button>
          </template>
        </n-result>
      </n-space>
      <n-space align="center" vertical v-else-if="displayStatus==='running'">
        <n-gradient-text :size="20" type="warning">
          正在向官网查询中...
        </n-gradient-text>
        <n-gradient-text :size="14" gradient="linear-gradient(90deg, red 0%, green 50%, purple 100%)">程序本没有慢，查的人多了，就变成了慢</n-gradient-text>
      </n-space>
      <n-space vertical v-else-if="displayStatus==='notfound'">
        <n-result size="small" status="404" title="未在官网中找到该用户" description="这下是真找不到了，是不是名字输错了？">
          <template #footer>
            <n-button type="error" secondary @click="refreshInfoQueries(route.params.nick)">我不信，我要再查</n-button>
          </template>
        </n-result>

      </n-space>
      <n-space vertical v-else>
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
    let done = props.queryList.gaijin[0].status === 'done'
    let queryId = props.queryList.gaijin[0].query_id
    for (let item of props.queryList.gaijin) {
      if (item.found) {
        found = true
        done = item.status === 'done'
        console.log(item.done);
        queryId = item.query_id
        break
      }
    }
    console.log(done);
    if (found && done) {
      displayStatus.value = 'find'
      getInfo(queryId)
    } else if (!found && !done) {
      displayStatus.value = 'running'
    } else {
      displayStatus.value = 'notfound'
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
