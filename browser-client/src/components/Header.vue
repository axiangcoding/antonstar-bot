<template>
	<n-layout-header position="absolute">
		<div class="header-content">
			<router-link class="logo" to="/">
				<n-image height="46" :src="LOGO_URL" preview-disabled />
			</router-link>
      <menu-overides></menu-overides>
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
import { onBeforeUnmount, onMounted, ref } from 'vue'
import logo from '@/assets/logo/logo.png'

let header: any = null
let scrollBox: any = null
let scrollFunc: any = null

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
import MenuOverides from './MenuOverides.vue'

export default defineComponent({
	components: {
		NLayoutHeader,
    MenuOverides
	},
})
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
