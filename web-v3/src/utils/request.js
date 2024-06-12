import axios from 'axios'
import { ElMessageBox, ElMessage } from 'element-plus'
import { getToken } from '@/utils/auth'
import router from '@/router'
import { useUserStore } from '@/store/user'

const RE_LOGIN_CODE = [10001, 10002, 10003]

// create an axios instance
const service = axios.create({
  baseURL: process.env.BASE_API, // url = base url + request url
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 10000 // request timeout
})

// request interceptor
service.interceptors.request.use(
  config => {
    // before request sent
    const store = useUserStore()
    if (store.token) {
      config.headers['ADMIN-TOKEN'] = getToken()
    }
    return config
  },
  error => {
    console.error(error) // for debug
    return Promise.reject(error)
  }
)

// response interceptor
service.interceptors.response.use(
  /**
   * If you want to get http information such as headers or status
   * Please return  response => response
  */

  /**
   * Determine the request status by custom code
   * Here is just an example
   * You can also judge the status by HTTP Status Code
   */
  response => {
    const res = response.data
    const store = useUserStore()
    if (typeof res.code === 'undefined') {
      ElMessage({
        message: res.msg ?? 'Wrong response data format! Please contact the system administrator!',
        type: 'error',
        duration: 5 * 1000
      })
      return Promise.reject(new Error('Error Response Data'))
    }

    if (RE_LOGIN_CODE.indexOf(res.code) !== -1) { // for relogin
      if (router.currentRoute.path !== '/login') {
        ElMessageBox.confirm('You have been logged out, you can cancel to stay on this page, or log in again', 'Confirm logout', {
          confirmButtonText: 'Re-Login',
          cancelButtonText: 'Cancel',
          type: 'warning'
        }).then(() => {
          store.resetToken().then(() => {
            location.reload()
          })
        })
      } else {
        ElMessage({
          message: res.msg || 'Error',
          type: 'error',
          duration: 5 * 1000
        })
      }
      const error = new Error(res.msg || 'Error')
      error.code = res.code
      return Promise.reject(error)
    } else { // other error
      if (res.code !== 0) {
        ElMessage({
          message: res.msg || 'Error',
          type: 'error'
        })
        const error = new Error(res.msg || 'Error')
        error.code = res.code
        error.data = res.data ?? {}
        return Promise.reject(error)
      } else {
        return res
      }
    }
  },
  error => {
    console.error('err' + error) // for debug
    ElMessage({
      message: error.message,
      type: 'error',
      duration: 5 * 1000
    })
    return Promise.reject(error)
  }
)

export default service
