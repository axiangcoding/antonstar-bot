<template>
	<div class="loading" id="loading" v-show="showLoading">
		<n-icon color="#f2f2f2" size="250">
			<Instalod />
		</n-icon>
	</div>
</template>

<script lang="ts" setup>
import { Instalod } from '@vicons/fa'
import anime from 'animejs'
import { watch, ref, onMounted } from 'vue'

const props = defineProps({
	loading: {
		type: Boolean,
		default: false,
	},
})

const showLoading = ref(false)
let timer: any = ref(null)

onMounted(() => {
	anime({
		targets: '#loading .n-icon',
		rotate: '1turn',
		loop: true,
		easing: 'easeInOutSine',
	})
})

watch(
	() => props.loading,
	(newv) => {
		if (newv) {
      console.log('正在加载中...')
			if (timer.value) clearTimeout(timer.value)
			showLoading.value = true
			anime({
				targets: '#loading',
				left: '0%',
			})
		} else {
      console.log('加载结束')
			anime({
				targets: '#loading',
				left: '100%',
			})
			timer.value = setTimeout(() => {
				showLoading.value = false
			}, 1000)
		}
	}
)
</script>

<style lang="scss" scoped>
.loading {
	position: fixed;
	left: 100%;
	top: 0;
	width: 120vw;
	height: 100vh;
	background: rgba($color: #000000, $alpha: 0.5);
	z-index: 10000;
	.n-icon {
		position: absolute;
		left: calc(50vw - 125px);
		top: calc(50% - 125px);
	}
}
</style>
