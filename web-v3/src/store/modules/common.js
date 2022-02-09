import { getPermissionMenu } from '@/api/common'

const getDefaultState = () => {
  return {
    platform: [],
    accessRoutes: []
  }
}

const state = getDefaultState()

const mutations = {
  SET_PLATFORM: (state, platform) => {
    state.platform = platform
  },
  SET_ACCESS_ROUTES: (state, accessRoutes) => {
    state.accessRoutes = accessRoutes
  }
}

const actions = {
  getAccessRoutes({ commit }) {
    return new Promise((resolve, reject) => {
      getPermissionMenu().then(response => {
        const { data } = response
        commit('SET_ACCESS_ROUTES', data)
        resolve(data)
      }).catch(error => {
        reject(error)
      })
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
