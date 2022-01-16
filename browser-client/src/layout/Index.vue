<template>
	<n-layout embedded class="layout">
		<Header></Header>
		<n-layout-content content-style="padding: 10px;" class="content">
			<router-view></router-view>
		</n-layout-content>
		<Footer></Footer>
	</n-layout>
</template>

<script lang="ts">
import { NLayout, NLayoutContent } from 'naive-ui'
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
	min-height: calc(100vh - $footerHeight);
	max-width: 1200px;
	margin: 0 auto;
}

.layout {
  background-image: linear-gradient(135deg, #FFF5C3 10%, #9452A5 100%);
}
</style>
