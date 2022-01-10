import {createApp} from 'vue'
import App from './App.vue'
import router from "./router"
import store from "./store"
import 'element-plus/dist/index.css'
import 'element-plus/theme-chalk/display.css'


createApp(App)
    .use(router)
    .use(store)
    .mount('#app')
