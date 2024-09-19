import { isArray } from 'lodash'
import Model from './Model'

export default class GroupUser extends Model {

  static MAP_STATUS = {
    valid: 1,
    invalid: 0,
  }

  constructor() {
    super()
    this.init()
  }

  get db(): string {
    return 'default'
  }

  get table(): string {
    return 'group_user'
  }

  get attr(): KV {
    return {
      id: 'int',
      uid: 'int',
      gid: 'int',
    }
  }

  async getGids(uid: Int|Int[]): Promise<Int[]> {
    const where: KV = {}
    isArray(uid) ? where['uid[IN]'] = uid : where.uid = uid
    const list = await this.find(where)
    return list.map(v => v.gid)
  }
  
  async getUids(gids: Int[]): Promise<Int[]> {
    const list = await this.find({ 'gid[IN]': gids })
    return list.map(v => v.uid)
  }

}