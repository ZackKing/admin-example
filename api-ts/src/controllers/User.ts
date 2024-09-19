import UserMdl from '~/models/User'
import Controller from './Controller'
import Koa from 'koa'
import { isEmpty } from 'lodash'
import UserLogic from '~/logics/User'
import { LogicError } from '~/types/error'
import AuthLogic from '~/logics/Auth'

export default class extends Controller {

  static async self(ctx: Koa.Context) {
    const uid = super.uid(ctx)

    const data = await UserMdl.instance().findOne({ uid })
    if (data) {
      delete data.password
      delete data.salt
    }
    super.succ(ctx, data)
  }

  static async editSelf(ctx: Koa.Context) {
    const data = super._validData(ctx, [
      { key: 'real_name', type: 'string', ignore: ['', null] },
      { key: 'mobile', type: 'string', ignore: ['', null] },
      { key: 'email', type: 'string', ignore: ['', null] },
      { key: 'desc', type: 'string', ignore: ['', null] },
    ])
    if (isEmpty(data)) {
      return super.succ(ctx, null)
    }

    const uid = super.uid(ctx)
    const ok = await UserMdl.instance().update({ uid }, data)
    return super.succ(ctx, ok)
  }

  static async list(ctx: Koa.Context) {
    const query = super._validQuery(ctx, [
      { key: 'name', type: 'string', ignore: ['', null] },
      { key: 'group_ids', type: 'array', default: []},
      { key: 'size', type: 'int', default: 10 },
      { key: 'page', type: 'int', default: 1 },
    ])

    const search: KV = {}
    if (query.name) search.name = query.name
    if (query.group_ids.length > 1) search.group_ids = query.group_ids

    const opts: KV = {
      limit: query.size,
      offset: query.size * (query.page - 1),
    }

    const data = await UserLogic.instance().list(search, opts)
    super.succ(ctx, data)
  }

  static async info(ctx: Koa.Context) {
    const query = super._validQuery(ctx, [
      { key: 'uid', type: 'int', min: 1, required: true },
    ])
    const cols = ['uid', 'name', 'real_name', 'mobile', 'email', 'desc', 'status']
    const info = await UserMdl.instance().findOne({ uid: query.uid }, cols)
    if (info) {
      await UserLogic.instance().addGroupInfo([info], ['id', 'name', 'status'])
    }
    super.succ(ctx, info)
  }

  static async add(ctx: Koa.Context) {
    const data = super._validData(ctx, [
      { key: 'name', type: 'string', min: 6, required: true },
      { key: 'password', type: 'string', min: 6, required: true },
      { key: 'real_name', type: 'string' },
      { key: 'mobile', type: 'string' },
      { key: 'email', type: 'string' },
      { key: 'desc', type: 'string' },
    ])
    const uid = await UserLogic.instance().add(data)
    if (!uid) {
      throw new LogicError(10006)
    }
    super.succ(ctx, { uid })
  }

  static async update(ctx: Koa.Context) {
    const data = super._validData(ctx, [
      { key: 'uid', type: 'int', min: 1, required: true },
      { key: 'name', type: 'string', min: 6, required: true },
      { key: 'real_name', type: 'string' },
      { key: 'mobile', type: 'string' },
      { key: 'email', type: 'string' },
      { key: 'desc', type: 'string' },
    ])

    const ok = await UserMdl.instance().update({ uid: data.uid }, data, 1)
    super.succ(ctx, ok)
  }

  static async status(ctx: Koa.Context) {
    const data = super._validData(ctx, [
      { key: 'uid', type: 'int', min: 1, required: true },
      { key: 'status', type: 'enum', values: [0, 1] },
    ])

    const ok = await UserMdl.instance().update({ uid: data.uid }, { status: data.status }, 1)
    if (!ok) {
      throw new LogicError(1, 'update user status error')
    }
    super.succ(ctx)
  }

  static async setGroup(ctx: Koa.Context) {
    const data = super._validData(ctx, [
      { key: 'id', type: 'int', min: 1, required: true },
      { key: 'group_ids', type: 'array', default: [] },
    ])

    const ok = await AuthLogic.instance().setUserGroup(data.id, data.group_ids)
    if (!ok) {
      throw new LogicError(1, 'set user group error')
    }
    super.succ(ctx)
  }

}