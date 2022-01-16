/// <reference types="vite/client" />

declare module '*.vue' {
  import { DefineComponent } from 'vue'
  // eslint-disable-next-line @typescript-eslint/no-explicit-any, @typescript-eslint/ban-types
  const component: DefineComponent<{}, {}, any>
  export default component
}

interface ImportMetaEnv {
  readonly VITE_APP_REQUEST_BASE_URL: string
  readonly VITE_APP_REQUEST_TIMEOUT: undefined
  // 更多环境变量...
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
