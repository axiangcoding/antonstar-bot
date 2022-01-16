<template>
	<n-layout-header position="absolute">
		<div class="header-content">
      <n-space>
        <div class="logo">
          <n-image height="46" :src="LOGO_URL" preview-disabled />
        </div>
        <n-menu
          v-model:value="activeKey"
          mode="horizontal"
          :options="menuOptions"
        />
      </n-space>
      <div class="right">
        <n-avatar
          size="large"
          src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"
        />
      </div>
		</div>
	</n-layout-header>
</template>

<script lang="ts" setup>
import { useRoute, RouterLink } from 'vue-router'
import { computed, h, onBeforeUnmount, onMounted, ref } from 'vue'
import { NIcon, NMenu } from 'naive-ui'
import logo from '@/assets/logo/logo.png'
import { Award } from '@vicons/fa'

let header: any = null
let scrollBox: any = null
let scrollFunc: any = null

const route = useRoute()
const activeKey = computed(() => {
  return route.name as string
})

function renderIcon(icon: any) {
	return () => h(NIcon, null, { default: () => h(icon) })
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
							nick: 'zh-CN',
						},
					},
				},
				{ default: () => '战绩查询' }
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
							nick: 'zh-CN',
						},
					},
				},
				{ default: () => '关于我们' }
			),
	},
]
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
</script>

<script lang="ts">
import { NLayoutHeader } from 'naive-ui'
import { defineComponent } from 'vue'

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
	background-color: var(--header-back-color);
  color: var(--header-color);
	transition: 0.3s all;

  :deep(.n-menu) {
    &.n-menu--horizontal {
      transform: translateY(-50%);
      position: relative;
      top: 50%;
    }
    .n-menu-item-content {
      margin: auto;
      .n-menu-item-content-header {
        color: var(--header-color);
      }
      .n-menu-item-content__icon {
        color: var(--header-color);
      }
    }
  }

	&.immersive {
		padding-top: 10px;
		background: inherit !important;
	}

	:deep(.n-grid) {
		margin: auto 0;
	}

	.header-content {
    max-width: 1200px;
    min-width: 1100px;
    margin: 0 auto;
	}

  .right {
    margin-left: auto;
  }

	:deep(.n-image) {
		display: block;

		img {
			display: block;
		}
	}
}
</style>
