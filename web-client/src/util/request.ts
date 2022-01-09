import axios from 'axios';
import {ElMessage, ElLoading} from 'element-plus';

// 创建axios实例
const service = axios.create({
    // 服务接口请求
    // @ts-ignore
    baseURL: Config.serverUrl,
    // 超时设置
    // @ts-ignore
    timeout: Config.requestTimeout,
    headers: {'Content-Type': 'application/json;charset=utf-8'}
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
    showLoading()
    // 是否需要设置 token
    // config.headers['Authorization'] = 'Bearer ' + getToken() // 让每个请求携带自定义token 请根据实际情况自行修改
    // get请求映射params参数
    if (config.method === 'get' && config.params) {
        let url = config.url + '?';
        for (const propName of Object.keys(config.params)) {
            const value = config.params[propName];
            var part = encodeURIComponent(propName) + "=";
            if (value !== null && typeof (value) !== "undefined") {
                if (typeof value === 'object') {
                    for (const key of Object.keys(value)) {
                        let params = propName + '[' + key + ']';
                        var subPart = encodeURIComponent(params) + "=";
                        url += subPart + encodeURIComponent(value[key]) + "&";
                    }
                } else {
                    url += part + encodeURIComponent(value) + "&";
                }
            }
        }
        url = url.slice(0, -1);
        config.params = {};
        config.url = url;
    }
    return config
}, error => {
    console.log(error)
    Promise.reject(error)
})

// 响应拦截器
service.interceptors.response.use((res: any) => {
        hideLoading()
        // TODO: 校验业务错误代码
        return Promise.resolve(res.data)

    },
    error => {
        console.log('err' + error)
        hideLoading()
        let {message} = error;
        if (message == "Network Error") {
            message = "后端接口连接异常";
        } else if (message.includes("timeout")) {
            message = "系统接口请求超时";
        } else if (message.includes("Request failed with status code")) {
            message = "系统接口" + message.substr(message.length - 3) + "异常";
        }
        ElMessage.error({
            message: message,
            duration: 5 * 1000
        })
        return Promise.reject(error)
    }
)

export default service;
