import router from '~/router'
import { ElMessage } from 'element-plus'
import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css' // progress bar style
import { getToken } from '~/utils/auth' // get token from cookie
import getPageTitle from '~/utils/get-page-title'
import { useUserStore } from './store/user'
import { usePermissionStore } from './store/permission'

NProgress.configure({ showSpinner: false }) // NProgress Configuration
const whiteList = ['/login'] // no redirect whitelist

router.beforeEach(async (to, from, next) => {
  NProgress.start() // start progress bar
  
  document.title = getPageTitle(to.meta.title) // set page title

  const hasToken = getToken()
  const store = useUserStore()
  const permissionStore = usePermissionStore()

  if (hasToken) {
    if (to.path === '/login') {
      next({ path: '/' })
      NProgress.done()
    } else {
      if (store.uid > 0) {
        next()
      } else {
        try {
          await store.getInfo()
          const accessRoutes = await permissionStore.generateRoutes()
          accessRoutes.map(r => {
            router.addRoute(r)
          })
          next({ ...to, replace: true })
        } catch (error) {
          await store.resetToken()
          ElMessage.error(error || 'Auth Token Error!')
          next(`/login?redirect=${to.path}`)
          NProgress.done()
        }
      }
    }
  } else {
    if (whiteList.indexOf(to.path) !== -1) {
      next()
    } else {
      next(`/login?redirect=${to.path}`)
      NProgress.done()
    }
  }
})

router.afterEach(() => {
  NProgress.done()
})
