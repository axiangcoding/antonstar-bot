<template>
	<n-layout-header position="absolute" v-injectThemes="'Layout'">
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
	height: var(--header-height);
	z-index: 10;
	display: flex;
	align-items: center;
	transition: 0.3s all;
	.logo {
		display: block;
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
		width: 100%;
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
