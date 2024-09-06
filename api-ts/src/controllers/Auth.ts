import AuthLogic from '~/logics/Auth'
import Controller from './Controller'
import Koa from 'koa'

export default class extends Controller {

  static async login(ctx: Koa.Context) {
    const sd = super._validData(ctx, [
      { key: 'account', type: 'string', min: 3, required: true, rename: 'name' },
      { key: 'password', type: 'string', min: 6, required: true },
    ])

    const ip = super.ip(ctx)
    const data = await AuthLogic.instance().login(sd, { ip })
    return super.succ(ctx, data)
  }

  static async renewToken(ctx: Koa.Context) {
  }
  
  static async jwtInfo(ctx: Koa.Context) {
  }
  
  static async menu(ctx: Koa.Context) {
    super.succ(
      ctx,
      await AuthLogic.instance().userMenu(super.uid(ctx))
    )
  }
  
  static async changPwd(ctx: Koa.Context) {
  }
  
  static async groupList(ctx: Koa.Context) {
  }
  
  static async addGroup(ctx: Koa.Context) {
  }
  
  static async editGroup(ctx: Koa.Context) {
  }
  
  static async setGroupUser(ctx: Koa.Context) {
  }
  
  static async setGroupMenu(ctx: Koa.Context) {
  }
  
  static async menuTree(ctx: Koa.Context) {
  }
  
  static async menuInfo(ctx: Koa.Context) {
  }
  
  static async setMenuGroup(ctx: Koa.Context) {
  }
  


}