<template>
  <n-card :bordered="false" size="small">
    <template #header>
      <n-h2 prefix="bar" align-text type="primary">友情链接</n-h2>
    </template>
    <n-h4>
      <ul>
        <li v-for="i in displayList" :key="i.url">
          {{ i.name }} -
          <n-button text size="large" type="info" @click="openWin(i.url,'_blank')">
            {{ i.url }}
          </n-button>
        </li>
      </ul>
    </n-h4>


  </n-card>
</template>

<script lang="ts" setup>
import {openWin} from "@/util/window";
import {ref} from "vue";

const displayList = ref([
  {name: '官方WiKi', url: 'https://wiki.warthunder.com/Main_Page'},
  {name: '分房热力图', url: 'https://wt.controlnet.space/'},
  {name: 'Thunder Skill', url: 'https://thunderskill.com/en'},
  {name: 'BiliGame WiKi', url: 'https://wiki.biligame.com/warthunder/'},
  {name: '小老板 摸鱼的33 的b站直播间', url: 'https://live.bilibili.com/5295790'}
])
</script>

<style scoped>

</style>
