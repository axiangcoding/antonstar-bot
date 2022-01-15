<template>
	<n-layout>
		<n-layout-header>
			<Header></Header>
		</n-layout-header>
		<n-layout-content>
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
	created() {
		console.log(this.$route)
	}
})
</script>
