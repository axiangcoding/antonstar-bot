import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import AutoImport from "unplugin-auto-import/vite";
import Components from "unplugin-vue-components/vite";
import {NaiveUiResolver} from "unplugin-vue-components/resolvers";
import {resolve} from "path";


// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        vue(),
        AutoImport({
            imports: [
                'vue',
                {
                    'naive-ui': [
                        'useDialog',
                        'useMessage',
                        'useNotification',
                        'useLoadingBar'
                    ]
                }
            ]
        }),
        Components({
            resolvers: [NaiveUiResolver()]
        })
    ],
    base: "./",
    resolve: {
        alias: {
            "@": resolve(__dirname, "src")
        }
    },
    server:{
        proxy:{
            '/api':{
                target:'http://localhost:8888'
            }
        }
    }
})
