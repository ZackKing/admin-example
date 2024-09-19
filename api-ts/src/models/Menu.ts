import comm from '~/utils/comm'
import Model from './Model'

export default class Menu extends Model {

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
    return 'menu'
  }

  get attr(): KV {
    return {
      id: 'int',
      name: 'string',
      uri: 'string',
      level: 'int',
      pid: 'int',
      icon: 'string',
      status: 'int',
      sort: 'int',
      remark: 'string',
      created_time: 'string',
      updated_time: 'string',
    }
  }

  async getAll(tree = true): Promise<KV | KV[]>
  {
    const list = await this.find({})
    return tree ? comm.genTree(list, 'id', 'pid', 'sub_menu') : list
  }

}