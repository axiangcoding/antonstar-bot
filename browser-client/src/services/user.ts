import http from "@/services/request";

export function userValueExist(key: string, val: string | number) {
    return http.post('/v1/user/value/exist', {
        'key': key,
        'value': val
    })
}

export function captcha(filename: string, reload: boolean) {
    return http.get('/v1/captcha/' + filename, {
        params: {
            'reload': reload
        }
    })
}
