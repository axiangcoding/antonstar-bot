<template>
  <div id="home">
    <n-layout>
      <n-layout-header>
       <home-header/>
      </n-layout-header>
      <n-layout-content content-style="margin:10px 10px">
        <n-grid :x-gap="12" :y-gap="8" item-responsive cols="4">

          <n-gi span="0:4 600:4 800:2 1400:1">
            <bot-status/>
          </n-gi>
          <n-gi span="0:4 600:4 800:2 1400:1">
            <kook-bot-status/>
          </n-gi>
          <n-gi span="0:4 600:4 800:2 1400:1">
            <get-help/>
          </n-gi>
        </n-grid>
      </n-layout-content>
    </n-layout>
  </div>
</template>

<script setup lang="ts">


import {ref} from "vue";
import BotStatus from "@/page/home/components/QQBotStatus.vue";

import KookBotStatus from "@/page/home/components/KookBotStatus.vue";
import HomeHeader from "@/page/home/components/HomeHeader.vue";
import GetHelp from "@/page/home/components/GetHelp.vue";

const testData = ref([{
  title: "机器人在线状态",

}])

</script>

<style scoped>
#home {
}

.n-layout-header {
  background: #ffffff;
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 70px;
  padding: 0 10px;
}

.n-layout-content {
  background: inherit;
  background-image: linear-gradient(135deg, #ABDCFF 10%, #0396FF 100%);
  height: calc(100vh - 140px);
}


</style>