<template>
  <n-card :content-style="{padding:0}" :bordered="false">
    <n-carousel autoplay draggable>
      <n-image height="240" width="1200" object-fit="cover" v-for="item in imgList" :key="item.index" :src="item.src"/>
    </n-carousel>
    <n-card :content-style="{padding:'4px'}" :bordered="false">
      <n-space vertical style="text-align: left">
        <Notice/>
        <Statistics/>
        <Links></Links>
      </n-space>
    </n-card>

  </n-card>


</template>

<script setup lang="ts">
import {ref} from "vue";
import Statistics from "@/views/home/components/Statistics.vue";
import Notice from "@/views/home/components/Notice.vue";
import Links from "@/views/home/components/Links.vue";

const imgList = ref([
  {
    index: 1,
    src: 'https://pic.imgdb.cn/item/6242dc5927f86abb2a3bb53b.png'
  },
  {
    index: 2,
    src: 'https://pic.imgdb.cn/item/6242dc5927f86abb2a3bb6d9.png'
  },
  {
    index: 3,
    src: 'https://pic.imgdb.cn/item/6242dc5a27f86abb2a3bb898.png'
  },
  {
    index: 4,
    src: 'https://pic.imgdb.cn/item/6242dc5b27f86abb2a3bba78.png'
  },
])
</script>

<style lang="scss" scoped>
.carousel-img {
  width: 100%;
  height: 240px;
  object-fit: cover;
}
</style>
