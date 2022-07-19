import http from "@/services/request";


export interface CaptchaForm {
    captchaVal: string,
    captchaId: string
}


export function captcha(filename: string, reload: boolean) {
    return http.get('/v1/captcha/' + filename, {
        params: {
            'reload': reload
        }
    })
}



