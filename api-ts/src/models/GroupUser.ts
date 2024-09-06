import Model from './Model'

export default class GroupUser extends Model {

  private static _instance: GroupUser

  static instance(): GroupUser {
    if (!GroupUser._instance) {
      GroupUser._instance = new GroupUser()
    }
    return GroupUser._instance
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
    return 'group_user'
  }

  get attr(): KV {
    return {
      id: 'int',
      uid: 'int',
      gid: 'int',
    }
  }

  async getGids(uid: Int): Promise<Int[]> {
    const list = await this.find({ uid })
    return list.map(v => v.gid)
  }

}