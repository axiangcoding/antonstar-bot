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

export interface LoginForm {
    password: string,
    username: string
}

export function userLogin(auth: string, form: LoginForm, captcha: CaptchaForm) {
    return http.post('/v1/user/login', {
        "password": form.password,
        "username": form.username
    }, {
        params: {
            'captcha_id': captcha.captchaId,
            'captcha_val': captcha.captchaVal
        },
        headers: {
            'Authorization': auth
        }
    })
}

export function userLogout(auth: string,) {
    return http.post('/v1/user/logout', {}, {
        headers: {
            'Authorization': auth
        }
    })
}

export function userInfo(auth: string, user_id?: string) {
    return http.post('/v1/user/info', {}, {
        headers: {
            'Authorization': auth
        },
        params: {
            'user_id': user_id
        }
    })
}
