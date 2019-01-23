
import api from '../../api/user'

export default {
    namespaced: true,
    state: {
        user: null
    },
    actions: {
        login({commit}, credentials) {
            api.login(credentials , (user) => {
                commit("login", user)
            })
        }
    },
    mutations: {
        login (state, user) {
            state.user = user
        }
    }
}