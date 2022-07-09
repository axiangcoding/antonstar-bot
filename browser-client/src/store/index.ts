import {createStore} from 'vuex'
import createPersistence from 'vuex-persistedstate'
// 创建一个新的 store 实例
const store = createStore({
    plugins: [
        createPersistence({
            key: 'vuex',
            storage: localStorage
        }),
    ],
    state() {
        return {
            count: 0
        }
    },
    mutations: {
        increment(state) {
            state.count++
        }
    }
})

export default store