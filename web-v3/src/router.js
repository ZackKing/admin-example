
import { createRouter, createWebHistory } from 'vue-router'

import Layout from '~/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * 
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 */
const constantRoutes = [
  {
    path: '/redirect',
    component: Layout,
    hidden: true,
    children: [{
        path: '/redirect/:path(.*)',
        component: () => import('~/views/redirect/index')
    }]
  }, {
    path: '/login',
    component: () => import('~/views/login/index'),
    hidden: true
  }, {
    path: '/self',
    name: '/self',
    redirect: '/self/info',
    component: Layout,
    hidden: true,
    children: [{
      path: 'info',
      name: '/self/info',
      hidden: true,
      component: () => import('~/views/self/info'),
      meta: { title: 'Info' }
    }]
  }, {
    path: '/404',
    component: () => import('~/views/404Page'),
    hidden: true
  }, {
    path: '/:pathMatch(.*)',
    component: () => import('~/views/404Page'),
    hidden: true
  }, {
    path: '/',
    component: Layout
  }
]

/**
 * asyncRoutes
 */
const asyncRoutes = [
  {
    path: '/home',
    name: '/home',
    redirect: '/home/index',
    component: Layout,
    children: [
      { path: '/index', name: '/home/index', component: () => import('~/views/dashboard/index.vue') },
    ],
  }, {
    path: '/panel',
    name: '/panel',
    component: Layout,
    alwaysShow: true,
    children: [
      { path: '/users', name: '/panel/users', component: () => import('~/views/panel/user.vue') },
      { path: '/groups', name: '/panel/groups', component: () => import('~/views/panel/group.vue') },
      { path: '/menus', name: '/panel/menus', component: () => import('~/views/panel/menu.vue') },
    ],
  }, {
    // path: '*',
    path: '/:catchAll(.*)',
    name: '/404',
    redirect: '/404',
    hidden: true,
  }
]

const router = createRouter({
  history: createWebHistory(),
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher
}

export { constantRoutes, asyncRoutes }
export default router
