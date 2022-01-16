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

      <UserInfo @refresh=""></UserInfo>
      <!--<n-empty size="huge">-->
      <!--</n-empty>-->
    </n-space>

  </n-card>

</template>

<script lang="ts" setup>
import {NCard, NAlert, NInputGroup, NInput, NDivider, NGrid, NGi} from "naive-ui";
import {Search} from "@vicons/fa";
import UserInfo from "@/views/record/components/UserInfo.vue";
import {ref} from "vue";

const nick = ref()

</script>

<script lang="ts">
export default {
  name: "Index",
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
