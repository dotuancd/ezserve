export default {
    state: {
        user: {}
    },
    mutations: {
        login (state, user) {
            state.user = user
        }
    }
}