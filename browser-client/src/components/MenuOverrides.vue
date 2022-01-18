<template>
	<div id="menuOverides">
		<n-space>
			<template v-for="item in menuOptions">
				<n-popselect
					v-if="item.children?.length"
					v-model:value="item.active"
					:options="item.children"
					trigger="click"
				>
					<n-button
						:class="{ 'menu-active': activeKey == item.routerName }"
						quaternary
						:key="item.key"
						:disabled="item.disabled"
						@click="routerClick(item.key, $event)"
						>{{ item.label }}
						<template #icon>
							<an-icon :vdom="item.icon"></an-icon>
						</template>
					</n-button>
				</n-popselect>
				<n-button
					v-else
					:class="{ 'menu-active': activeKey == item.routerName }"
					quaternary
					:key="item.key"
					:disabled="item.disabled"
					@click="routerClick(item.key, $event)"
				>
					{{ item.label }}
					<template #icon>
						<an-icon :vdom="item.icon"></an-icon>
					</template>
				</n-button>
			</template>
		</n-space>
		<div class="nav-active"></div>
	</div>
</template>

<script lang="ts" setup>
import { useRoute, useRouter } from 'vue-router'
import { ref, h, onMounted, nextTick } from 'vue'
import { NIcon } from 'naive-ui'
import { Award, CommentsRegular } from '@vicons/fa'

const route = useRoute()
const router = useRouter()
let activeKey: any = ref(null)

function renderIcon(icon: any) {
	return () => h(NIcon, null, { default: () => h(icon) })
}

onMounted(() => {
	activeKey.value = route.name
	nextTick(() => {
		const navActive = document.querySelector('#menuOverides .nav-active')
		const menuActive = document.querySelector('#menuOverides .menu-active')
		const width = menuActive?.clientWidth
		const left = menuActive?.offsetLeft
		navActive?.setAttribute('style', `width: ${width}px; left: ${left}px`)
	})
})

const menuOptions = [
	{
		key: 'record',
		icon: renderIcon(Award),
		label: '战绩查询',
		routerName: 'record',
		params: {
			nick: '',
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

const routerClick = (key: String, event: any) => {
	const { routerName: name, params } = menuOptions.find((o) => o.key == key)
	activeKey.value = name
	nextTick(() => {
		const navActive = document.querySelector('#menuOverides .nav-active')
		const menuActive = document.querySelector('#menuOverides .menu-active')
		const width = menuActive?.clientWidth
		const left = menuActive?.offsetLeft
		navActive?.setAttribute('style', `width: ${width}px; left: ${left}px`)
	})
	router.push({
		name,
		params,
	})
}
</script>

<style lang="scss" scoped>
#menuOverides {
	position: relative;
	margin-left: 10px;
	:deep(.n-button) {
		color: var(--header-text-color) !important;
		&:hover {
			color: var(--header-nav-hover-color) !important;
		}
		&.menu-active {
			color: var(--header-nav-active-color) !important;
			font-weight: bold !important;
		}
	}
	.nav-active {
		position: absolute;
		bottom: -2px;
		height: 2px;
		border-radius: 2px;
		background-color: var(--header-nav-active-color);
		transition: 0.3s all;
	}
}
</style>
