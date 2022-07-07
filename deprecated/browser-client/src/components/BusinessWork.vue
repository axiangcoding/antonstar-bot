<template>
	<n-modal :show="show" @update:show="(val: any) => $emit('update:show', val)">
		<n-card
			style="width: 600px"
			title="商业合作"
			:bordered="false"
			size="huge"
			role="dialog"
			aria-modal="true"
		>
			<template #header-extra></template>
			<div>
				邮箱：<a href="mailto:axiangcoding@gmail.com" target="_blank"
					>axiangcoding@gmail.com</a
				>
			</div>
			<!--<template #footer> 人也是要恰饭的嘛</template>-->
		</n-card>
	</n-modal>
</template>

<script lang="ts" setup>
import { toRefs } from 'vue'
const props = defineProps({
	show: {
		type: Boolean,
	},
})
const emits = defineEmits(['update:show'])
const refProps = toRefs(props)


</script>

<style scoped></style>
