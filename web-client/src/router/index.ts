import {createRouter, createWebHashHistory, RouteRecordRaw} from "vue-router";
import service from "../util/request";
import store from "../store";

declare module 'vue-router' {
    interface RouteMeta {
        // 是可选的
        title: string
    }
}

const routes: Array<RouteRecordRaw> = [
    {
        path: "/",
        name: "Home",
        component: () => import("../views/home/Index.vue"),
        meta: {
            title: '首页'
        }
    },
    {
        path: "/about",
        name: "About",
        component: () => import("../views/About.vue"),
        meta: {
            title: '关于我们'
        }
    }
];

const router = createRouter({
    history: createWebHashHistory(),
    routes
});

router.afterEach((to, from, failure) => {
    if (!failure) {
        document.title = to.meta.title
        service.post('/v1/visits/visit', {
            client_id: store.state.clientId,
            page: window.location.href,
            user_id: store.state.userId
        }).then()
    }
})

export default router;
