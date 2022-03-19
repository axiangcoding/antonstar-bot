<template>
  <n-card class="record" :bordered="false">
    <n-space vertical>
      <n-alert title="战绩查询使用须知" type="info">
        1. 查询到的数据来自官网和Thunder Skill，本站不对数据的可靠性负责
        <br/>
        2. 昵称大小写敏感，注意不要写错
        <br/>
        3. 如果发现数据错位情况，请点击页面最下方的问题反馈栏目进行反馈
        <br/>
        4. 由于发现了滥用的情况，目前开启全站查询限制，本站一天仅允许100次查询，后续将进一步调整，内测期间请谅解！
        <br/>
        5. 如果是第一次查询，或者是刷新快照，需要登录才能使用！敬请谅解！
      </n-alert>

      <n-grid item-responsive>
        <n-gi offset="0 768:6 1200:6 1920:6" span="24 768:12 1200:12 1920:12">
          <n-space vertical>
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

            <n-space v-if="store.state.login && searchHistory!=null" justify="space-between">
              <n-space>
                搜索历史：
                <n-button v-for="i in searchHistory.slice(0, 3)" size="tiny" @click="handleSearch(i)" secondary
                          type="info" :key="i">{{ i }}
                </n-button>
              </n-space>
              <n-button text type="primary" @click="showMore=!showMore" v-if="searchHistory.length>5">
                查看更多
              </n-button>
            </n-space>
            <n-collapse-transition :show="showMore && searchHistory!=null">
              <n-space>
                <n-button v-for="i in searchHistory.slice(3)" size="tiny" @click="handleSearch(i)" secondary
                          type="info" :key="i">{{ i }}
                </n-button>
              </n-space>
            </n-collapse-transition>
          </n-space>
        </n-gi>
      </n-grid>

    </n-space>
    <n-divider/>
    <n-space vertical>
      <n-h3 prefix="bar" align-text>
        <n-text type="primary">
          本站已记录玩家
        </n-text>
      </n-h3>
      <n-grid cols="1 768:3 1200:2 1920:3" :x-gap="12" :y-gap="16">
        <n-gi v-for="i in gameUserList" :key="i.nick">
          <PlayerCard :data="i"/>
        </n-gi>
      </n-grid>
      <n-space justify="end">
        <n-pagination v-model:page="pagination.pageNum" :page-size="pagination.pageSize" :item-count="total"
                      :page-slot="7" @update:page="pageUpdate"/>
      </n-space>

    </n-space>
  </n-card>
</template>

<script lang="ts" setup>
import {Search} from "@vicons/fa";
import UserInfo from "@/views/record/player/components/UserInfo.vue";
import {onMounted, ref} from "vue";
import {useRoute, useRouter} from "vue-router";
import {getWTUserInfoQueries} from "@/services/war_thunder";
import {useStore} from "vuex";
import {getWTQueryHistory} from "@/services/user";
import PlayerCard from "@/views/record/components/PlayerCard.vue";
import {getGameUsers} from "@/services/game_user";
import {Pagination} from "@/services/request";

const store = useStore();
const route = useRoute();
const nick = ref('')
const router = useRouter();

const btnLoading = ref(false)

const searchHistory = ref([])

const pagination = ref<Pagination>({
  pageSize: 6,
  pageNum: 1,
  filter: ''
})
const total = ref(0)

const gameUserList = ref([])

onMounted(() => {
  // 搜索记录
  if (store.state.login) {
    getWTQueryHistory(store.state.auth).then(res => {
      searchHistory.value = res.data.list
    })
  }
  // 获取游戏玩家列表
  getGameUsers(store.state.auth, pagination.value).then(res => {
    gameUserList.value = res.data.users;
    total.value = res.data.total
  })
})

const handleSearch = (i: string) => {
  nick.value = i
}
const showMore = ref(false)

const doSearch = () => {
  router.push({name: 'player', params: {'nick': nick.value}})
}
const pageUpdate = (page: number) => {
  getGameUsers(store.state.auth, pagination.value).then(res => {
    gameUserList.value = res.data.users;
  })

}
</script>

<style scoped lang="scss">
.record {
  text-align: left;
}

:deep(.n-input-group) {
  justify-content: center;
}


</style>
