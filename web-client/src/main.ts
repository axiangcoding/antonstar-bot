import {createApp} from 'vue'
import App from './App.vue'
import router from "./router"
import store from "./store"
import 'element-plus/dist/index.css'
import 'element-plus/theme-chalk/display.css'
import * as ElIconModules from '@element-plus/icons-vue'

const app = createApp(App);

// 统一注册Icon图标
for (const iconName in ElIconModules) {
    if (Reflect.has(ElIconModules, iconName)) {
        const item = ElIconModules[iconName]
        app.component(iconName, item)
    }
}
app.use(router)
    .use(store)
    .mount('#app')
