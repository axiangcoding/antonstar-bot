import {createStore} from "vuex";
import createPersistence from 'vuex-persistedstate'

export default createStore({
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
});
