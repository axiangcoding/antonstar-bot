<template>
  <n-card class="home-card" hoverable title="获取帮助">
    <n-list hoverable clickable>
      <n-list-item @click="openTab(repo,'_blank')">
        <n-thing title="GitHub" >
          <template #avatar>
            <n-icon>
              <github/>
            </n-icon>
          </template>
          <template #description>
            {{ repo }}
          </template>

        </n-thing>
      </n-list-item>
      <n-list-item @click="openTab(manual,'_blank')">
        <n-thing title="使用文档" >
          <template #avatar>
            <n-icon>
              <book-filled/>
            </n-icon>
          </template>
          <template #description>
            {{ manual }}
          </template>

        </n-thing>
      </n-list-item>
    </n-list>
  </n-card>
</template>

<script setup lang="ts">
import {Github} from "@vicons/fa";
import {BookFilled} from "@vicons/material";

const openTab = (url: string, target: string) => {
  window.open(url, target)
}

const repo = "https://github.com/axiangcoding/antonstar-bot"
const manual = "https://www.yuque.com/axiangcoding/anton_star/sfw0d8"
</script>

<style scoped>

</style>