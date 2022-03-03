import http from "@/services/request";

export function getLastSiteNotice() {
    return http.get('/v1/site/notice/last', {})
}

export interface NoticeForm {
    content: string,
    title: string
}

export function postSiteNotice(auth: string, form: NoticeForm) {
    return http.post('/v1/site/notice/', {
        'content': form.content,
        'title': form.title
    }, {
        headers: {
            'Authorization': auth
        }
    })
}
