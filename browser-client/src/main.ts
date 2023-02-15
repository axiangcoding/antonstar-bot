import {createApp} from 'vue'
import './style.css'
import App from './App.vue'

import router from "@/router";
import store from "@/store";


import 'vfonts/Lato.css'


const app = createApp(App);
app.use(router).use(store)
app.mount('#app')
