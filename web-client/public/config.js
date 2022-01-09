// 这个文件主要是为了镜像中动态部署参数而设置的，为空的主要目的是为了devServer启动时不会报错
// eslint-disable-next-line no-unused-vars
const config = (() => {
    return {
        VUE_APP_API_URL: "/api"
    }
})()
