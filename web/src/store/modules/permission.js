import { asyncRoutes, constantRoutes } from '@/router'
import store from '@/store'
import _ from 'lodash'
// import { deepClone } from '@/utils'

/**
 * Use meta.role to determine if the current user has permission
 * @param roles
 * @param route
 */
// function hasPermission(roles, route) {
//   if (route.meta && route.meta.roles) {
//     return roles.some(role => route.meta.roles.includes(role))
//   } else {
//     return true
//   }
// }

/**
 * Filter asynchronous routing tables by recursion
 * @param routes asyncRoutes
 * @param apiRoutes
 */
export function filterAsyncRoutes(routes, apiRoutes, accessUrls) {
  const access = {
    '/404': { allow: true, name: '', icon: '', sort: 0 }
  }

  function getAccessName(arr) {
    arr.forEach(item => {
      access[item.uri] = {
        allow: true,
        name: item.name,
        icon: item.icon,
        sort: item.sort
      }
      if (item.sub_menu) {
        getAccessName(item.sub_menu)
      }

      // 记录可访问的url
      accessUrls.push(item.uri)
    })
  }

  getAccessName(apiRoutes)

  function filterValidRoute(array) {
    return array.filter((item, index) => {
      if (!access[item.name] || !access[item.name].allow) {
        return false
      } else {
        item.sort = access[item.name].sort
        if (typeof item.meta === 'undefined') {
          item.meta = {}
        }
        item.meta.title = access[item.name].name
        item.meta.icon = access[item.name].icon
        if (item.children && item.children.length > 0) {
          item.children = filterValidRoute(item.children)
        }
        item.children = _.sortBy(item.children, ['sort'])
        return true
      }
    })
  }
  routes = filterValidRoute(routes)
  routes.sort((a, b) => { return a.sort - b.sort })
  return routes
}

const state = {
  routes: [],
  addRoutes: [],
  accessUrls: []
}

const mutations = {
  SET_ROUTES: (state, routes) => {
    state.addRoutes = routes
    state.routes = constantRoutes.concat(routes)
  },
  SET_ACCESS_URLS: (state, accessUrls) => {
    state.accessUrls = accessUrls
  }
}

const actions = {
  generateRoutes({ commit }, roles) {
    return new Promise(async(resolve) => {
      const apiRoutes = await store.dispatch('common/getAccessRoutes')
      const accessUrls = []
      const filterRoutes = filterAsyncRoutes(asyncRoutes, apiRoutes, accessUrls)

      commit('SET_ROUTES', filterRoutes)
      // 记录可访问的url到store，用于做button级的权限控制
      commit('SET_ACCESS_URLS', accessUrls)
      resolve(filterRoutes)
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
