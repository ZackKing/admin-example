import { defineStore } from 'pinia'
import { ref } from 'vue'
import { asyncRoutes, constantRoutes } from '@/router'
import { useCommonStore } from './common'
import { sortBy } from 'lodash-es'

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
        item.children = sortBy(item.children, ['sort'])
        return true
      }
    })
  }
  routes = filterValidRoute(routes)
  routes.sort((a, b) => { return a.sort - b.sort })
  return routes
}

export const usePermissionStore = defineStore('permission', () => {
  const routes = ref([])
  const addRoutes = ref([])
  const accessUrls = ref([])

  function generateRoutes() {
    return new Promise(async (resolve) => {
      const commonStore = useCommonStore()
      const apiRoutes = await commonStore.getAccessRoutes()
      const aUrls = []
      const filterRoutes = filterAsyncRoutes(asyncRoutes, apiRoutes, aUrls)

      // set routes
      addRoutes.value = filterRoutes
      routes.value = constantRoutes.concat(filterRoutes)

      // 记录可访问的url到store，用于做button级的权限控制
      accessUrls.value = aUrls
      resolve(filterRoutes)
    })
  }

  return {
    routes,
    addRoutes,
    accessUrls,
    generateRoutes,
  }
})

