// store.ts
import { InjectionKey } from 'vue'
import createPersistence from 'vuex-persistedstate'

import { createStore, useStore as baseUseStore, Store } from 'vuex'
import themes from './themes'

export interface State {
	theme: String
	count: number
	clientId: string
	userId: number
	loading: Boolean
}

export const key: InjectionKey<Store<State>> = Symbol()

export const store = createStore<State>({
	plugins: [
		createPersistence({
			key: 'vuex',
		}),
	],
	state: {
		theme: 'anton_star',
		count: 2,
		clientId: '',
		userId: 0,
		loading: false,
	},
	mutations: {
		increment(state) {
			state.count++
		},
		setUserId(state, userId) {
			state.userId = userId
		},
		setClientId(state, clientId) {
			state.clientId = clientId
		},
		setThemes(state, theme) {
			state.theme = theme
		},
		setLoading(state, val) {
			state.loading = val
		},
	},
	getters: {
		getThemesOverides(state) {
			return state.themes.themesOverides[state.theme]
		}
	},
	actions: {},
	modules: {
		themes,
	},
})

// 定义自己的 `useStore` 组合式函数
export function useStore() {
	return baseUseStore(key)
}
