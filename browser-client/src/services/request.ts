import axios from 'axios';

// 创建axios实例
const http = axios.create({
    // 服务接口请求
    baseURL: import.meta.env.VITE_APP_REQUEST_BASE_URL,
    // 超时设置
    timeout: import.meta.env.VITE_APP_REQUEST_TIMEOUT,
    headers: {
        'Authorization': ''
    }
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
            window.$message.error('服务器暂时无法连接！')
        } else if (error.response.status == 400) {
            window.$message.warning('请输入正确的参数')
        }
        // return Promise.reject(error)  这个会默认报错显示
    }
)

export default http;


