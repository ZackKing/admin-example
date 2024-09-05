import http from 'http'
import Koa from 'koa'
import koaBody from 'koa-body'
import conf from '~/conf'
import Task from '~/tasks'
import { log } from '~/services/log'
import errMid from '~/middlewares/error'
import authMid from '~/middlewares/auth'
import reporter from '~/services/reporter'
import routers from '~/controllers/router'
import _ from 'lodash'

log.info(`server start with env: ${conf.env}`)

// http server
const app = new Koa()
app.use(koaBody({ multipart: true }))
app.use(errMid)
app.use(authMid)
app.use(routers.routes()).use(routers.allowedMethods({ throw: true }))
const server = http.createServer(app.callback())

// start listen
server.listen({ host: conf.host, port: conf.port }, () => {
  const msg = `${conf.app} init ok, listen in ${conf.host}:${conf.port} `
  reporter.info(msg)
})

// run task
Task.instance().init()

// process error notice
process.on('uncaughtException', function (error) {
  reporter.error(`uncaughtException: ${error.message}`)
})