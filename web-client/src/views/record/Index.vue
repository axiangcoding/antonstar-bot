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
        <el-divider></el-divider>
        <el-input
            style="width: 40%"
            v-model="nick"
            placeholder="请输入待查询的游戏昵称"
            class="input-with-select"
            size="large"
        >
          <template #append>
            <el-button type="primary">查询</el-button>
          </template>
        </el-input>
        <el-divider></el-divider>
        <UserInfo></UserInfo>
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
import {ref} from "vue";

const nick = ref()
</script>

<style scoped>

</style>
