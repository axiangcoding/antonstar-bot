/// <reference types="vite/client" />

interface ImportMetaEnv {
    readonly VITE_APP_REQUEST_BASE_URL: string
    readonly VITE_APP_REQUEST_TIMEOUT: number
    // 更多环境变量...
}

interface ImportMeta {
    readonly env: ImportMetaEnv
}
