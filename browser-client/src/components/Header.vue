<template>
	<n-layout-header position="absolute">
		<div class="header-content">
			<router-link class="logo" to="/">
				<n-image height="46" :src="LOGO_URL" preview-disabled />
			</router-link>
			<n-menu
				v-model:value="activeKey"
				mode="horizontal"
				:options="menuOptions"
			/>
			<div class="right">
        <n-avatar
            round
            size="large"
            src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"
				/>
			</div>
		</div>
	</n-layout-header>
</template>

<script lang="ts" setup>
import {useRoute, RouterLink} from 'vue-router'
import {computed, defineComponent, h, onBeforeUnmount, onMounted, ref, resolveComponent} from 'vue'
import {NIcon, NMenu} from 'naive-ui'
import logo from '@/assets/logo/logo.png'
import { Award,CommentsRegular } from '@vicons/fa'

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
            {default: () => '战绩查询'}
        ),
  },
  {
    key: 'about',
    icon: renderIcon(CommentsRegular),
    label: () =>
        h(
            RouterLink,
            {
              to: {
                name: 'about',
              },
            },
            {default: () => '关于我们'}
        ),
  },
  {
    key: 'realtime',
    // FIXME: 渲染为router-link时，disabled元素错位
    disabled: true,
    label: () =>
        h(
            RouterLink,
            {
              to: {
                name: 'realtime',
              },
            },
            {default: () => '实时战绩'}
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

<style lang="scss" scoped>
.logo{
  cursor: pointer;
}

.n-layout-header {
	height: $headerHeight;
	z-index: 10;
	background-color: $headerBackColor;
	color: $headerColor;
	transition: 0.3s all;
	.logo {
		display: block;
	}
	:deep(.n-menu) {
		&.n-menu--horizontal {
			position: relative;
			top: 50%;
		}
		.n-menu-item--selected {
			.n-menu-item-content {
				.n-menu-item-content-header {
					color: $itemIconColorActive !important;
					font-weight: bold;
					font-size: 17px;
				}
				.n-menu-item-content__icon {
					color: $itemIconColorActive !important;
					font-weight: bold;
					font-size: 22px !important;
				}
			}
		}
		.n-menu-item-content {
			margin: auto;
			&:hover {
				.n-menu-item-content-header {
					transition: all 0.3s;
					color: $itemTextColorHover !important;
				}
				.n-menu-item-content__icon {
					color: $itemTextColorActive !important;
				}
			}
			.n-menu-item-content-header {
				color: $headerColor;
				font-size: 16px;
			}
			.n-menu-item-content__icon {
				color: $headerColor;
				font-size: 21px !important;
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
		display: flex;
		align-items: center;
		margin: 0 auto;
	}

	.right {
		margin-left: auto;
    display: flex;
    align-items: center;
	}

	:deep(.n-image) {
		display: block;

		img {
			display: block;
		}
	}
}
</style>
