import Router from 'koa-router'
import Auth from './Auth'
import User from './User'

const routerMap: RouterMap = {
  '': [
    { method: 'post', path: '/login', ctl: Auth.login },
    { method: 'post', path: '/renewToken', ctl: Auth.renewToken },
  ],

  'auth': [
    { method: 'get', path: '/jwt', ctl: Auth.jwtInfo },
    { method: 'get', path: '/menu', ctl: Auth.menu },
    { method: 'post', path: '/password', ctl: Auth.changPwd },
  ],

  'group': [
    { method: 'get', path: '/', ctl: Auth.groupList },
    { method: 'post', path: '/', ctl: Auth.addGroup },
    { method: 'post', path: '/edit', ctl: Auth.editGroup },
    { method: 'post', path: '/user', ctl: Auth.setGroupUser },
    { method: 'post', path: '/menu', ctl: Auth.setGroupMenu },
  ],

  'menu': [
    { method: 'get', path: '/', ctl: Auth.menuTree },
    { method: 'get', path: '/info', ctl: Auth.menuInfo },
    { method: 'post', path: '/group', ctl: Auth.setMenuGroup },
  ],

  'user': [
    { method: 'get', path: '/self', ctl: User.self },
    { method: 'post', path: '/self', ctl: User.editSelf },
    { method: 'get', path: '/list', ctl: User.list },
    { method: 'get', path: '/info', ctl: User.info },
    { method: 'post', path: '/add', ctl: User.add },
    { method: 'post', path: '/update', ctl: User.updateInfo },
    { method: 'post', path: '/status', ctl: User.status },
    { method: 'post', path: '/group', ctl: User.setGroup },
  ]

}

const methodMap: { [key in RouteConf['method']]: (router: Router, path: string, ctl: Router.IMiddleware) => Router } = {
  get: (router, path, ctl) => router.get(path, ctl),
  post: (router, path, ctl) => router.post(path, ctl),
  put: (router, path, ctl) => router.put(path, ctl),
  del: (router, path, ctl) => router.del(path, ctl),
  all: (router, path, ctl) => router.all(path, ctl),
}

const routers = new Router()

for (const k in routerMap) {
  const routes = routerMap[k]
  const r = new Router()
  routes.forEach(route => registerRoute(r, route))
  routers.use(k == '' ? '' : `/${k}`, r.routes(), r.allowedMethods())
}

export default routers

function registerRoute(router: Router, route: RouteConf) {
  const { method, path, ctl } = route
  if (methodMap[method]) {
    methodMap[method](router, path, ctl)
  } else {
    throw new Error(`Unsupported method: ${method}`)
  }
}
