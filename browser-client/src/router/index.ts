import {createRouter, createWebHashHistory} from "vue-router";


const routes = [
    {
        path: '/',
        component: () => import('@/page/home/Home.vue')
    },
    {
        path: '/query',
        component: () => import('@/page/query/Query.vue')
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
});

export default router