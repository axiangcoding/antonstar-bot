<template>
  <n-layout-header position="absolute">
    <n-grid item-responsive>
      <!-- 左侧栅栏 -->
      <n-grid-item span="0 768:1 1200:2 1920:4">
        <div class="logo">
          <n-image height="46" :src="LOGO_URL" preview-disabled/>
        </div>
      </n-grid-item>
      <n-grid-item span="24 768:22 1200:20 1920:16" class="header-content">
        <n-menu v-model:value="activeKey" mode="horizontal" :options="menuOptions"/>
      </n-grid-item>
      <!-- 右侧栅栏 -->
      <n-grid-item span="0 768:1 1200:2 1920:4">
        <n-avatar
            size="large"
            src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"
        />
      </n-grid-item>
    </n-grid>
  </n-layout-header>
</template>


<script lang="ts" setup>
// const internalInstance = getCurrentInstance()
// const globalProperties =
// 	internalInstance?.appContext.config.globalProperties
// const anime = globalProperties?.anime
import {useRoute, RouterLink} from "vue-router";
import {computed, h, onBeforeUnmount, onMounted, ref} from "vue";
import {NIcon, NMenu} from "naive-ui";
import logo from "@/assets/logo/logo.png";
import {Award} from "@vicons/fa";

let header: any = null
let scrollBox: any = null
let scrollFunc: any = null

const route = useRoute()
const activeKey = computed(() => {
  return route.name as string
})

function renderIcon(icon: any) {
  return () => h(NIcon, null, {default: () => h(icon)})
}

const menuOptions = [
  {
    key: 'record',
    icon: renderIcon(Award),
    label: () =>
        h(
            RouterLink,
            {
              to: {
                name: 'record',
                params: {
                  nick: 'zh-CN'
                }
              }
            },
            {default: () => '战绩查询'}
        ),
  },
  {
    key: 'about',
    icon: renderIcon(Award),
    label: () =>
        h(
            RouterLink,
            {
              to: {
                name: 'record',
                params: {
                  nick: 'zh-CN'
                }
              }
            },
            {default: () => '关于我们'}
        ),
  }
]
onMounted(() => {
  header = document
      .querySelector('.n-layout .n-layout-header')
  scrollBox = document
      .querySelector('.n-layout .n-layout-scroll-container')
  scrollBox?.addEventListener(
      'scroll',
      (scrollFunc = (event: any) => {
        const scrollTop = event.target.scrollTop
        console.log(header?.className)
        if (scrollTop == 0) header.classList.add('immersive')
        else header.classList.remove('immersive')
      })
  )
})
onBeforeUnmount(() => {
  scrollBox?.removeEventListener('scroll', scrollFunc)
})
const LOGO_URL = ref(logo)

</script>

<script lang="ts">
import {NLayoutHeader} from 'naive-ui'
import {defineComponent} from 'vue'

export default defineComponent({
  components: {
    NLayoutHeader,
  },
})
</script>

<style lang="scss" scoped>
.n-layout-header {
  height: $headerHeight;
  z-index: 10;
  display: flex;
  align-items: center;
  background-color: rgba($color: #d4e2eb, $alpha: 0.8);
  transition: 0.3s all;

  &.immersive {
    padding-top: 10px;
    background: inherit !important;
  }

  :deep(.n-grid) {
    margin: auto 0;
  }

  .header-content {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  :deep(.n-image) {
    display: block;

    img {
      display: block;
    }
  }
}
</style>
