<template>
  <n-card style="text-align: left" :bordered="false">
    <n-space vertical>
      <n-page-header @back="handleBack">
        <template #title>
          {{ route.params.nick }} 的战绩
        </template>
        <template #avatar>
          <n-avatar src="https://pic.imgdb.cn/item/6242d7ae27f86abb2a2fda29.png"></n-avatar>
        </template>
        <template #extra>
          <n-button @click="refreshInfo" type="success">刷新
            <template #icon>
              <n-icon>
                <SunRegular/>
              </n-icon>
            </template>
          </n-button>
        </template>
      </n-page-header>
      <UserInfo v-if="render" :nick="route.params.nick"></UserInfo>
    </n-space>
  </n-card>

</template>

<script lang="ts" setup>
import {useStore} from "vuex";
import {useRoute, useRouter} from "vue-router";
import UserInfo from "@/views/player/components/UserInfo.vue";
import {nextTick, onMounted, ref} from "vue";
import {SunRegular} from "@vicons/fa";

const store = useStore();
const route = useRoute();
const router = useRouter();

const render = ref(true)
const handleBack = function () {
  router.back()
}

const refreshInfo = () => {
  render.value = false
  nextTick(() => {
    render.value = true
  })
}

onMounted(() => {
  document.title = document.title + " - " + route.params.nick
})

</script>

<style scoped>

</style>