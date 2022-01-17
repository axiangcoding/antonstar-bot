<template>
  <n-card class="record" :bordered="false">
    <n-space vertical>
      <n-alert title="使用须知" type="warning">
        1. 查询到的数据来自官网和Thunder Skill，本站不对数据的可靠性负责
        <br/>
        2. 为了防止滥用，同个游戏用户的最新数据每24小时只会刷新一次
        <br/>
        3. 如果发现数据错位情况，请进行反馈
      </n-alert>
      <n-grid item-responsive>
        <n-gi offset="0 768:6 1200:6 1920:6" span="24 768:12 1200:12 1920:12">

          <n-input-group>
            <n-input v-model:value="nick" maxlength="20" show-count size="large" round
                     :style="{ width: '100%' }" placeholder="请输入游戏昵称"/>
            <n-button @click="doSearch" :loading="btnLoading" size="large" type="primary" round
                      :disabled="nick.length===0">
              <template #icon>
                <Search/>
              </template>
              查询
            </n-button>
          </n-input-group>
        </n-gi>
      </n-grid>
      <n-divider/>
      <UserInfo :query-list="queryIdList"></UserInfo>
    </n-space>
  </n-card>
</template>

<script lang="ts" setup>
import {useMessage} from "naive-ui";
import {Search} from "@vicons/fa";
import UserInfo from "@/views/record/components/UserInfo.vue";
import {ref} from "vue";
import http from "@/services/request";
import {useRouter} from "vue-router";

const nick = ref('WT_GodFather')
const message = useMessage()
const showInfo = ref('none')
const btnLoading = ref(false)
let messageReactive = null

const doSearch = async () => {
  messageReactive = message.loading('正在查询，请稍后', {duration: 0})
  btnLoading.value = true
  try {
    await getInfoQueries(nick.value)
  } finally {
    messageReactive.destroy()
    messageReactive = null
    btnLoading.value = false
  }

}

const router = useRouter();
const queryIdList = ref()
const getInfoQueries = async (nick: string) => {
  await http.get('v1/war_thunder/userinfo/queries',
      {
        params: {
          "nickname": nick
        }
      }).then(res => {
    queryIdList.value = res.data
    router.push({
      name: 'record', params: {
        nick: nick
      }
    })
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


</script>

<style scoped>
.record {
  /* min-height: 2000px; */
  text-align: left;
}

:deep(.n-input-group) {
  justify-content: center;
}
</style>
