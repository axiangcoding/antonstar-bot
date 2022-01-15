<template>
	<n-grid item-responsive>
		<!-- 左侧栅栏 -->
		<n-grid-item span="0 768:1 1200:2 1920:4"> </n-grid-item>
		<n-grid-item span="24 768:22 1200:20 1920:16" class="header-content">
			<div class="logo">
				<n-image
					height="60"
					:src="LOGO_URL"
					preview-disabled
				/>
			</div>
		</n-grid-item>
		<!-- 右侧栅栏 -->
		<n-grid-item span="0 768:1 1200:2 1920:4"> </n-grid-item>
	</n-grid>
</template>

<script setup lang="ts">
import logo from '@/assets/logo/logo.png'
import { ref } from 'vue'

const LOGO_URL = ref(logo)
</script>

<style lang="scss" scoped>

</style>
