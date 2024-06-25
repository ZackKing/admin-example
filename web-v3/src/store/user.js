import { login as userLogin, getInfo as userGetInfo } from '~/api/user'
import { getToken, setToken, removeToken } from '~/utils/auth'
import { resetRouter } from '~/router'
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useUserStore = defineStore('user', () => {
  const token = ref(getToken())
  const uid = ref(0)
  const name = ref('')
  const nickname = ref('')
  const avatar = ref('')

  function reset() {
    token.value = getToken()
    uid.value = 0
    name.value = ''
    nickname.value = ''
    avatar.value = ''
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
          reject('Auth timeout!')
        }
        uid.value = data.uid
        name.value = data.name
        nickname.value = data.real_name
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
    nickname,
    avatar,
    login,
    getInfo,
    logout,
    resetToken,
  }
})
