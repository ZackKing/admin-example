import Model from './Model'

export default class AccessLog extends Model {

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
    return 'access_log'
  }

  get attr(): KV {
    return {
      id: 'int',
      access_id: 'string',
      uid: 'int',
      method: 'string',
      path: 'string',
      header: 'string',
      query: 'string',
      body: 'string',
      ip: 'string',
      response: 'string',
      created_time: 'string',
      updated_time: 'string',
    }
  }

}