import {createRouter, createWebHashHistory, RouteRecordRaw} from 'vue-router'
import http from '../services/request'
import {store} from '@/store'

declare module 'vue-router' {
    interface RouteMeta {
        // 是可选的
        title: string
    }
}

const routes: Array<RouteRecordRaw> = [
    {
        path: '/',
        component: () => import('@/layout/Index.vue'),
        children: [
            {
                path: '/',
                name: 'home',
                component: () => import('@/views/home/Index.vue'),
                meta: {
                    title: '首页',
                },
            },
            {
                path: '/record/:nick',
                name: 'record',
                component: () => import('@/views/record/Index.vue'),
                meta: {
                    title: '战绩查询',
                },
            },
            {
                path: '/realtime',
                name: 'realtime',
                component: () => import('@/views/realtime/Index.vue'),
                meta: {
                    title: '实时数据',
                },
            },
            {
                path: '/about',
                name: 'about',
                component: () => import('@/views/about/Index.vue'),
                meta: {
                    title: '关于我们',
                },
            },
            {
                path: '/:pathMatch(.*)*',
                name: 'not_found',
                component: () => import("@/views/common/NotFound.vue"),
                meta: {
                    title: '页面未找到'
                }
            }
        ],
    },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

router.afterEach((to, from, failure) => {
    if (!failure) {
        let concat = ' - '
        let prefix = document.title.split(concat)[0]
        document.title = prefix + concat + to.meta.title
        http.post('/v1/visits/visit', {
            client_id: store.state.clientId,
            page: window.location.href,
            user_id: store.state.userId,
        })
            .then()
    }
})

export default router
