import {createStore} from "vuex";

export default createStore({
    state: {
        count: 1
    },
    getters: {
        getTest(state) {
            return state.count
        }
    },
    mutations: {
        addTest(state) {
            ++state.count
        }
    },
    actions: {},
    modules: {}
});
