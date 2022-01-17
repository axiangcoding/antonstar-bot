<template>
  <n-card class="record" :bordered="false">
    <n-space vertical>
      <n-alert title="查询到的数据来自官网和Thunder Skill，本站不对数据的可靠性负责" type="warning"/>
      <n-alert title="如果发现数据错位情况，请进行反馈" type="info"/>
      <n-alert title="为了防止滥用，同个游戏用户的最新数据每24小时只会刷新一次" type="warning"/>
      <n-grid item-responsive>
        <n-gi offset="0 768:6 1200:6 1920:6" span="24 768:12 1200:12 1920:12">

          <n-input-group>
            <n-input v-model:value="nick" maxlength="20" show-count size="large" round
                     :style="{ width: '100%' }" placeholder="请输入游戏昵称"/>
            <n-button @click="doSearch" size="large" type="primary" round :disabled="nick.length===0">
              <template #icon>
                <Search/>
              </template>
              查询
            </n-button>
          </n-input-group>
        </n-gi>
      </n-grid>
      <n-divider/>

      <UserInfo @refresh="" :gaijin-data="gaijinInfo" :thunderskill-data="thunderskillInfo"></UserInfo>
      <!--<n-empty size="huge">-->
      <!--</n-empty>-->
    </n-space>
  </n-card>
</template>

<script lang="ts" setup>
import {useMessage} from "naive-ui";
import {Search} from "@vicons/fa";
import UserInfo from "@/views/record/components/UserInfo.vue";
import {ref} from "vue";
import http from "@/services/request";

const nick = ref('WT_GodFather')
const message = useMessage()
const showInfo = ref('notfound')
let messageReactive = null
const doSearch = async () => {
  messageReactive = message.loading('正在查询，请稍后', {duration: 0})
  try {
    await getInfoQueries(nick.value)
    let gaijinList = queryIdList['gaijin'];
    // 如果是唯一的一条记录，那么说明是第一次查询
    if (gaijinList != null && gaijinList.length == 1) {
      const item = gaijinList[0]
      if (item['status'] === 'running') {
        showInfo.value = "running"
      } else if (item['found'] === true) {
        await getInfo(item['query_id'])
        showInfo.value = "done"
      } else {
        showInfo.value = 'notfound'
      }
    }
    // 如果有多条记录，代表这个用户已经完成了多次查询
    else if (gaijinList != null && gaijinList.length > 1) {
      for (let key in gaijinList) {
        const item = gaijinList[key]
        if (item['found'] === true) {
          if (item['status'] === 'running') {
            showInfo.value = "running"
          } else if (item['found'] === true) {
            await getInfo(item['query_id'])
            showInfo.value = "done"
          } else {
            showInfo.value = 'notfound'
          }
          break;
        } else {
          showInfo.value = 'notfound'
        }
      }
    } else {
      refreshInfo(nick.value)
      showInfo.value = 'running'
    }
  } catch (e) {

  }

  console.log("gg");
  messageReactive.destroy()
  messageReactive = null
}

let queryIdList: any
const getInfoQueries = async (nick: string) => {
  await http.get('v1/war_thunder/userinfo/queries',
      {
        params: {
          "nickname": nick
        }
      }).then(res => {

    queryIdList = res.data

  })
}

const refreshInfo = (nick: string) => {
  http.post('v1/war_thunder/userinfo/refresh',
      {}, {
        params: {
          "nickname": nick
        }
      }).then(res => {
    if (res.data['refresh'] === true) {
      // ElMessage({
      //   message: '已发送查询请求',
      //   type: 'success',
      // })
    } else {
      // ElMessage({
      //   type: 'warning',
      //   message: "同一个玩家24小时内仅能查询一次！"
      // })
    }
  })
}

const gaijinInfo = ref({})
const thunderskillInfo = ref({})
const getInfo = async (queryId: string) => {
  await http.get('v1/war_thunder/userinfo',
      {
        params: {
          "query_id": queryId
        }
      }).then(res => {
    gaijinInfo.value = res.data['gaijin']
    thunderskillInfo.value = res.data['thunder_skill']
  })
}
</script>

<style scoped>
.record {
  min-height: 2000px;
  text-align: left;
}

:deep(.n-input-group) {
  justify-content: center;
}
</style>
