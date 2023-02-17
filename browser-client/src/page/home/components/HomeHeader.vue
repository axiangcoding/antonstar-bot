<template>
  <div style="display: inherit">
    <n-image preview-disabled height="56" :src="logoNoText"></n-image>

    <p class="welcome-text">欢迎来到安东星！</p>
  </div>

  <n-space>
    <!--<n-button @click="router.push('/query')" type="info">-->
    <!--  <template #icon>-->
    <!--    <find-in-page-outlined/>-->
    <!--  </template>-->
    <!--  战绩数据-->
    <!--</n-button>-->
    <n-button type="primary">
      <template #icon>
        <settings-filled/>
      </template>
      管理后台
    </n-button>
  </n-space>
</template>

<script setup lang="ts">
import logoNoText from '@/assets/logo/logo_no_text.png'
import {SettingsFilled, FindInPageOutlined} from "@vicons/material";
import {useRouter} from "vue-router";

const router = useRouter();
</script>

<style scoped>

.welcome-text {
  font-size: 16px;
  margin: auto 10px;
}

</style>