<template>
  <div>
    <el-card>
      <el-space style="width: 100%" direction="vertical" fill>
        <el-alert title="查询到的数据来自官网和Thunder Skill，本站不对数据的可靠性负责"
                  type="warning" effect="dark" center
                  :closable="false"></el-alert>
        <el-alert title="如果发现数据错位情况，请进行反馈"
                  type="info" effect="dark" center
                  :closable="false"></el-alert>
        <el-alert title="为了防止滥用，同个游戏用户的最新数据每24小时只会刷新一次"
                  type="error" effect="dark" center
                  :closable="false"></el-alert>
        <el-divider></el-divider>
        <el-row justify="space-around">
          <el-col :xs="24" :sm="16" :md="8" :lg="8" :xl="8">
            <el-input
                style="width: 100%"
                minlength="1"
                maxlength="20"
                show-word-limit
                v-model="nick"
                placeholder="请输入待查询的游戏昵称"
                size="large"
                autofocus
                @keyup.enter.native="doSearch"
            >
              <template #append>
                <el-button type="primary" @click="doSearch" :disabled="nick.length===0">查询</el-button>
              </template>
            </el-input>
          </el-col>
        </el-row>


        <div v-if="showInfo==='done'">
          <el-divider></el-divider>
          <el-space style="width: 100%" direction="vertical" fill>
            <!--<el-space size="5" spacer="|">-->
            <!--  <el-button size="small" type="success" plain disabled>这是我</el-button>-->
            <!--  <el-button size="small" type="danger" plain disabled>举报为疑似外挂</el-button>-->
            <!--</el-space>-->
            <UserInfo :gaijinInfo="gaijinInfo" :thunderskillInfo="thunderskillInfo"></UserInfo>
          </el-space>


        </div>
        <div v-else-if="showInfo==='running'">
          <el-divider></el-divider>
          <h1>已提交查询请求，正在查询，请耐心等待</h1>
          <el-link type="primary" :icon="Refresh" @click="doSearch">数据还没出来？点我刷新</el-link>
        </div>
        <div v-else-if="showInfo==='notfound'">
          <el-divider></el-divider>
          <h1>该玩家不存在，是否填写了正确的名称？</h1>
          <el-link type="primary" :icon="Refresh" @click="refreshInfo(nick)">点我重新查询</el-link>
        </div>
      </el-space>
    </el-card>
  </div>
</template>

<script lang="ts">
import UserInfo from './components/UserInfo.vue'

export default {
  name: "Index",
  components: {
    UserInfo
  }
}
</script>

<script lang="ts" setup>
import {
  Refresh,
} from '@element-plus/icons-vue'
import {ref} from "vue";
import service from "../../util/request";
import {useRoute, useRouter} from "vue-router";
import {ElMessage} from 'element-plus'

const gaijinInfo = ref({})
const thunderskillInfo = ref({})

const route = useRoute();
const router = useRouter()

const nick = ref('')

const showInfo = ref('none')

let queryIdList: {}
const doSearch = async () => {
  await getInfoQueries(nick.value)
  let gaijinList = queryIdList['gaijin'];
  // 如果是唯一的一条记录，那么说明是第一次查询
  if (gaijinList != null && gaijinList.length == 1) {
    const item = gaijinList[0]
    console.log(item);
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
}
const getInfoQueries = async (nick: string) => {
  await service.get('v1/war_thunder/userinfo/queries',
      {
        params: {
          "nickname": nick
        }
      }).then(res => {

    queryIdList = res.data

  })
}

const refreshInfo = (nick: string) => {
  service.post('v1/war_thunder/userinfo/refresh',
      {}, {
        params: {
          "nickname": nick
        }
      }).then(res => {
    if (res.data['refresh'] === true) {
      ElMessage({
        message: '已发送查询请求',
        type: 'success',
      })
    } else {
      ElMessage({
        type: 'warning',
        message: "同一个玩家24小时内仅能查询一次！"
      })
    }
  })
}

const getInfo = async (queryId: string) => {
  await service.get('v1/war_thunder/userinfo',
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

</style>
