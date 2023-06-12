import router from '@/router'
import { ElMessage } from 'element-plus'
import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css' // progress bar style
import { getToken } from '@/utils/auth' // get token from cookie
import getPageTitle from '@/utils/get-page-title'
import { useUserStore } from './store/user'
import { usePermissionStore } from './store/permission'

NProgress.configure({ showSpinner: false }) // NProgress Configuration
const whiteList = ['/login'] // no redirect whitelist

router.beforeEach(async (to, from, next) => {
  NProgress.start() // start progress bar

  document.title = getPageTitle(to.meta.title) // set page title

  const hasToken = getToken() // determine whether the user has logged in

  const store = useUserStore() // get user store
  const permissionStore = usePermissionStore()

  if (hasToken) {
    if (to.path === '/login') {
      next({ path: '/' }) // if is logged in, redirect to the home page
      NProgress.done()
    } else {
      const hasRoles = store.roles && store.roles.length > 0 // determine whether the user has obtained his permission roles through getInfo
      if (hasRoles) {
        next()
      } else {
        try {
          // get user info
          // note: roles must be a object array! such as: ['admin'] or ,['developer','editor']
          const { roles } = await store.getInfo()

          // generate accessible routes map based on roles
          const accessRoutes = await permissionStore.generateRoutes(roles)
          accessRoutes.map(r => {
            router.addRoute(r)
          })
          // router.addRoutes(accessRoutes)

          // hack method to ensure that addRoutes is complete
          // set the replace: true, so the navigation will not leave a history record
          next({ ...to, replace: true })
        } catch (error) {
          console.error(error)
          await store.resetToken() // remove token and go to login page to re-login
          ElMessage.error(error || 'Has Error')
          next(`/login?redirect=${to.path}`)
          NProgress.done()
        }
      }
    }
  } else {
    /* has no token*/
    if (whiteList.indexOf(to.path) !== -1) {
      // in the free login whitelist, go directly
      next()
    } else {
      // other pages that do not have permission to access are redirected to the login page.
      next(`/login?redirect=${to.path}`)
      NProgress.done()
    }
  }
})

router.afterEach(() => {
  NProgress.done() // finish progress bar
})
