import Model from './Model'

export default class GroupMenu extends Model {

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
    return 'group_menu'
  }

  get attr(): KV {
    return {
      id: 'int',
      mid: 'int',
      gid: 'int',
    }
  }

  async getMids(gids: Int[]): Promise<Int[]> {
    if (gids.length == 0) {
      return []
    }
    const list = await this.find({ 'gid[IN]': gids }, ['mid'])
    return list.map(v => v.mid)
  }

  async getGids(mids: Int[]): Promise<Int[]> {
    if (mids.length == 0) {
      return []
    }
    const list = await  await this.find({ 'mid[IN]': mids }, ['gid'])
    return list.map(v => v.gid)
  }

}