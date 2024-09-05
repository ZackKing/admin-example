import conf from '~/conf'
import { Redis } from 'ioredis'
import _ from 'lodash'

class RedisServer {
  private _connestions: { [key: string]: Redis } = {}
  keys: Array<string> = []

  constructor() {
    this.keys = _.keys(conf.redis)
    try { // init connection
      this.keys.map(k => {
        this._connestions[k] = this.newConn(k)
      })
    } catch (err) {
      console.error('redis init error', err)
      process.exit()
    }
  }

  /**
   * cache data with json format and get data with val is null
   * @param rk key
   * @param data default null
   * @param ex ttl
   * @param k redis connections key
   * @returns any
   */
  async cache(rk: string, conn?: string): Promise<KV | null>
  async cache(rk: string, data?: any, ex?: Int, conn?: string): Promise<"OK">
  async cache(rk: string, data: any = null, ex = 3600, conn?: string): Promise<KV | null | "OK"> {
    const rc = this.getConn(conn)
    if (_.isNull(data)) { // get
      const rv = await rc.get(rk)
      return _.isNull(rv) ? null : JSON.parse(rv)
    }

    const rv = JSON.stringify(data) // set with ex
    return await rc.set(rk, rv, 'EX', ex)
  }

  async tryLock(k: string, ex: Int = 10, conn = 'default'): Promise<boolean> {
    if (await this.locked(k, conn)) {
      return false
    }
    return await this.lock(k, ex, conn) > 0
  }

  async locked(k: string, conn = 'default'): Promise<boolean> {
    const ttl = await this.getConn(conn).ttl(`lock:${k}`)
    return ttl > 0
  }

  async lock(k: string, ex: Int = 10, conn = 'default'): Promise<Int> {
    return await this.getConn(conn).set(`lock:${k}`, 'locked', 'EX', ex) == 'OK' ? ex : 0
  }

  async unlock(k: string, conn = 'default') {
    await this.del(`lock:${k}`, conn)
  }

  async get(rk: string, k = 'default') {
    try {
      const rv = await this.getConn(k).get(rk)
      if(rv)
        return JSON.parse(rv)
      else 
        return null
    } catch (error) {
      console.error(error)
      return null
    }
  }

  async del(rk: string, conn = 'default') {
    return await this.getConn(conn).del(rk)
  }

  getConn(conn = 'default'): Redis {
    return this._connestions[conn]
  }

  newConn(k = 'default'): Redis {
    return new Redis({
      host: conf.redis[k].host,
      port: conf.redis[k].port,
      password: conf.redis[k].password,
      db: conf.redis[k].db,
      keyPrefix: conf.redis[k].prefix || '',
    })
  }

}

export default new RedisServer()