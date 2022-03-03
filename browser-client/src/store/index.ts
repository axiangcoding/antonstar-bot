// store.ts
import {InjectionKey} from 'vue'
import createPersistence from 'vuex-persistedstate'

import {createStore, useStore as baseUseStore, Store} from 'vuex'
import {userInfo} from "@/services/user";

export interface State {
    themes: string
    clientId: string
    login: boolean
    auth: string
    userId: number,
    userInfo: object,
    loading: boolean
}

export const key: InjectionKey<Store<State>> = Symbol()

export const store = createStore<State>({
    plugins: [
        createPersistence({
            key: 'vuex',
            storage: localStorage
        }),
    ],
    state: {
        themes: 'anton_star',
        clientId: '',
        login: false,
        auth: '',
        userId: 0,
        userInfo: {},
        loading: false,
    },
    mutations: {
        setClientId(state, clientId) {
            state.clientId = clientId
        },
        setLogin(state, login) {
            state.login = login
        },
        setAuth(state, auth) {
            state.auth = auth
        },
        setUserId(state, userId) {
            state.userId = userId
        },
        setThemes(state, themes) {
            state.themes = themes
        },
        setLoading(state, val) {
            state.loading = val
        },
        setUserInfo(state, val) {
            state.userInfo = val
        },
        logout(state) {
            state.login = false
            state.auth = ''
            state.userInfo = {}
            state.userId = 0
        }
    },
    getters: {},
    actions: {},
    modules: {},
})

// 定义自己的 `useStore` 组合式函数
export function useStore() {
    return baseUseStore(key)
}

export function getStore() {
    return store
}
