<template>
  <n-card class="record">
    <n-space vertical>
      <n-alert title="查询到的数据来自官网和Thunder Skill，本站不对数据的可靠性负责" type="warning"/>
      <n-alert title="如果发现数据错位情况，请进行反馈" type="info"/>
      <n-alert title="为了防止滥用，同个游戏用户的最新数据每24小时只会刷新一次" type="warning"/>
      <div style="width:100%;">
        <n-input-group>
          <n-input :style="{ width: '50%' }"/>
          <n-button type="primary">搜索</n-button>
        </n-input-group>
      </div>

      <n-empty size="huge">

      </n-empty>
      <n-button>Oops!</n-button>
      <n-button>Oops!</n-button>
    </n-space>

  </n-card>

</template>

<script lang="ts" setup>
import {NCard, NEmpty, NAlert, NInputGroup, NInput} from "naive-ui";
</script>

<script lang="ts">
export default {
  name: "Index"
}
</script>

<style scoped>
.record {
  text-align: left;
}

:deep(.n-input-group){
  margin: 0 auto;
}
</style>
