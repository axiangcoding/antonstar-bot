import { defineConfig } from 'vite'
import path from 'path'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [vue()],
	server: {
		open: false,
		proxy: {
			'/api': {
				target: 'http://localhost:8888',
				changeOrigin: true,
			},
		},
	},
	resolve: {
		alias: {
			'@': path.resolve(__dirname, 'src'),
		},
	},
	// 引入全局scss
	css: {
		preprocessorOptions: {
			scss: {
				additionalData: `@import "./src/scss/layout.scss";`,
			},
		},
	},
})
