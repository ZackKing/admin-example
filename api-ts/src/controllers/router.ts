import Router from 'koa-router'
import { forIn } from 'lodash'
import Auth from './Auth'

const routerMap: RouterMap = {
  'login': [
    { method: 'post', path: '/', ctl: Auth.login },
  ],
}

const routers = new Router()

const methodMap: { [key in RouteConf['method']]: (router: Router, path: string, ctl: Router.IMiddleware) => Router } = {
  get: (router, path, ctl) => router.get(path, ctl),
  post: (router, path, ctl) => router.post(path, ctl),
  put: (router, path, ctl) => router.put(path, ctl),
  del: (router, path, ctl) => router.del(path, ctl),
  all: (router, path, ctl) => router.all(path, ctl),
}

forIn(routerMap, (routes, prefix) => {
  const sr = new Router()
  routes.forEach(route => registerRoute(sr, route))
  routers.use('/' + prefix, sr.routes(), sr.allowedMethods())
})

export default routers

function registerRoute(router: Router, route: RouteConf) {
  const { method, path, ctl } = route
  if (methodMap[method]) {
    methodMap[method](router, path, ctl)
  } else {
    throw new Error(`Unsupported method: ${method}`)
  }
}
