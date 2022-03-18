import {createRouter, createWebHashHistory, RouteRecordRaw} from 'vue-router'
import http from '../services/request'
import {store} from '@/store'

declare module 'vue-router' {
    interface RouteMeta {
        // 是可选的
        title: string
    }
}

const routes: Array<RouteRecordRaw> = [{
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
            path: '/record',
            name: 'record',
            component: () => import('@/views/record/Index.vue'),
            meta: {
                title: '战绩查询',
            },
        },
        {
            path: '/player/:nick',
            name: 'recordPlayer',
            component: () => import('@/views/player/Index.vue'),
            meta: {
                title: '玩家战绩',
            },
        },
        {
            path: '/resource',
            name: 'resource',
            component: () => import('@/views/resource/Index.vue'),
            meta: {
                title: '游戏资源',
            },
        },
        {
            path: '/rank/',
            name: 'rank',
            component: () => import('@/views/rank/Index.vue'),
            meta: {
                title: '硬核狠人',
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
            path: '/user/login',
            name: 'login',
            component: () => import('@/views/user/Login.vue'),
            meta: {
                title: '用户登录',
            },
        },
        {
            path: '/user/register',
            name: 'register',
            component: () => import('@/views/user/Register.vue'),
            meta: {
                title: '用户注册',
            },
        },
        {
            path: 'admin',
            component: () => import('@/views/admin/Index.vue'),
            children: [
                {
                    path: '',
                    name: 'admin',
                    component: () => import('@/views/admin/dashboard/Index.vue'),
                    meta: {
                        title: '管理员 - 总览',
                    },
                },
                {
                    path: 'notice',
                    name: 'adminNotice',
                    component: () => import('@/views/admin/notice/Index.vue'),
                    meta: {
                        title: '管理员 - 通知管理',
                    },
                },
                {
                    path: 'global-config',
                    name: 'adminGlobalConfig',
                    component: () => import('@/views/admin/global_config/Index.vue'),
                    meta: {
                        title: '管理员 - 全局配置',
                    },
                }
            ]
        },
        {
            path: '/403',
            name: 'no_permission',
            component: () => import("@/views/common/NoPermission.vue"),
            meta: {
                title: '无权访问'
            }
        },
        {
            path: '/:pathMatch(.*)*',
            name: 'not_found',
            component: () => import("@/views/common/NotFound.vue"),
            meta: {
                title: '页面未找到'
            }
        }],
},]

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
