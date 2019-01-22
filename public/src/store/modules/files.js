
import fileRequest from '../../api/files'

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
    },
    create({commit}, file, callback) {
        fileRequest.create(file, callback)
    }
}

const mutations = {
    setFiles(state, files) {
        state.files = files
    },
    create(state, file) {
        state.files.items.push(file)
    }
}

export default {
    namespaced: true,
    state,
    actions,
    mutations
}