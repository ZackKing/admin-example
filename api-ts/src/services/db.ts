import knex, { Knex } from 'knex'
import conf from '~/conf'

export default class DB {

  static defaultKey = 'default'
  private static _instance: DB

  static instance(): DB {
    if (!DB._instance) {
      DB._instance = new DB()
    }
    return DB._instance
  }

  private _conn: { [key: string]: Knex } = {}
  private _conf: KV

  constructor() {
    this._conf = conf.db
  }

  getConn(k = DB.defaultKey): Knex {
    if (typeof this._conn[k] == 'undefined') {
      this._conn[k] = this._connect(this._conf[k])
    }
    return this._conn[k]
  }

  _connect(cfg: KV): Knex {
    return knex({
      client: 'mysql2',
      connection: {
        host: cfg.host,
        port: cfg.port,
        user: cfg.username,
        password: cfg.password,
        database: cfg.db,
      },
      pool: { min: cfg.min || 1, max: cfg.max || 100 }
    })
  }

}
