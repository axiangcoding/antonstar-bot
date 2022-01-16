import {createApp} from 'vue'
import App from './App.vue'
import {store} from './store'
import router from './router'
import {
    // create naive ui
    create,
    // component
    NButton,
    NImage,
    NGrid,
    NGridItem
} from 'naive-ui'

const naive = create({
    components: [NButton, NImage, NGrid, NGridItem],
})

const app = createApp(App)

app.use(store)
    .use(router)
    .use(naive)
    .mount('#app')
