<template>
	<div id="menuOverides">
		<template v-for="item in menuOptions">
			<n-popselect
				v-if="item.children?.length"
				v-model:value="item.active"
				:options="item.children"
				trigger="click"
			>
				<n-button
					:class="{ 'menu-ative': activeKey == item.key }"
					quaternary
					:key="item.key"
					@click="routerClick(item.key)"
					>{{ item.label }}
					<template #icon> </template>
				</n-button>
			</n-popselect>
			<n-button
				v-else
				:class="{ 'menu-ative': activeKey == item.key }"
				quaternary
				:key="item.key"
				@click="routerClick(item.key)"
			>
				{{ item.label }}
				<template #icon v-html="item.icon"></template>
			</n-button>
		</template>
		<div></div>
	</div>
</template>

<script lang="ts" setup>
import {
	useRoute,
	useRouter,
	RouteRecordName,
	RouteParamsRaw,
} from 'vue-router'
import { computed, h, onMounted } from 'vue'
import { NIcon } from 'naive-ui'
import { Award, CommentsRegular } from '@vicons/fa'

const route = useRoute()
const router = useRouter()
const activeKey = computed(() => {
	return route.name as string
})

function renderIcon(icon: any) {
	return () => h(NIcon, null, { default: () => h(icon) })
}

onMounted(() => {
	const NMenu = document.querySelector('#menuOverides .n-menu')
	// NMenu?.addEventListener('')
})

const menuOptions = [
	{
		key: 'record',
		icon: renderIcon(Award),
		label: '战绩查询',
		routerName: 'record',
		params: {
			nick: 'zh-CN',
		},
		active: null, // 选中
		children: [], // 子项
	},
	{
		key: 'about',
		icon: renderIcon(CommentsRegular),
		label: '关于我们',
		routerName: 'about',
	},
	{
		key: 'realtime',
		// icon: renderIcon(CommentsRegular),
		label: '实时战绩',
		disabled: true,
		routerName: 'realtime',
	},
]

const routerClick = (key: String) => {
	const { routerName: name, params } = menuOptions.find((o) => o.key == key)
	router.push({
		name,
		params,
	})
}
</script>

<style lang="scss" scoped>
#menuOverides {
	:deep(.n-button) {
		color: var(--header-text-color) !important;
		&:hover {
			color: var(--header-nav-hover-color) !important;
		}
		&.menu-ative {
			color: var(--header-nav-active-color) !important;
			font-weight: bold !important;
		}
	}
}
</style>
