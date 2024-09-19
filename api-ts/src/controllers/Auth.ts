import AuthLogic from '~/logics/Auth'
import Controller from './Controller'
import Koa from 'koa'
import GroupMdl from '~/models/Group'
import GroupUserMdl from '~/models/GroupUser'
import GroupMenuMdl from '~/models/GroupMenu'
import MenuMdl from '~/models/Menu'
import { LogicError } from '~/types/error'

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
    const uid = super.uid(ctx)
    return super.succ(ctx, {
      uid,
      token: await AuthLogic.instance().genToken(uid)
    })
  }
  
  static async jwtInfo(ctx: Koa.Context) {}
  
  static async menu(ctx: Koa.Context) {
    super.succ(
      ctx,
      await AuthLogic.instance().userMenu(super.uid(ctx))
    )
  }
  
  static async changPwd(ctx: Koa.Context) {
    const sd = super._validData(ctx, [
      { key: 'old_password', type: 'string', min: 6, required: true },
      { key: 'password', type: 'string', min: 6, required: true },
    ])
    const uid = super.uid(ctx)
    const ok = await AuthLogic.instance().changePwd(uid, sd.old_password, sd.password)
    if (!ok) {
      throw new LogicError(10004)
    }
    return super.succ(ctx)
  }
  
  static async groupList(ctx: Koa.Context) {
    const list = await GroupMdl.instance().find({})
    for (const v of list) {
      v.uids = await GroupUserMdl.instance().getUids([v.id])
      v.mids = await GroupMenuMdl.instance().getMids([v.id])
    }
    super.succ(ctx, list)
  }
  
  static async addGroup(ctx: Koa.Context) {
    const sd = super._validData(ctx, [
      { key: 'name', type: 'string', min: 2, required: true },
      { key: 'remark', type: 'string', default: '' },
    ])
    const id = await GroupMdl.instance().save({ id: 0 }, {  name: sd.name, remark: sd.remark })
    if (id  < 1) {
      throw new LogicError(1, 'add group error')
    }
    return super.succ(ctx, id)
  }
  
  static async editGroup(ctx: Koa.Context) {
    const sd = super._validData(ctx, [
      { key: 'id', type: 'int', min: 1, required: true },
      { key: 'name', type: 'string', min: 2, required: true },
      { key: 'remark', type: 'string' },
      { key: 'status', type: 'enum', values: [0, 1] },
    ])
    const info = await GroupMdl.instance().findOne({ id: sd.id }, ['id'])
    if (!info) {
      throw new LogicError(1, 'group not found')
    }
    const ok = await GroupMdl.instance().update({ id: sd.id }, sd)
    if (!ok) {
      throw new LogicError(1, 'group edit found')
    }
    super.succ(ctx)
  }
  
  static async setGroupUser(ctx: Koa.Context) {
    const sd = super._validData(ctx, [
      { key: 'id', type: 'int', min: 1, required: true },
      { key: 'uids', type: 'array', default: [] },
    ])
    const ok = await AuthLogic.instance().setGroupUser(sd.id, sd.uids)
    if (!ok) {
      throw new LogicError(1, 'set group user error')
    }
    super.succ(ctx)
  }
  
  static async setGroupMenu(ctx: Koa.Context) {
    const sd = super._validData(ctx, [
      { key: 'id', type: 'int', min: 1, required: true },
      { key: 'menu_ids', type: 'array', default: [] },
    ])
    const ok = await AuthLogic.instance().setGroupMenu(sd.id, sd.menu_ids)
    if (!ok) {
      throw new LogicError(1, 'set group menu error')
    }
    super.succ(ctx)
  }
  
  static async menuTree(ctx: Koa.Context) {
    super.succ(
      ctx,
      await MenuMdl.instance().getAll()
    )
  }
  
  static async menuInfo(ctx: Koa.Context) {
    const query = super._validQuery(ctx, [
      { key: 'id', type: 'int', min: 3, required: true }
    ])

    const info = await MenuMdl.instance().findOne({ id: query.id })
    if (!info) {
      return super.succ(ctx, {})
    }
    info.group_ids = await GroupMenuMdl.instance().getGids([info.id])
    return super.succ(ctx, info)
  }
  
  static async setMenuGroup(ctx: Koa.Context) {
    const sd = super._validData(ctx, [
      { key: 'id', type: 'int', min: 1, required: true },
      { key: 'group_ids', type: 'array', default: [] },
    ])
    const ok = await AuthLogic.instance().setMenuGroup(sd.id, sd.group_ids)
    if (!ok) {
      throw new LogicError(1, 'set menu group error')
    }
    super.succ(ctx)
  }
  
}