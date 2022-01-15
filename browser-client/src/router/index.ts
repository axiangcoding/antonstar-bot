import {createRouter, createWebHashHistory, RouteRecordRaw} from "vue-router";
import http from "../services/request";
import {store} from "../store";

declare module 'vue-router' {
    interface RouteMeta {
        // 是可选的
        title: string
    }
}

const routes: Array<RouteRecordRaw> = [
    // {
    //     path: "/",
    //     name: "Home",
    //     component: () => import("../views/home/Index.vue"),
    //     meta: {
    //         title: '首页'
    //     }
    // },
    // {
    //     path: "/record",
    //     name: "Record",
    //     component: () => import("../views/record/Index.vue"),
    //     meta: {
    //         title: '战绩查询'
    //     }
    // },
    // {
    //     path: "/realtime",
    //     name: "RealtimeGame",
    //     component: () => import("../views/realtime/Index.vue"),
    //     meta: {
    //         title: '战绩查询'
    //     }
    // },
    // {
    //     path: "/about",
    //     name: "About",
    //     component: () => import("../views/about/Index.vue"),
    //     meta: {
    //         title: '关于我们'
    //     }
    // },
    // {
    //     path: '/:pathMatch(.*)*',
    //     name: 'NotFound',
    //     component: () => import("../views/common/NotFound.vue"),
    //     meta: {
    //         title: '页面未找到'
    //     }
    // }
];

const router = createRouter({
    history: createWebHashHistory(),
    routes
});

router.afterEach((to, from, failure) => {
    if (!failure) {
        let concat = " - "
        let prefix = document.title.split(concat)[0]
        document.title = prefix + concat + to.meta.title
        http.post('/v1/visits/visit', {
            client_id: store.state.clientId,
            page: window.location.href,
            user_id: store.state.userId
        }).then()
    }
})

export default router;
