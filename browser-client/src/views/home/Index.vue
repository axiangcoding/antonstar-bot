<template>
  <n-card :content-style="{padding:0}" :bordered="false">
    <n-carousel autoplay draggable>
      <n-image height="240" width="1200" object-fit="cover" v-for="item in imgList" :key="item.index" :src="item.src"/>
    </n-carousel>
    <n-space vertical>
      <Notice/>
      <Statistics/>
    </n-space>
  </n-card>


</template>

<script setup lang="ts">
import {ref} from "vue";
import Statistics from "@/views/home/components/Statistics.vue";
import Notice from "@/views/home/components/Notice.vue";

const imgList = ref([
  {
    index: 1,
    src: 'https://pic.imgdb.cn/item/61ec14312ab3f51d913cbd09.jpg'
  },
  {
    index: 2,
    src: 'https://pic.imgdb.cn/item/61ec14312ab3f51d913cbd01.jpg'
  },
  {
    index: 3,
    src: 'https://pic.imgdb.cn/item/61ec19472ab3f51d9142c2bc.jpg'
  },
  {
    index: 4,
    src: 'https://pic.imgdb.cn/item/61ec19472ab3f51d9142c2b4.png'
  },
  {
    index: 5,
    src: 'https://pic.imgdb.cn/item/61ec19472ab3f51d9142c2ad.png'
  }
])
</script>

<style lang="scss" scoped>
.carousel-img {
  width: 100%;
  height: 240px;
  object-fit: cover;
}
</style>
