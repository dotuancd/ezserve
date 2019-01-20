import Vuex from "vuex"
import Vue from "vue"
import defaultLogger from "vuex/dist/logger"
import files from "./modules/files";

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'

export default new Vuex.Store({
    modules: {
        files
    },
    strict: debug,
    plugins: debug ? [defaultLogger()] : []
})
