import UserMdl from '~/models/User'
import Logic from './Loigc'
import GroupUserMdl from '~/models/GroupUser'
import GroupMdl from '~/models/Group'
import AuthLogic from './Auth'

export default class User extends Logic {

  async list(search: KV, opts: KV): Promise<KV> {

    const where: KV = {
      'status[IN]': [UserMdl.MAP_STATUS.valid, UserMdl.MAP_STATUS.invalid]
    }

    for (const k in search) {
      const v = search[k]
      switch (k) {
        case 'name':
          where['name[~]'] = `%${v}%`
          break
        case 'group_ids':
          where['uid[IN]'] = await GroupUserMdl.instance().getUids(v)
          break
        default:
          where[k] = v
          break
      }
    }
    const fields = ['uid', 'name', 'real_name', 'mobile', 'email', 'desc', 'status']
    const data = await UserMdl.instance().findWithTotal(where, fields, opts)
    if (data.total > 0) { // add ext info
      await this.addGroupInfo(data.list)
    }
    return data
  }

  async addGroupInfo(list: KV[], fields: string[] = ['*']): Promise<KV[]> {
    const uids: Int[] = []
    const data: KV = {}

    list.map(v => {
      uids.push(v.uid)
      data[v.uid] = v
      data[v.uid].group = []
    })

    if (uids.length < 1) {
      return Object.values(data)
    }

    const gu = await GroupUserMdl.instance().find({ 'uid[IN]': uids })
    if (gu.length < 1) {
      return Object.values(data)
    }
    const gids = gu.map(v => v.gid)
    const gm = await GroupMdl.instance().findMap({ 'id[IN]': gids }, 'id', fields)

    gu.map(v => {
      if (!gm[v.gid]) {
        return
      }
      data[v.uid].group.push(gm[v.gid])
    })

    return Object.values(data)
  }

  async add(data: KV): Promise<Int> {
    const mdl = UserMdl.instance()
    const cnt = await mdl.count({ name: data.name })
    if (cnt > 0) {
      this.throw(10007)
    }
    const authLogic = AuthLogic.instance()
    data.salt = authLogic.genSalt()
    data.password = authLogic.encodePwd(data.password, data.salt)
    data.status = UserMdl.MAP_STATUS.valid
    data.login_time = 0
    data.pwd_wrong = 0
    return await mdl.insert(data)
  }

}