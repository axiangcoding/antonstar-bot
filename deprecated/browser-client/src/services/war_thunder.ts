import http from "@/services/request";

export function getWTUserInfoQueries(auth: string, nick: string) {
    return http.get('v1/war_thunder/userinfo/queries',
        {
            params: {
                "nickname": nick
            },
            headers: {
                'Authorization': auth
            }
        })
}

export function postWTUserInfoRefresh(auth: string, nick: string) {
    return http.post('v1/war_thunder/userinfo/refresh',
        {}, {
            params: {
                "nickname": nick
            },
            headers: {
                'Authorization': auth
            }
        })
}

export function getWTUserInfo(queryId: string) {
    return http.get('v1/war_thunder/userinfo',
        {
            params: {
                "query_id": queryId
            }
        })
}
