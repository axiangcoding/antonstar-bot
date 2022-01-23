import axios from 'axios';
import { store } from '@/store';
import { mapMutations } from 'vuex';

// 创建axios实例
const http = axios.create({
    // 服务接口请求
    baseURL: import.meta.env.VITE_APP_REQUEST_BASE_URL,
    // 超时设置
    timeout: import.meta.env.VITE_APP_REQUEST_TIMEOUT,
    headers: {}
})

const setLoading = mapMutations(['setLoading'])['setLoading']

// 请求拦截
http.interceptors.request.use(config => {
    // 调用接口时是否启用loading特效
    if (config.loading) {
        store.commit('setLoading', true)
    }
    return config
}, error => {
    store.commit('setLoading', false)
    return Promise.reject(error)
})

// 响应拦截器
http.interceptors.response.use((res: any) => {
        store.commit('setLoading', false)
        return Promise.resolve(res.data)
    },
    error => {
        store.commit('setLoading', false)
        if (error.code === 'ECONNABORTED') {
            window.$message.error('服务器暂时无法连接！')
        } else if (error.response?.status == 400) {
            window.$message.warning('请输入正确的参数')
        }
        return Promise.reject(error)
    }
)

export default http;
