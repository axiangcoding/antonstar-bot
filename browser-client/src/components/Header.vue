<template>
  <div class="headerBox">
    <n-layout-header
        class="immersive"
        position="absolute"
        v-injectThemes="'Layout'"
    >
      <div class="header-content">
        <router-link class="logo" to="/">
          <n-space size="small">
            <n-image height="44" :src="LOGO_URL" preview-disabled/>
            <n-image
                height="44"
                :src="LOGO_TEXT_URL"
                preview-disabled
            />
          </n-space>
        </router-link>
        <menu-overrides ref="menu"></menu-overrides>
        <div class="right">
          <n-avatar
              v-if="$store.state.login"
              round
              size="large"
              src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"
          />
          <n-button v-else ghost type="primary" @click="router.push({name:'login'})">
            <template #icon>
              <User/>
            </template>
            登录 / 注册
          </n-button>
        </div>
      </div>
    </n-layout-header>
    <div class="rope" :class="{'pushDown': hasPush, 'pushUp': !hasPush}" @click="handlePushMenu">
      <div class="ropeLine"></div>
      <div class="laugh">
        <n-icon size="30">
          <LaughWinkRegular/>
        </n-icon>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {onBeforeUnmount, onMounted, ref} from 'vue'
import logo from '@/assets/logo/logo_no_text.png'
import logoText from '@/assets/logo/logo_text_white.png'
import MenuOverrides from './MenuOverrides.vue'
import {LaughWinkRegular, User} from '@vicons/fa'
import {useRouter} from "vue-router";

let header: any = null
let scrollBox: any = null
let scrollFunc: any = null

const menu = ref(null)

let hasPush = ref(false)
const router = useRouter();
const handlePushMenu = () => {
  const submenu = document.querySelector('.n-menu .n-submenu')
  if (submenu) {
    menu.value.changeExpand()
    hasPush.value = !hasPush.value
  }
}

onMounted(() => {
  header = document.querySelector('.n-layout .n-layout-header')
  scrollBox = document.querySelector('.n-layout .n-layout-scroll-container')
  scrollBox?.addEventListener(
      'scroll',
      (scrollFunc = (event: any) => {
        const scrollTop = event.target.scrollTop
        if (scrollTop == 0) header.classList.add('immersive')
        else header.classList.remove('immersive')
      })
  )
})
onBeforeUnmount(() => {
  scrollBox?.removeEventListener('scroll', scrollFunc)
})
const LOGO_URL = ref(logo)
const LOGO_TEXT_URL = ref(logoText)
</script>

<style lang="scss" scoped>
.headerBox {
  position: absolute;
  top: 0;
  width: 100%;
  height: var(--header-height);

  .rope {
    display: none;
    cursor: pointer;

    .ropeLine {
      width: 2px;
      height: 140px;
      background: #FCA706dd;
      transition: 0.3s transform;
    }

    .laugh {
      color: #FCA706dd;
      margin-top: -1px;
    }
  }

  @include xs {
    .rope {
      display: flex;
      flex-direction: column;
      align-items: center;
      position: absolute;
      left: 20px;
      z-index: 1;
      top: -30%;
      animation-timing-function: cubic-bezier(0.42, 0, 1, 1);

      &.pushDown {
        animation: 0.3s ease-in-out pushDown forwards;
      }

      &.pushUp {
        animation: 0.3s ease-in-out pushUp forwards;
      }
    }
    #menuOverides {
      opacity: 0;
      position: absolute;
      left: 60px;
    }
  }
}

.n-layout-header {
  height: var(--header-height);
  z-index: 10;
  display: flex;
  align-items: center;
  transition: 0.3s all;
  box-sizing: content-box;
  padding: 3px 0;

  .hidden {
    display: none;
  }

  :deep(.n-space) {
    flex-wrap: nowrap !important;
  }

  .logo {
    display: block;
    cursor: pointer;
  }

  &.immersive {
    padding-top: 8px;
    background-color: var(--header-transparent-color) !important;
  }

  :deep(.n-grid) {
    margin: auto 0;
  }

  .header-content {
    max-width: 1200px;
    display: flex;
    align-items: center;
    width: 100%;
    margin: 0 auto;
    padding: 0 5px;
    box-sizing: border-box;
  }

  @keyframes pushDown {
    0% {
      top: -100%;
    }
    60% {
      top: 0%;
    }
    100% {
      top: -50%;
    }
  }

  @keyframes pushUp {
    0% {
      top: -50%;
    }
    60% {
      top: 0%;
    }
    100% {
      top: -100%;
    }
  }

  .right {
    margin-left: auto;
    display: flex;
    align-items: center;
  }

  .hiddenAnimate {
    animation: 0.3s ease-in show forwards;
    transition: 0.3s all;
    @media screen and (max-width: 1000px) {
      animation: 0.3s ease-in hidden forwards;
    }
  }

  @keyframes hidden {
    from {
      transform: translateX(0px);
      width: 79px;
      opacity: 1;
    }
    to {
      transform: translateX(-6px);
      opacity: 0;
      width: 0px;
    }
  }

  @keyframes show {
    from {
      opacity: 0;
      width: 0;
      transform: translateX(-6px);
    }
    to {
      transform: translateX(0px);
      width: 79px;
      opacity: 1;
    }
  }

  :deep(.n-image) {
    display: block;

    img {
      display: block;
    }
  }
}
</style>
