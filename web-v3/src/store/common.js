import { defineStore } from 'pinia'
import { getPermissionMenu } from '~/api/common'
import { ref } from 'vue'

// Setup Store example
export const useCommonStore = defineStore('common', () => {
  const platform = ref([])
  const accessRoutes = ref([])

  function getAccessRoutes() {
    return new Promise((resolve, reject) => {
      getPermissionMenu().then(response => {
        const { data } = response
        accessRoutes.value = data
        resolve(data)
      }).catch(error => {
        reject(error)
      })
    })
  }

  function setPlatform(plf) {
    platform.value = plf
  }

  return { platform, accessRoutes, getAccessRoutes, setPlatform }
})
