import {createApp} from 'vue'
import App from './App.vue'
import '@/global.css'
import router from "@/router";
import store from "@/store";


import 'vfonts/Lato.css'


const app = createApp(App);
app.use(router).use(store)
app.mount('#app')
