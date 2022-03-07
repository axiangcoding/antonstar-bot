<template>
  <n-card embedded :bordered="true">
    <n-spin :show="reloading">
      <n-space v-if="displayStatus==='find'" vertical>
        <div>
          <n-button quaternary type="error" size="small" disabled>
            <template #icon>
              <Angry/>
            </template>
            站内举报
          </n-button>
          <n-divider vertical/>
          <n-button quaternary type="info" size="small" @click="copyLink">
            <template #icon>
              <ShareAlt/>
            </template>
            分享链接
          </n-button>
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
        <TSCommonInfo :data="thunderskillData"/>
        <n-divider/>

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
            </n-grid>
          </n-tab-pane>
        </n-tabs>

      </n-space>
      <n-space align="center" vertical v-else-if="displayStatus==='nothing'">
        <n-result size="small" status="success" title="已发起查询">
          <template #default>
            <n-space vertical align="center">
              <n-gradient-text :size="14" gradient="linear-gradient(90deg, red 0%, green 50%, purple 100%)">
                第一次总是值得期待的
              </n-gradient-text>
              <n-button type="warning" ghost tag="a"
                        :href="'https://warthunder.com/zh/community/userinfo/?nick='+route.params.nick" target="_blank">
                在线等，很急，点我直接去官网
              </n-button>
            </n-space>
          </template>
        </n-result>
      </n-space>
      <n-space align="center" vertical v-else-if="displayStatus==='running'">
        <n-result size="small" status="info" title="正在查询中...">
          <template #default>
            <n-space vertical align="center">
              <n-gradient-text :size="14" gradient="linear-gradient(90deg, red 0%, green 50%, purple 100%)">
                程序本没有慢，查的人多了，就变成了慢
              </n-gradient-text>
              <n-button type="info" ghost @click="refreshInfoQueries(route.params.nick)">很久没刷出来？点我重新查询</n-button>
              <n-button type="warning" ghost tag="a"
                        :href="'https://warthunder.com/zh/community/userinfo/?nick='+route.params.nick" target="_blank">
                在线等，很急，点我直接去官网
              </n-button>
            </n-space>
          </template>
        </n-result>
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
import {ShareAlt, Angry} from "@vicons/fa";
import TSCommonInfo from "@/views/record/components/TSCommonInfo.vue";
import {getWTUserInfo, postWTUserInfoRefresh} from "@/services/war_thunder";
import {useStore} from "vuex";

const props = defineProps({
  queryList: Object
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
        queryId = item.query_id
        break
      }
    }
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
    refreshInfoQueries(route.params.nick)
  }
})


const reloading = ref(false)
const getInfo = async (queryId: string) => {
  reloading.value = true
  activeQuery.value = queryId
  try {
    await getWTUserInfo(queryId).then(res => {
      gaijinData.value = res.data['gaijin']
      thunderskillData.value = res.data['thunder_skill']
    })
  } finally {
    reloading.value = false
  }
}

const message = useMessage();
const store = useStore();
const refreshInfoQueries = (nick: any) => {
  postWTUserInfoRefresh(store.state.auth, nick).then((res: any) => {
    if (res.code === 13000) {
      message.warning('无法刷新，已达到今天的全站限额！')
      return
    }
    if (res.data['refresh'] === true) {
      message.success("正在获取最新快照，请稍后")
    } else {
      message.warning('同一个玩家24小时内只能刷新一次！')
    }
  })
}


const copyLink = async () => {
  if (!navigator.clipboard) {
    message.error('复制到剪切板失败，请直接复制网页地址')
    return
  }
  navigator.clipboard.writeText(window.location.href).then(() => {
    message.success('链接已复制到剪切板，欢迎分享！')
  }).catch(() => {
    message.error('复制到剪切板失败，请直接复制网页地址')
  })

}


const toGaijin = () => {

}
</script>

<style scoped>

</style>
