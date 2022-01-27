<template>
  <n-card class="record" :bordered="false">
    <n-space vertical>
      <n-alert title="使用须知" type="warning">
        1. 查询到的数据来自官网和Thunder Skill，本站不对数据的可靠性负责
        <br/>
        2. 昵称大小写敏感，注意不要写错
        <br/>
        3. 如果发现数据错位情况，请进行反馈
        <br/>
        4. 由于发现了滥用的情况，目前开启全站查询限制，本站一天仅允许100次查询，后续将进一步调整，内测期间请谅解！
      </n-alert>
      <n-grid item-responsive>
        <n-gi offset="0 768:6 1200:6 1920:6" span="24 768:12 1200:12 1920:12">

          <n-input-group>
            <n-input v-model:value="nick" maxlength="16" show-count size="large" round
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
import {onMounted, ref} from "vue";
import http from "@/services/request";
import {useRoute, useRouter} from "vue-router";

const route = useRoute();
const nick = ref('')
const message = useMessage()
const showInfo = ref('none')
const btnLoading = ref(false)
let messageReactive = null

onMounted(() => {
  nick.value = route.params.nick as string
  if (nick.value) {
    doSearch()
  }
})

const doSearch = async () => {
  // messageReactive = message.loading('正在查询，请稍后', {duration: 0})
  btnLoading.value = true
  try {
    await getInfoQueries(nick.value as string)
  } finally {
    // messageReactive.destroy()
    // messageReactive = null
    btnLoading.value = false
  }
}


const router = useRouter();
const queryIdList = ref()
const getInfoQueries = async (nick: string) => {
  // 点击查询后，应先进行跳转，这样组件才能获得正确的nickname
  await router.push({
    name: 'record', params: {
      nick: nick
    }
  })
  await http.get('v1/war_thunder/userinfo/queries',
      {
        params: {
          "nickname": nick
        }
      }).then(async res => {
    queryIdList.value = res.data
  })
}


</script>

<style scoped>
.record {
  text-align: left;
}

:deep(.n-input-group) {
  justify-content: center;
}
</style>
