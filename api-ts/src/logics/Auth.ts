import conf, { constant } from '~/conf'
import Logic from './Loigc'
import jwt from 'jsonwebtoken'
import { LogicError } from '~/types/error'
import UserMdl from '~/models/User'
import comm from '~/utils/comm'
import time from '~/utils/time'
import GroupUser from '~/models/GroupUser'
import MenuMdl from '~/models/Menu'
import GroupMenu from '~/models/GroupMenu'

export default class Auth extends Logic {

  async login(data: KV, opts: KV): Promise<KV> {
    const user = await UserMdl.instance().findOne({ name: data.name })
    if (!user) {
      this.throw(10008)
    }
    if (this.encodePwd(data.password, user.salt) != user.password) {
      const limit = await this.checkBan(user)
      this.throw(10008, '', {}, [limit])
    }

    await this.loginOk(user.uid, opts)

    return {
      uid: user.uid,
      token: this.genToken(user.uid)
    }
  }

  async loginOk(uid: Int, opts: KV): Promise<boolean> {
    const up: KV = { pwd_wrong: 0, login_time: time.unix() }
    if (opts.ip) {
      up.login_ip = up.ip
    }
    const ok = await UserMdl.instance().update({ uid }, up, 1)
    return ok > 0
   }

  genToken(uid: number): string {
    const token = jwt.sign({ uid }, conf.jwt.secret, conf.jwt.options)
    return token
  }

  decodeToken(token: string): KV {
    try {
      const data = jwt.verify(token, conf.jwt.secret, conf.jwt.options)
      return data as KV
    } catch (error) {
      throw new LogicError(101)
    }
  }

  async checkBan(user: KV): Promise<Int> {
    user.pwd_wrong += 1
    const up = { pwd_wrong: user.pwd_wrong++, status: user.status }

    if (user.pwd_wrong >= constant.PWD_WRONG_LIMIT) {
      up.status = UserMdl.MAP_STATUS.invalid
    }

    await UserMdl.instance().update({ uid: user.uid }, up)
    
    if (up.status == UserMdl.MAP_STATUS.invalid) {
      this.throw(10009)
    }

    return constant.PWD_WRONG_LIMIT - user.pwd_wrong
  }

  encodePwd(pwd: string, salt: string): string {
    return comm.md5(`${salt}${comm.md5(pwd)}`)
  }

  async userMenu(uid: Int): Promise<KV> {
    const gids = await GroupUser.instance().getGids(uid)
    let tree = {}
    if (this.isSuperAdmin(gids)) {
      tree = await this.getMenuTree()
    } else {
      const mids = await GroupMenu.instance().getMids(gids)
      tree = mids.length > 0 ? await this.getMenuTree(mids) : {}
    }
    return tree
  }

  isSuperAdmin(gids: Int[]): boolean {
    return gids.indexOf(1) != -1
  }


  async getMenuTree(mids?: Int[]): Promise<KV> {
    const where:KV = { status: MenuMdl.MAP_STATUS.valid }
    if (mids && mids.length > 0) {
      where['id[in]'] = mids
    }
    const list = await MenuMdl.instance().find(where, ['*'], { orderBy: ['sort', 'asc'] })
    return comm.genTree(list, 'id', 'pid', 'sub_menu')
  }

}