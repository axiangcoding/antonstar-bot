import axios from 'axios';
import {store} from "@/store";
import router from "@/router";
// 创建axios实例
const http = axios.create({
    // 服务接口请求
    baseURL: import.meta.env.VITE_APP_REQUEST_BASE_URL,
    // 超时设置
    timeout: import.meta.env.VITE_APP_REQUEST_TIMEOUT,
    withCredentials: false
})


// 请求拦截
http.interceptors.request.use(config => {
    return config
}, error => {
    return Promise.reject(error)
})

// 响应拦截器
http.interceptors.response.use((res: any) => {
        return Promise.resolve(res.data)
    },
    error => {
        if (error.code === 'ECONNABORTED') {
            window.$message.error('服务器暂时无法连接！请稍后重试')
        } else if (error.response.status == 400) {
            // window.$message.warning('请输入正确的参数')
        } else if (error.response.status == 401) {
            // window.$message.warning('功能仅限登录用户使用，请登录后再访问！')
            store.commit('logout')
            // router.push({name: 'login'})
        } else if (error.response.status == 403) {
            window.$message.warning('对不起，你无权访问这个资源')
            router.push({name: 'no_permission'})
        }
        // 这个会默认报错显示
        return Promise.reject(error)
    }
)

export default http;


export interface Pagination {
    pageNum: number,
    pageSize: number,
    filter: string
}