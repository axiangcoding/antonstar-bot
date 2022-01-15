<template>
	<n-layout>
		<n-layout-header>
			<!-- 命名视图==>头部栏 默认为Header -->
			<router-view name="header">
				<Header></Header>
			</router-view>
		</n-layout-header>
		<n-layout-content>
			<router-view></router-view>
		</n-layout-content>
		<n-layout-footer>
			<!-- 命名视图==>底边栏 默认为Footer -->
			<router-view name="footer">
				<Footer></Footer>
			</router-view>
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

<style lang="scss" scoped></style>
