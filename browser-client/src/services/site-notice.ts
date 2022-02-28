import http from "@/services/request";

export function getLastSiteNotice() {
    return http.get('/v1/site/notice/last', {})
}
