import Model from './Model'

export default class Group extends Model {

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
    return 'group'
  }

  get attr(): KV {
    return {
      id: 'int',
      name: 'string',
      status: 'int',
      remark: 'string',
      created_time: 'string',
      updated_time: 'string',
    }
  }

}