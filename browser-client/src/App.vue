<script setup lang="ts">
import { useStore } from 'vuex'
import { computed } from 'vue'
import Loading from './components/Loading.vue';

const store = useStore()

const loading = computed(() => store.state.loading)
const themeOverrides = computed(() => store.getters.getThemesOverides)
</script>

<template>
	<!-- 调整主题变量 -->
	<n-config-provider :theme-overrides="themeOverrides">
		<n-message-provider>
			<n-el tag="div">
				<router-view />
			</n-el>
		</n-message-provider>
		<n-back-top :right="30" />
		<Loading :loading="loading"/>
	</n-config-provider>
</template>

<style lang="scss">
html,
body {
	margin: 0;
	overflow: hidden;
}

#app {
	font-family: Avenir, Helvetica, Arial, sans-serif;
	-webkit-font-smoothing: antialiased;
	-moz-osx-font-smoothing: grayscale;
	text-align: center;
	font-size: 16px;
	height: 100vh;
	overflow-x: hidden;
}
</style>
