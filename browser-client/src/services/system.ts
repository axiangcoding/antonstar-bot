import http from "@/services/request";

export function getSystemInfo() {
    return http.get('/v1/system/info/')
}
