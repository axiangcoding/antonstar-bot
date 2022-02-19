import http from "@/services/request";
import {CaptchaForm} from "@/services/captcha";

export function userValueExist(key: string, val: string | number) {
    return http.post('/v1/user/value/exist', {
        'key': key,
        'value': val
    })
}

export interface RegForm {
    avatarUrl?: string,
    email: string,
    invitedCode: string,
    password: string,
    phone?: string,
    username: string
}

export function userRegister(form: RegForm, captcha: CaptchaForm) {
    return http.post('/v1/user/register', {
        "avatar_url": form.avatarUrl,
        "email": form.email,
        "invited_code": form.invitedCode,
        "password": form.password,
        "phone": form.phone,
        "username": form.username
    }, {
        params: {
            'captcha_id': captcha.captchaId,
            'captcha_val': captcha.captchaVal
        }
    })
}


