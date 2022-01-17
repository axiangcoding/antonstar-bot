<template>
	<div>
		<n-menu
			v-model:value="activeKey"
			mode="horizontal"
			:options="menuOptions"
		/>
	</div>
</template>

<script lang="ts" setup>
import { useRoute, RouterLink } from 'vue-router'
import { computed, h } from 'vue'
import { NIcon, NMenu } from 'naive-ui'
import { Award, CommentsRegular } from '@vicons/fa'

const route = useRoute()
const activeKey = computed(() => {
	return route.name as string
})

function renderIcon(icon: any) {
	return () => h(NIcon, null, { default: () => h(icon) })
}

const menuOptions = [
	{
		key: 'record',
		icon: renderIcon(Award),
		label: () =>
			h(
				RouterLink,
				{
					to: {
						name: 'record',
						params: {
							nick: 'zh-CN',
						},
					},
				},
				{ default: () => '战绩查询' }
			),
	},
	{
		key: 'about',
		icon: renderIcon(CommentsRegular),
		label: () =>
			h(
				RouterLink,
				{
					to: {
						name: 'record',
						params: {
							nick: 'zh-CN',
						},
					},
				},
				{ default: () => '关于我们' }
			),
	},
]
</script>

<style lang="scss" scoped></style>
