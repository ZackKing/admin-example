import Koa from 'koa'
import { LogicError } from '~/types/error'
import { isArray, join } from 'lodash'
import Auth from '~/logics/Auth'

const WHITE_LIST = new Set(['/login'])

export default async (ctx: Koa.Context, next: Koa.Next) => {

  if (!WHITE_LIST.has(ctx.path) && ctx.request.method != 'OPTIONS') {
    _valid(ctx)
  }

  await next()

}

function _valid(ctx: Koa.Context) {
  const token = _getStrHeader(ctx, 'admin-token')
  if (!token) {
    throw new LogicError(100)
  }

  const au = Auth.instance().decodeToken(token)
  if (!au || !au.uid) {
    throw new LogicError(101)
  }

  ctx.uid = au.uid
}

function _getStrHeader(ctx: Koa.Context, key: string): string {
  return isArray(ctx.header[key]) ? join(ctx.header[key], '') : ctx.header[key] ?? ''
}
