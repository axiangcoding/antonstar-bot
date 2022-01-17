import moment from "moment-timezone";


export function parseLocalTime(str: string) {
    return moment(str).tz('Asia/Shanghai').format('YYYY-MM-DD HH:mm:ss')
}
