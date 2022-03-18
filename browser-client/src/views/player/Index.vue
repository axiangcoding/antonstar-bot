<template>
  <n-card style="text-align: left" :bordered="false">
    <n-space vertical>
      <n-page-header @back="handleBack">
        <template #title>
          {{ route.params.nick }} 的战绩
        </template>
        <template #avatar>
          <n-avatar src="https://pic.imgdb.cn/item/622ea5435baa1a80ab949b2d.png"></n-avatar>
        </template>
      </n-page-header>
      <UserInfo :nick="route.params.nick"></UserInfo>
    </n-space>
  </n-card>

</template>

<script lang="ts" setup>
import {useStore} from "vuex";
import {useRoute, useRouter} from "vue-router";
import UserInfo from "@/views/player/components/UserInfo.vue";

const store = useStore();
const route = useRoute();
const router = useRouter();


const handleBack = function () {
  router.back()
}

</script>

<style scoped>

</style>