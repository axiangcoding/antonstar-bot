"use strict";

import axios from "axios";

const axiosInstance = axios.create({
    // eslint-disable-next-line no-undef
    baseURL: config.VUE_APP_API_URL || process.env.VUE_APP_API_URL,
    headers: {
        // Authorization: 'Bearer {token}'
    }
});

axiosInstance.interceptors.request.use(config => {
    // Do something before request is sent
    return config;
}, error => {
    // Do something with request error
    return Promise.reject(error);
});

// axiosInstance.interceptors.response.use(response => {
//     return response
// }, error => {
//     return Promise.reject(error)
// })

export const HTTP = axiosInstance
