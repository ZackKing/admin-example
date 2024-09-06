import Koa from 'koa'
import redis from '~/services/redis'
import { errMsg, LogicError } from '~/types/error'
import time from '~/utils/time'
import Validator, { ValidatorRule } from '~/utils/validator'

export default class Controller {

  static uid(ctx: Koa.Context): number {
    return ctx.uid ?? 0
  }

  static succ(ctx: Koa.Context, data?: any, msg?: string): void {
    ctx.body = Controller._genRes(0, data || null, msg || 'ok')
  }

  static err(ctx: Koa.Context, code = 0, data: any = null, msg?: string): void {
    ctx.body = Controller._genRes(code, data, msg || 'error')
  }

  static async lock(k: string, ex: Int = 10): Promise<void> {
    if (!await redis.tryLock(k, ex)) {
      throw new LogicError(3)
    }
  }

  static async unlock(k: string): Promise<void> {
    await redis.unlock(k)
  }

  static ip(ctx: Koa.Context): string {
    let realip = ''
    if (ctx.headers['x-forwarded-for']) {
      const forwarded = ctx.headers['x-forwarded-for'] as string
      const ips = forwarded.split(',').map(ip => ip.trim())
      for (const ip of ips) {
        if (ip != 'unknown') {
          realip = ip
          break
        }
      }
    } else if (ctx.headers['x-client-ip']) {
      realip = ctx.headers['x-client-ip'] as string
    } else if (ctx.req.socket.remoteAddress) {
      realip = ctx.req.socket.remoteAddress
    } else {
      realip = '0.0.0.0'
    }

    const ipMatch = realip.match(/[\d\.]{7,15}/)
    realip = ipMatch ? ipMatch[0] : '0.0.0.0'
    return realip
  }

  static _validQuery(ctx: Koa.Context, rules: Array<ValidatorRule>) {
    const query = ctx.query ?? {}
    return Controller._valid(ctx, rules, query)
  }

  static _validData(ctx: Koa.Context, rules: Array<ValidatorRule>) {
    const data = ctx.request.body?.fields ?? ctx.request.body ?? {}
    return Controller._valid(ctx, rules, data)
  }

  static _valid(ctx: Koa.Context, rules: Array<ValidatorRule>, data: KV): KV {
    try {
      return Validator.check(data, rules)
    } catch (error: any) {
      ctx.assert(false, 400, error instanceof Error ? error.message : JSON.stringify(error))
    }
  }

  static _genRes(code: number, data: any = null, msg?: string) {
    const res: ResData = { code, data, msg: msg ?? Controller._genMsg(code), ts: time.unix() }
    return res
  }

  static _genMsg(code: number) {
    return errMsg[code] ?? 'unknow error'
  }

}
