<template>
	<n-layout class="h100">
		<Header />
		<n-layout-content class="content">
			<router-view></router-view>
		</n-layout-content>
		<Footer />
	</n-layout>
</template>

<script lang="ts">
import { NLayout, NLayoutContent, NLayoutFooter } from 'naive-ui'
import Header from '@/components/Header.vue'
import Footer from '@/components/Footer.vue'
import { defineComponent } from 'vue'

export default defineComponent({
	components: {
		NLayout,
		NLayoutContent,
		Header,
		Footer,
	},
})
</script>

<style lang="scss" scoped>
.content {
	padding-top: $headerHeight;
	min-height: calc(100% - $footerHeight);
}
</style>
