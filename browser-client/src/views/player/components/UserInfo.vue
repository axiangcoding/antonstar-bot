<template>
  <n-card embedded :bordered="true">
    <n-spin :show="reloading">
      <n-space v-if="displayStatus==='found'" vertical>
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
                         @refresh="refreshInfoQueries(route.params.nick, false)"></CrawlerInfo>
          </n-gi>
        </n-grid>
        <TSCommonInfo :data="thunderskillData"/>

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
      <n-space align="center" vertical v-else-if="displayStatus==='first'">
        <n-result size="small" status="success" title="已发起查询，正在生成玩家战绩数据">
          <template #default>
            <n-space vertical align="center">
              <n-gradient-text :size="14" gradient="linear-gradient(90deg, red 0%, green 50%, purple 100%)">
                第一次总是值得期待的
              </n-gradient-text>
              <n-button type="primary" @click="refreshInfoQueries(nick)">等待太久，重新更新下</n-button>
              <n-button type="info" ghost tag="a"
                        :href="'https://warthunder.com/zh/community/userinfo/?nick='+nick" target="_blank">
                进入Gaijin官网确认
              </n-button>
            </n-space>
          </template>
        </n-result>
      </n-space>
      <n-space vertical v-else-if="displayStatus==='notExist'">
        <n-result size="small" status="404" title="该玩家不存在" description="请检查是否昵称拼写错误？">
          <template #footer>
            <n-space vertical align="center">
              <n-button type="primary" @click="refreshInfoQueries(nick)">我不信，我要重新查询</n-button>

              <n-button type="info" ghost tag="a"
                        :href="'https://warthunder.com/zh/community/userinfo/?nick='+nick" target="_blank">
                进入Gaijin官网确认
              </n-button>
            </n-space>
          </template>
        </n-result>
      </n-space>
      <n-space vertical v-else>
        <div>Q: 为什么我无法查询 {{ nick }} 的战绩？</div>
        <div>A: 如果该玩家是第一次被查询，那么必须要
          <n-button size="small" type="primary" text @click="router.push({name:'login'})">登录</n-button>
          才能查。如果该玩家的战绩已在安东星上已有记录，那么无需登录就能查看
        </div>
        <div>Q: 为什么要登录才能查战绩？</div>
        <div>A: 因为查战绩需要向官网发起查询，如果游客都能使用，太过频繁会导致战绩查询无法使用。同时，成为安东星的用户可以为网站带来客观的增益，站长才会有动力更新更多的功能...</div>
        <div>Q: 垃圾网站，要我输入账号密码肯定是想盗我 DK3，1500天大会员的战雷帐号，给👴爬</div>
        <div>A: 请注意，安东星的账号是独立的，和包括gaijin游戏账号在内的其他任何地方没有一毛钱关系。你需要在安东星上注册一个全新的账号。为了安全，你应该设置一个完全独立的用户名和密码</div>
      </n-space>
    </n-spin>
  </n-card>
</template>

<script lang="ts" setup>
import {NCard, NGrid, NGi, NSpace, NTabs, NTabPane, NDivider, useMessage, useDialog} from "naive-ui";
import CommonInfo from "@/views/player/components/CommonInfo.vue";
import GaijinStatCard from "@/views/player/components/GaijinStatCard.vue";
import CrawlerInfo from "@/views/player/components/CrawlerInfo.vue";
import {onMounted, ref, watch} from "vue";
import {useRoute, useRouter} from "vue-router";
import GaijinAviationCard from "@/views/player/components/GaijinAviationCard.vue";
import GaijinGroundCard from "@/views/player/components/GaijinGroundCard.vue";
import GaijinFleetCard from "@/views/player/components/GaijinFleetCard.vue";
import {ShareAlt, Angry} from "@vicons/fa";
import TSCommonInfo from "@/views/player/components/TSCommonInfo.vue";
import {getWTUserInfo, getWTUserInfoQueries, postWTUserInfoRefresh} from "@/services/war_thunder";
import {useStore} from "vuex";

const props = defineProps({
  nick: String
});

const message = useMessage();
const dialog = useDialog();
const store = useStore();
const router = useRouter();
const route = useRoute();

const activeQuery = ref()

const gaijinData = ref({})
const thunderskillData = ref({})

const displayStatus = ref('nothing')

const queryList = ref()

onMounted(() => {
  getWTUserInfoQueries(store.state.auth, props.nick as string).then(res => {
    let qList
    queryList.value = res.data
    qList = res.data
    if (qList != undefined && qList.gaijin != undefined) {
      // 是否有找到的记录
      let found = false
      // 是否有执行完的查询
      let done = false
      let queryId
      for (let item of qList.gaijin) {
        done = (done || item.status === 'done')
        found = (found || item.found as boolean)
        if (item.found) {
          queryId = item.query_id
          break
        }
      }
      if (found && done) {
        displayStatus.value = 'found'
        getInfo(queryId)
      } else if (!found && done) {
        displayStatus.value = 'notExist'
      } else {
        displayStatus.value = 'first'
      }
    } else {
      refreshInfoQueries(props.nick)
    }
  })
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


const refreshInfoQueries = (nick: any, jumpNone?: boolean) => {
  postWTUserInfoRefresh(store.state.auth, nick).then((res: any) => {
    if (res.code === 13000) {
      message.warning('无法刷新，已达到今天的全站限额！')
      return
    }
    if (res.data['refresh'] === true) {
      message.success("正在获取最新快照，请稍后")
      router.go(0)
    } else {
      message.warning('同一个玩家24小时内只能刷新一次！')
    }
  }).catch(err => {
    if (err.response.status == 401) {
      if(jumpNone){
        displayStatus.value = 'none'
      }
      dialog.warning({
        title: '查询受限',
        content: '对不起，该玩家的战绩无法刷新。请登录后再访问',
        closable: false,
        positiveText: '登 录',
        negativeText: '取消',
        onPositiveClick: () => {
          router.push({name: 'login'})
        }
      })
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

</script>

<style scoped>

</style>
