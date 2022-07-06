import axios from 'axios';
import {ElLoading, ElMessage} from 'element-plus';

// 创建axios实例
const service = axios.create({
    // 服务接口请求
    baseURL: import.meta.env.VITE_APP_REQUEST_BASE_URL,
    // 超时设置
    timeout: import.meta.env.VITE_APP_REQUEST_TIMEOUT,
    headers: {}
})

let loading: any;
//正在请求的数量
let requestCount: number = 0
//显示loading
const showLoading = () => {
    if (requestCount === 0 && !loading) {
        loading = ElLoading.service({
            text: "拼命加载中，请稍后...",
            background: 'rgba(0, 0, 0, 0.7)',
            spinner: 'el-icon-loading',
        })
    }
    requestCount++;
}
//隐藏loading
const hideLoading = () => {
    requestCount--
    if (requestCount == 0) {
        loading.close()
    }
}

// 请求拦截
service.interceptors.request.use(config => {
    return config
}, error => {

    return Promise.reject(error)
})

// 响应拦截器
service.interceptors.response.use((res: any) => {
        return Promise.resolve(res.data)
    },
    error => {
        console.log(error.code);
        if (error.code === 'ECONNABORTED') {
            ElMessage.error('连接到服务器超时，请稍后重试！')
        } else if (error.response.status == 400) {
            ElMessage({
                message: '请输入正确的参数！',
                type: 'warning',
            })
        }
        return Promise.reject(error)
    }
)

export default service;
