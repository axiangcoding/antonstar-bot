<template>
  <n-card embedded :bordered="true">
    <n-space vertical>
      <n-grid cols="1 768:2 1200:2 1920:2" :x-gap="12" :y-gap="8">
        <n-gi>
          <CommonInfo></CommonInfo>
        </n-gi>
        <n-gi>
          <CommonInfo></CommonInfo>
        </n-gi>
      </n-grid>
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
import {NCard, NGrid, NGi, NSpace, NTabs, NTabPane} from "naive-ui";
import CommonInfo from "@/views/record/components/CommonInfo.vue";
</script>

<script lang="ts">
export default {
  name: "UserInfo"
}
</script>

<style scoped>

</style>
