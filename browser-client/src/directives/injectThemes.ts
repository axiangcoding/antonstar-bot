import themes from '@/themes/index.ts'
import { getStore } from '@/store/index.ts'

const injectThemes: Function = (app: any) => {
  app.directive('injectThemes', {
    mounted(el: any, binding: any, vnode: any) {
      const store = getStore()
			const themeOverrides = themes[store.state.themes]
      const style = themeOverrides[binding.value]
      for (let i in style) {
        const item = i.replace(/([A-Z])/g, '-$1').toLowerCase()
        const styleAttr = el.getAttribute('style')
        el.setAttribute('style', styleAttr + (`--${item}:${style[i]};`))
      }
    },
	})
}

export default injectThemes
