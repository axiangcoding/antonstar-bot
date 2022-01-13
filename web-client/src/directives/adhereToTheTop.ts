const func = {}

export function adhereToTheTop(app: any) {
  app.directive('adhereTop', {
		// 当被绑定的元素挂载到 DOM 中时……
    mounted(el: any, binding: any, vnode: any) {
			// 获取当前vueComponent的ID。作为存放各种监听事件的key
      const uid = '123'
			// 获取当前滚动的容器是什么。如果是document滚动。则可默认不传入parent参数
			const scrollParent =
				document.querySelector(binding.value.parent) || document
			el.style.position = 'relative'
      scrollParent.addEventListener('scroll', (func[uid] = (event: any) => {
        const scrollTop = event.target.scrollTop
        el.style.top = scrollTop + 'px'
      }))
		},
		// 节点取消绑定时 移除各项监听事件。
		unbind(el: any, binding: any, vnode: any) {
      const uid = '123'
      const scrollParent =
        document.querySelector(binding.value.parent) || document
      scrollParent.removeEventListener('scroll', func[uid])
		},
	})
}
