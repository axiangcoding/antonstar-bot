import axios from 'axios';

// 创建axios实例
const http = axios.create({
    // 服务接口请求
    baseURL: import.meta.env.VITE_APP_REQUEST_BASE_URL,
    // 超时设置
    timeout: import.meta.env.VITE_APP_REQUEST_TIMEOUT,
    headers: {}
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
        console.log(error.code);
        if (error.code === 'ECONNABORTED') {

        } else if (error.response.status == 400) {

        }
        return Promise.reject(error)
    }
)

export default http;
