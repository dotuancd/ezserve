
import fileRequest from '../../api/files'
import logger from "vuex/dist/logger";

const state = {
    files: {
        items: []
    }
}

const actions = {
    loadPage({commit}, page) {
        console.log(arguments)
        fileRequest.files(files => {
            commit('setFiles', files)
        }, {page})
    }
}

const mutations = {
    setFiles(state, files) {
        state.files = files
    }
}

export default {
    namespaced: true,
    state,
    actions,
    mutations
}