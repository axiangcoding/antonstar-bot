<template>
  <div id="menuOverides">
    <n-menu
        :value="activeKey"
        mode="horizontal"
        :options="menuOptions"
        :dropdown-props="{
				show: expand
			}"
    />
    <div class="nav-active" ref="navActive"></div>
  </div>
</template>


<script lang="ts" setup>
import {useRoute, RouterLink} from 'vue-router'
import {h, ref, onMounted, nextTick, watch, computed} from 'vue'
import {NIcon} from 'naive-ui'
import {Award, CommentsRegular, Gamepad, Biohazard, EllipsisH, Toolbox, FortAwesome} from '@vicons/fa'
import {renderIcon} from "@/util/naive";
import {useStore} from "vuex";

const route = useRoute()
const navActive = ref(null)
let activeKey: any = computed(() => {
  return route.name as string
})

let flag = ref(false)
let expand = ref(false)

const changeExpand = () => {
  expand.value = !expand.value
}

defineExpose({
  changeExpand,
})


const basicRoute = [
  {
    key: 'record',
    icon: renderIcon(Award),
    label: () =>
        h(RouterLink, {to: {name: 'record'}}, {default: () => '战绩查询'}),
  },
  {
    key: 'resource',
    icon: renderIcon(Toolbox),
    label: () =>
        h(RouterLink, {to: {name: 'resource'}}, {default: () => '游戏资源'}),
  },
  {
    key: 'rank',
    icon: renderIcon(Biohazard),
    label: () =>
        h(RouterLink, {to: {name: 'rank'}}, {default: () => '硬核狠人'}),
  },
  {
    key: 'about',
    icon: renderIcon(CommentsRegular),
    label: () =>
        h(RouterLink, {to: {name: 'about'}}, {default: () => '关于我们'}),
  },
]

const adminRoute = [{
  key: 'admin',
  icon: renderIcon(FortAwesome),
  label: () =>
      h(RouterLink, {to: {name: 'admin'}}, {default: () => '管理界面'})
}]

const options = ref(basicRoute)

const others = ref({
  key: 'others',
  label: '更多',
  icon: renderIcon(EllipsisH),
  children: [...options.value],
})

const menuOptions = computed(() => {
  const result = [...options.value]
  if (flag.value) return [others.value]
  return result
})

function navAnimation() {
  nextTick(() => {
    const menuActive = document.querySelector(
        '#menuOverides .n-menu-item--selected'
    )
    const width = menuActive?.clientWidth - 20
    const left = menuActive?.offsetLeft + 10
    navActive.value?.setAttribute('style', `width: ${width}px; left: ${left}px`)
  })
}

const store = useStore();
// 重设头部可选栏目
const resetHeader = () => {
  const roles = store.state.userInfo.roles;
  if (roles != undefined && roles.indexOf('admin') >= 0) {
    options.value = basicRoute.concat(adminRoute)
  } else {
    options.value = basicRoute
  }
}

onMounted(() => {
  navAnimation()
  if (window.innerWidth < 992) {
    flag.value = true
    navAnimation()
  } else {
    flag.value = false
    navAnimation()
  }
  window.onresize = (event) => {
    if (window.innerWidth < 768) {
      flag.value = true
      navAnimation()
    } else {
      flag.value = false
      navAnimation()
    }
  }
  resetHeader()
})

watch(
    () => store.state.userInfo,
    () => {
      resetHeader()
    }
)

watch(
    () => route.name,
    () => {
      nextTick(() => {
        navAnimation()
      })
    }
)
</script>

<style lang="scss" scoped>
#menuOverides {
  position: relative;
  margin-left: 10px;
  // overflow: hidden;
  // flex: 1;
  // text-align: left;
  :deep(.n-menu) {
    .n-menu-item.n-menu-item--selected {
      .n-menu-item-content {
        color: var(--header-nav-active-color);
        font-weight: bold;
      }
    }

    .n-menu-item-content {
      color: var(--header-text-color);
      display: flex;

      &:hover {
        color: var(--header-nav-hover-color);
      }

      .n-menu-item-content__icon {
        color: inherit;
      }

      .n-menu-item-content-header {
        color: inherit;
      }
    }
  }

  .nav-active {
    position: absolute;
    bottom: 2px;
    height: 2px;
    border-radius: 2px;
    background-color: var(--header-nav-active-color);
    transition: 0.3s all;
  }
}
</style>
