<template>
	<n-layout-footer></n-layout-footer>
</template>

<script lang="ts">
  import { NLayoutFooter } from 'naive-ui';
  import { defineComponent } from 'vue'

  export default defineComponent({
	components: {
		NLayoutFooter,
	},
})
</script>

<script setup lang="ts"></script>

<style lang="scss" scoped>
.n-layout-footer {
		height: $footerHeight;
	}
</style>
