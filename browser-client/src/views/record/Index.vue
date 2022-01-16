<template>
  <n-card class="record" :bordered="false">
    <n-space vertical>
      <n-alert title="查询到的数据来自官网和Thunder Skill，本站不对数据的可靠性负责" type="warning"/>
      <n-alert title="如果发现数据错位情况，请进行反馈" type="info"/>
      <n-alert title="为了防止滥用，同个游戏用户的最新数据每24小时只会刷新一次" type="warning"/>
      <n-grid item-responsive>
        <n-gi offset="0 768:8 1200:8 1920:8" span="24 768:8 1200:8 1920:8">
          <n-input-group>
            <n-input maxlength="20" show-count size="large" round
                     :style="{ width: '100%' }" placeholder="请输入游戏昵称"/>
            <n-button size="large" type="primary" round>
              <template #icon>
                <Search/>
              </template>
              查询
            </n-button>
          </n-input-group>
        </n-gi>
      </n-grid>
      <n-divider/>
      <UserInfo></UserInfo>
      <!--<n-empty size="huge">-->
      <!--</n-empty>-->
    </n-space>

  </n-card>

</template>

<script lang="ts" setup>
import {NCard, NEmpty, NAlert, NInputGroup, NInput, NDivider, NGrid, NGi} from "naive-ui";
import UserInfo from "@/views/record/components/UserInfo.vue";

</script>

<script lang="ts">
import {Search} from "@vicons/fa";

export default {
  name: "Index",
  components: {
    Search
  }
}
</script>

<style scoped>
.record {
  text-align: left;
}

:deep(.n-input-group){
  justify-content: center;
}
</style>
