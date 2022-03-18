import http from "@/services/request";


export interface BugReportForm {
    type: string,
    title: string,
    content: string,
    anonymous: boolean
}


export function postBugReport(auth: string, form: BugReportForm) {
    return http.post('/v1/bug_report/', {
        'content': form.content,
        'title': form.title,
        'type': form.type,
        'anonymous': form.anonymous
    }, {
        headers: {
            'Authorization': auth
        }
    })
}
