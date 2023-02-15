<template>
  <div id="home">
    <n-layout>
      <n-layout-header>
        <div style="display: inherit">
          <n-image preview-disabled height="56" :src="logoNoText"></n-image>

          <p class="welcome-text">欢迎使用安东星！请使用默认密码登录管理后台</p>
        </div>


        <n-button type="primary">
          <template #icon>
            <settings-filled/>
          </template>
          管理后台
        </n-button>
      </n-layout-header>
      <n-layout-content content-style="margin:10px 10px">
        <n-grid :x-gap="12" :y-gap="8" item-responsive cols="4">

          <n-gi span="0:4 600:4 800:2 1400:1">
            <bot-status/>
          </n-gi>
          <n-gi span="0:4 600:4 800:2 1400:1" v-for="i in 6" :key="i">
            <n-card title="卡片" hoverable>
              卡片内容 ces
            </n-card>
          </n-gi>
        </n-grid>

      </n-layout-content>
      <n-layout-footer>
        Powered by
        <n-image
            preview-disabled
            width="60"
            :src="logo"
        />

      </n-layout-footer>
    </n-layout>
  </div>
</template>

<script setup lang="ts">
import logo from '@/assets/logo/logo_text_color.png'
import logoNoText from '@/assets/logo/logo_no_text.png'
import {SettingsFilled} from "@vicons/material";
import {ref} from "vue";
import BotStatus from "@/page/home/components/BotStatus.vue";

const testData = ref([{
  title: "机器人在线状态",

}])

</script>

<style scoped>
#home {
  height: 100vh;
}

.welcome-text {
  font-size: 16px;
  margin: auto 10px;
}

.n-layout-header {
  background: #ffffff;
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 70px;
  padding: 0 10px;
}

.n-layout-footer {
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f5f5f5;
  margin: auto auto;
  height: 70px;
}

.n-layout-content {
  background: inherit;
  background-image: linear-gradient(135deg, #ABDCFF 10%, #0396FF 100%);
  height: calc(100vh - 140px);
}


</style>