import Koa from 'koa'
import time from '~/utils/time'
import { log } from '~/services/log'
import { LogicError } from '~/types/error'

const code401 = new Set([100, 101])

export default async (ctx: Koa.Context, next: Koa.Next) => {

  ctx.set('Access-Control-Allow-Origin', '*')
  ctx.set('Access-Control-Allow-Methods', 'GET, POST, PUT, DELETE, OPTIONS')
  ctx.set('Access-Control-Allow-Headers', '*')

  if (ctx.method === 'OPTIONS') {
    ctx.status = 204
    return ctx
  }

  try {
    
    await next()
    
  } catch (err: any) {
    log.error(`[${ctx.method}]${ctx.path} has error:  ${err.message ?? 'unknow'}`)

    ctx.response.status = 200
    const res: KV = {
      code: -1,
      msg: err.message ?? 'unknow',
      ts: time.unix(),
    }

    if (err instanceof LogicError) {
      res.code = err.code
      res.data = err.data ?? ''

      if (code401.has(err.code)) {
        ctx.response.status = 401
      }
    }

    ctx.body = res
  }
}