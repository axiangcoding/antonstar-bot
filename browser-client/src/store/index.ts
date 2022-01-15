// store.ts
import {InjectionKey} from 'vue'
import createPersistence from 'vuex-persistedstate'

import {createStore, useStore as baseUseStore, Store} from 'vuex'

export interface State {
    count: number,
    clientId: string,
    userId: number
}

export const key: InjectionKey<Store<State>> = Symbol()

export const store = createStore<State>({
    plugins: [
        createPersistence({
            key: 'vuex'
        })
    ],
    state: {
        count: 2,
        clientId: '',
        userId: 0
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
