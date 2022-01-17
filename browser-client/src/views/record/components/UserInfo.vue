<template>
  <n-card embedded :bordered="true">
    <n-space vertical>
      <div>
        <n-button quaternary type="info" size="small" disabled>这是我</n-button>
        <n-divider vertical/>
        <n-button quaternary type="primary" size="small" disabled>给大佬点赞</n-button>
        <n-divider vertical/>
        <n-button quaternary type="error" size="small" disabled>站内举报</n-button>
        <n-divider vertical/>
      </div>
      <n-grid cols="1 768:2 1200:2 1920:2" :x-gap="12" :y-gap="8">
        <n-gi>
          <CommonInfo></CommonInfo>
        </n-gi>
        <n-gi>
          <CommonInfo></CommonInfo>
        </n-gi>
      </n-grid>
      <n-grid cols="1 768:3 1200:2 1920:3" :x-gap="12" :y-gap="8">
        <n-gi v-for="(index,key) in gaijinData['user_stat']" :key="key">
          <GaijinStatCard :data="index" :title="key"></GaijinStatCard>
        </n-gi>
      </n-grid>
      <n-tabs type="line">
        <n-tab-pane name="aviation" tab="空军">
          <n-grid cols="1 768:3 1200:2 1920:3" :x-gap="12" :y-gap="8">
            <n-gi>
              <CommonInfo></CommonInfo>
            </n-gi>
            <n-gi>
              <CommonInfo></CommonInfo>
            </n-gi>
            <n-gi>
              <CommonInfo></CommonInfo>
            </n-gi>
          </n-grid>
        </n-tab-pane>
        <n-tab-pane name="陆军" tab="陆军">
          <n-grid cols="1 768:3 1200:2 1920:3" :x-gap="12" :y-gap="8">
            <n-gi>
              <CommonInfo></CommonInfo>
            </n-gi>
            <n-gi>
              <CommonInfo></CommonInfo>
            </n-gi>
            <n-gi>
              <CommonInfo></CommonInfo>
            </n-gi>
          </n-grid>
        </n-tab-pane>
        <n-tab-pane name="海军" tab="陆军">
          <n-grid cols="1 768:3 1200:2 1920:3" :x-gap="12" :y-gap="8">
            <n-gi>
              <CommonInfo></CommonInfo>
            </n-gi>
            <n-gi>
              <CommonInfo></CommonInfo>
            </n-gi>
            <n-gi>
              <CommonInfo></CommonInfo>
            </n-gi>
          </n-grid>
        </n-tab-pane>
      </n-tabs>
    </n-space>

  </n-card>
</template>

<script lang="ts" setup>
import {NCard, NGrid, NGi, NSpace, NTabs, NTabPane, NDivider} from "naive-ui";
import CommonInfo from "@/views/record/components/CommonInfo.vue";
import GaijinStatCard from "@/views/record/components/GaijinStatCard.vue";

const props = defineProps({
  gaijinData: {},
  thunderskillData: {}
});

const refresh = () => {
  emit("refresh", props.gaijinData)
}

const emit = defineEmits(["refresh"])

</script>

<script lang="ts">
export default {
  name: "UserInfo",
}
</script>

<style scoped>

</style>
