import User from '~/models/User'
import Controller from './Controller'
import Koa from 'koa'

export default class extends Controller {

  static async self(ctx: Koa.Context) {
    const uid = super.uid(ctx)

    const data = await User.instance().findOne({ uid })
    if (data) {
      delete data.password
      delete data.salt
    }
    super.succ(ctx, data)
  }

  static async editSelf(ctx: Koa.Context) {

  }

  static async list(ctx: Koa.Context) {

  }

  static async info(ctx: Koa.Context) {

  }

  static async add(ctx: Koa.Context) {

  }

  static async updateInfo(ctx: Koa.Context) {

  }

  static async status(ctx: Koa.Context) {

  }

  static async setGroup(ctx: Koa.Context) {

  }

}