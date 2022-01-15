<template>
	<n-layout>
		<n-layout-header position="absolute">
			<Header></Header>
		</n-layout-header>
		<n-layout-content class="content">
			<router-view></router-view>
		</n-layout-content>
		<n-layout-footer>
			<Footer></Footer>
		</n-layout-footer>
	</n-layout>
</template>

<script lang="ts">
import { NLayout, NLayoutHeader, NLayoutContent, NLayoutFooter } from 'naive-ui'
import Header from '@/components/Header.vue'
import Footer from '@/components/Footer.vue'
import { defineComponent } from 'vue'

export default defineComponent({
	components: {
		NLayout,
		NLayoutHeader,
		NLayoutContent,
		NLayoutFooter,
		Header,
		Footer,
	},
})
</script>

<style lang="scss" scoped>
	::v-deep .n-layout-header {
		height: $headerHeight;
		z-index: 10;
	}
	.content {
		padding-top: $headerHeight;
		min-height: calc(100vh - $footerHeight);
	}
	.n-layout-footer {
		height: $footerHeight;
	}
</style>
