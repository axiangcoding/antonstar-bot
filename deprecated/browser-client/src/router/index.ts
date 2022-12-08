import {createRouter, createWebHashHistory, RouteRecordRaw} from "vue-router";

const routes: RouteRecordRaw[] = [
    {
        path: "/",
        component: () => import("@/layout/Index.vue"),
        children: []
    }
];

export default createRouter({
    history: createWebHashHistory(),
    routes,
});