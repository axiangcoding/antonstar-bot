import { store } from '@/store/index.ts'

function changeThemes(el: any, binding: any, vnode: any) {
	const themeOverrides = {}
	const style = JSON.parse(JSON.stringify(store.getters.getThemesOverides))[
		binding.value
	]
	for (let i in style) {
		const item = i.replace(/([A-Z])/g, '-$1').toLowerCase()
		const styleAttr = el.getAttribute('style')
		el.setAttribute('style', styleAttr + `--${item}:${style[i]};`)
	}
}

const injectThemes: Function = (app: any) => {
  app.directive('injectThemes', {
    mounted(el: any, binding: any, vnode: any) {
      changeThemes(el, binding, vnode)
    },
		updated(el: any, binding: any, vnode: any) {
			changeThemes(el, binding, vnode)
		},
	})
}

export default injectThemes
