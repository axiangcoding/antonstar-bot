import moment, {Moment} from "moment-timezone";


export function parseLocalTime(str: string) {
    return moment(str).tz('Asia/Shanghai').format('YYYY-MM-DD HH:mm:ss')
}

export function toISOTime(date: Date | Moment) {
    return moment(date).toISOString()
}

export function today() {
    return moment()
}
