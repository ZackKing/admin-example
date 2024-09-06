import Model from './Model'

export default class User extends Model {

  private static _instance: User

  static instance(): User {
    if (!User._instance) {
      User._instance = new User()
    }
    return User._instance
  }

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
    return 'user'
  }

  get attr(): KV {
    return {
      uid: 'int',
      name: 'string',
      password: 'string',
      salt: 'string',
      real_name: 'string',
      mobile: 'string',
      email: 'string',
      desc: 'string',
      login_time: 'int',
      pwd_wrong: 'int',
      status: 'int',
      created_time: 'string',
      updated_time: 'string',
    }
  }

}