import {createStore} from "vuex";

export default createStore({
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
