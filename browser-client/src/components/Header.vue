<template>
	<n-layout-header position="absolute">
		<n-grid item-responsive>
			<!-- 左侧栅栏 -->
			<n-grid-item span="0 768:1 1200:2 1920:4"> </n-grid-item>
			<n-grid-item span="24 768:22 1200:20 1920:16" class="header-content">
				<div class="logo">
					<n-image height="60" :src="LOGO_URL" preview-disabled />
				</div>
			</n-grid-item>
			<!-- 右侧栅栏 -->
			<n-grid-item span="0 768:1 1200:2 1920:4"> </n-grid-item>
		</n-grid>
	</n-layout-header>
</template>

<script lang="ts">
import logo from '@/assets/logo/logo.png'
import { NLayoutHeader } from 'naive-ui'
import { defineComponent, ref, onMounted, onBeforeUnmount } from 'vue'

export default defineComponent({
	components: {
		NLayoutHeader,
	},
	setup() {
		// const internalInstance = getCurrentInstance()
		// const globalProperties =
		// 	internalInstance?.appContext.config.globalProperties
		// const anime = globalProperties?.anime
    let header: any = null
    let scrollBox: any = null
		let scrollFunc: any = null

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
						if(scrollTop == 0) header.classList.add('immersive')
            else header.classList.remove('immersive')
					})
				)
		})
    onBeforeUnmount(() => {
      scrollBox?.removeEventListener('scroll', scrollFunc)
    })
		const LOGO_URL = ref(logo)
		return {
			LOGO_URL
		}
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
