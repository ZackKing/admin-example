import { login as userLogin, getInfo as userGetInfo } from '@/api/user'
import { getToken, setToken, removeToken } from '@/utils/auth'
import { resetRouter } from '@/router'
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useUserStore = defineStore('user', () => {
  const token = ref(getToken())
  const uid = ref('')
  const name = ref('')
  const avatar = ref('')
  const roles = ref([])

  function reset() {
    token.value = getToken()
    uid.value = ''
    name.value = ''
    avatar.value = ''
    roles.value = []
  }

  function login(userInfo) {
    return new Promise((resolve, reject) => {
      userLogin({ ...userInfo }).then(res => {
        const { data } = res
        token.value = data.token
        setToken(data.token)
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  }

  function getInfo() {
    return new Promise((resolve, reject) => {
      userGetInfo().then(res => {
        const { data } = res
        if (!data) {
          reject('Verification failed, please Login again.')
        }
        data.roles = ['admin']
        // const { roles, name, head_icon, uid } = data
        if (!data.roles || data.roles.length <= 0) { // roles must be a non-empty array
          reject('getInfo: roles must be a non-null array!')
        }
        roles.value = data.roles
        uid.value = data.uid
        name.value = data.name
        avatar.value = data.avatar
        resolve(data)
      }).catch(error => {
        reject(error)
      })
    })
  }

  function logout() {
    return new Promise((resolve, reject) => {
      removeToken()
      resetRouter()
      reset()
      resolve()
    })
  }

  function resetToken() {
    return new Promise(resolve => {
      removeToken()
      reset()
      resolve()
    })
  }

  return {
    token,
    uid,
    name,
    avatar,
    roles,
    login,
    getInfo,
    logout,
    resetToken,
  }
})
