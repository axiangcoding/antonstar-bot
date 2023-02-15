import {createRouter, createWebHashHistory} from "vue-router";


const routes = [
    {
        path: '/',
        component: () => import('@/components/home/Home.vue')
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
});

export default router