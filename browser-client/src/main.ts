import { createApp } from 'vue'
import App from './App.vue'
import { store } from './store'
import router from './router'
import anime from 'animejs'
import naive from 'naive-ui'
import injectDirectives from './directives/index.ts'

// 不放入全局scss配置中，全局scss配置导致页面style标签将其多次引入。
import '@/scss/global.scss'

const app = createApp(App)

injectDirectives(app)

// 全局变量
app.config.globalProperties.anime = anime

app.use(store).use(router).use(naive).mount('#app')
