import {createStore} from "vuex";
import createPersistence from 'vuex-persistedstate'

export default createStore({
    plugins: [
        createPersistence({
            key: 'vuex'
        })
    ],
    state: {
        count: 2
    },
    mutations: {
        increment(state) {
            state.count++
        }
    },
    getters: {},
    actions: {},
    modules: {},
});
