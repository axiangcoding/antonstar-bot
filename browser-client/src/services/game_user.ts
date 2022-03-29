import http, {Pagination} from "@/services/request";

export function getGameUsers(auth: string, form: Pagination) {
    return http.get('/v1/game_users/', {
        params: {
            'page_num': form.pageNum,
            'page_size': form.pageSize,
            'filter': form.filter
        }
    })
}