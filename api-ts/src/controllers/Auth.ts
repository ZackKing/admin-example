import Controller from './Controller'
import Koa from 'koa'

export default class extends Controller {

  static async login(ctx: Koa.Context) {
    const sd = super._validData(ctx, [
      { key: 'account', type: 'string', required: true },
      { key: 'password', type: 'string', required: true },
    ])
    return super.succ(ctx, {
      uid: 1,
      token: 'xxxxx',
    })
  }

}