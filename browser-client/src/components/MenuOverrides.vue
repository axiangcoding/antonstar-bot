<template>
	<div id="menuOverides">
		<!-- <n-space> -->
		<template v-for="item in menuOptions">
			<n-dropdown
				v-if="item.options?.length"
				@select="handleSelect"
				:options="item.options"
				trigger="click"
			>
				<n-button
					:class="{ 'menu-active': activeKey == item.routerName }"
					quaternary
					:key="item.key"
					:disabled="item.disabled"
					@click="routerClick(item.key)"
					>{{ item.label }}
					<template #icon v-if="item.icon">
						<an-icon :vdom="item.icon"></an-icon>
					</template>
				</n-button>
			</n-dropdown>
			<n-button
				v-else
				:class="{ 'menu-active': activeKey == item.routerName }"
				quaternary
				:key="item.key"
				:disabled="item.disabled"
				@click="routerClick(item.key)"
			>
				{{ item.label }}
				<template #icon v-if="item.icon">
					<an-icon :vdom="item.icon"></an-icon>
				</template>
			</n-button>
		</template>
		<!-- </n-space> -->
		<div class="nav-active" ref="navActive"></div>
	</div>
</template>

<script lang="ts" setup>
import { useRoute } from 'vue-router'
import { h, ref, onMounted, nextTick, watch, computed } from 'vue'
import { NIcon } from 'naive-ui'
import { Award, CommentsRegular } from '@vicons/fa'

const route = useRoute()
const navActive = ref(null)
let activeKey: any = computed(() => {
	return route.name as string
})

function navAnimation() {
	const menuActive = document.querySelector('#menuOverides .menu-active')
	const width = menuActive?.clientWidth
	const left = menuActive?.offsetLeft
	navActive.value?.setAttribute('style', `width: ${width}px; left: ${left}px`)
}

onMounted(() => {
	navAnimation()
})

watch(
	() => route.name,
	() => {
		nextTick(() => {
			navAnimation()
		})
	}
)
</script>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
	data() {
		return {
			options: [],
			optionsTemp: []
		}
	},
	computed: {
		menuOptions() {
			const options = [
				{
					key: 'record',
					icon: this.renderIcon(Award),
					label: '战绩查询',
					routerName: 'record',
					params: {
						nick: '',
					},
					active: null, // 选中
					options: [], // 子项
				},
				{
					key: 'about',
					icon: this.renderIcon(CommentsRegular),
					label: '关于我们',
					routerName: 'about',
				},
				{
					key: 'realtime',
					// icon: renderIcon(CommentsRegular),
					label: '实时数据',
					disabled: true,
					routerName: 'realtime',
				},
			]
			const optionsList = JSON.parse(JSON.stringify(this.options))
			const result = options.filter(
				(o) => !optionsList.find((p) => p.key == o.key)
			)
			if (optionsList.length) {
				result.push({
					key: 'others',
					label: '更多',
					options: optionsList,
				})
			}
			return result
		},
	},
	mounted() {
		const menu = document.querySelector('#menuOverides')
		const menuItems = document.querySelectorAll('#menuOverides .n-button')
		const widthLists: any[] = []
		const optionsTemp: any[] = []
		menuItems.forEach((o, index) => {
			widthLists.push(o.clientWidth)
			optionsTemp.push(this.menuOptions[index])
		})
		this.optionsTemp = optionsTemp
		const resizeObserver = new ResizeObserver((resizeObj) => {
			let menuWidth = -12
			const optionsMap: any[] = []
			let flag = true
			widthLists.forEach((o, index) => {
				if (
					menuWidth + 12 + o <= resizeObj[0].target.clientWidth - 56 &&
					flag
				) {
					menuWidth += 12 + o
				} else {
					flag = false
					optionsMap.push({
						label: optionsTemp[index].label,
						key: optionsTemp[index].key,
						disabled: optionsTemp[index].disabled,
					})
				}
			})
			this.options = optionsMap
			this.navAnimation()
		})
		resizeObserver.observe(menu)
	},
	methods: {
		renderIcon(icon: any) {
			return () => h(NIcon, null, { default: () => h(icon) })
		},
		routerClick(key: String) {
			if (key == 'others') {
				return
			}
			const { routerName: name, params } = this.optionsTemp.find(
				(o) => o.key == key
			)
			this.$router.push({
				name,
				params,
			})
		},
		handleSelect(key: any) {
			this.routerClick(key)
		},
	},
})
</script>

<style lang="scss" scoped>
#menuOverides {
	position: relative;
	margin-left: 10px;
	overflow: hidden;
	display: flex;
	flex: 1;
	:deep(.n-button) {
		color: var(--header-text-color) !important;
		&:not(:first-child) {
			margin-left: 12px;
		}
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
		bottom: 0px;
		height: 2px;
		border-radius: 2px;
		background-color: var(--header-nav-active-color);
		transition: 0.3s all;
	}
}
</style>
