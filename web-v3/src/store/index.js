// import Vue from 'vue'
import { createStore } from 'vuex'
import getters from './getters'
import app from './modules/app'
import permission from './modules/permission'
import settings from './modules/settings'
import user from './modules/user'
import common from './modules/common'
import tagsView from './modules/tagsView'

// Vue.use(Vuex)

// const store = new Vuex.Store({
//   modules: {
//     app,
//     permission,
//     settings,
//     user,
//     common,
//     tagsView
//   },
//   getters
// })

export default createStore({
  modules: {
    app,
    permission,
    settings,
    user,
    common,
    tagsView
  },
  getters
})
