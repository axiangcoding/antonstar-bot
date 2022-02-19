const username = '^[a-zA-Z0-9_]{5,16}$'
const password = '^.{8,16}$'
const email = '^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$'

export function getRegex(key: string) {
    switch (key) {
        case 'username':
            return username
        case 'password':
            return password
        case 'email':
            return email
        default:
            return null
    }
}
