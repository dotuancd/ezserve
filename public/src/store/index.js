import Vuex from "vuex"
import Vue from "vue"
import defaultLogger from "vuex/dist/logger"
import files from "./modules/files";
import user from "./modules/user";

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'

export default new Vuex.Store({
    modules: {
        files,
        user,
    },
    strict: debug,
    plugins: debug ? [defaultLogger()] : []
})
