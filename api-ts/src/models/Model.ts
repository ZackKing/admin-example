import { Knex } from 'knex'
import { isArray, keys } from 'lodash'
import DB from '~/services/db'
import { tLog } from '~/services/log'

type FindOpts = {
  limit?: Int
  offset?: Int
  orderBy?: string[]
}

export default abstract class Model {

  private static instances: { [key: string]: Model } = {}

  static instance<T extends Model>(this: new () => T): T {
    const className = this.name
    if (!Model.instances[className]) {
      Model.instances[className] = new this()
    }
    return Model.instances[className] as T
  }

  abstract get db(): string
  abstract get table(): string
  abstract get attr(): KV
  conn: Knex

  init() {
    this.conn = DB.instance().getConn(this.db)
  }

  async raw(sql: string, bindings?: any[]) {
    return await this.conn.raw(sql, bindings ?? [])
  }

  async count(where: KV): Promise<number> {
    const rs = await this.builderWhere(where).count()
    return rs[0]['count(*)'] as number
  }

  async getById(id: number, fields: (KV | string)[] = ['*']): Promise<KV | undefined> {
    return await this.findOne({ id }, fields)
  }

  async getByField(value: any, field: string, fields: (KV | string)[] = ['*']): Promise<KV | undefined>  {
    const where: KV = {}
    where[field] = value
    return await this.findOne(where, fields)
  }

  async findOne(where: KV, fields: (KV | string)[] = ['*']): Promise<KV | undefined> {
    const res = await this.find(where, fields, { limit: 1 })
    return res[0]
  }

  async findCol(where: KV, field: string, defaultVal?: any): Promise<any> {
    const row = await this.findOne(where, [field])
    return row ? row[field] : defaultVal ?? null
  }

  async findWithTotal(where: KV, fields: (KV | string)[] = ['*'], opts: FindOpts = {}): Promise<{ list: KV[], total: Int }> {
    return {
      list: await this.find(where, fields, opts),
      total: await this.count(where)
    }
  }

  async findMap(where: KV, key: string, fields: (KV | string)[] = ['*'], opts: FindOpts = {}): Promise<KV> {
    const map: KV = {}
    const list = await this.find(where, fields, opts)
    list.map(v => {
      map[v[key]] = v
    })
    return map
  }

  async find(where: KV, fields: (KV | string)[] = ['*'], opts: FindOpts = {}): Promise<KV[]> {
    const builder = this.builderWhere(where)
    if (!isArray(fields)) fields = [fields]
    if (fields.includes('*') && this.fields.length != 0) fields = this.fields

    if (opts.limit) builder.limit(opts.limit) // limit / offset
    builder.offset(opts.offset ?? 0)
    if (opts.orderBy) builder.orderBy(opts.orderBy[0] ?? 'id', opts.orderBy[1] ?? 'desc')

    return await builder.select(fields)
  }

  async insert(data: KV): Promise<Int>
  async insert(data: KV | KV[]): Promise<Int | Int[]> {
    try {
      const ids = await this.conn(this.table).insert(data)
      return ids.length == 1 ? (ids[0] ?? 0) : ids
    } catch (error) {
      tLog.error('model_insert', { error })
      return 0
    }
  }

  async update(where: KV, data: KV, limit: number = 0): Promise<number> {
    const builder = this.builderWhere(where)
    if (limit > 0) {
      builder.limit(limit)
    }
    return await builder.update(data)
  }

  async save(where: KV, data: KV): Promise<number> {
    const row = await this.findOne(where, ['id'])
    if (row) {
      await this.update({ id: row.id }, data, 1)
      return row.id
    } else {
      return await this.insert(data)
    }
  }

  async delete(where: KV, num: Int = 0): Promise<Int> {
    const builder = this.builderWhere(where)
    if (num > 0) {
      const cnt = await builder.count()
      if (cnt != num) {
        return 0
      }
    }
    return await builder.delete()
  }

  get fields(): string[] {
    return keys(this.attr)
  }

  builder(): Knex.QueryBuilder {
    return this.conn(this.table)
  }

  builderWhere(where: KV): Knex.QueryBuilder {
    const builder = this.builder()
    for (const k in where) {
      const m = k.match(/(.*)\[(.*)\]/)
      let v = where[k]
      if (isArray(m) && m.length == 3) {
        const [_, col, op] = m // example: status[in] => ['status[in]', 'status', 'in']
        switch (op.toLowerCase()) {
          case 'in':
            builder.whereIn(col, v)
            break
          case '~':
          case '#':
            v = `%${v}%`
          case 'like':
            builder.where(col, 'like', v)
            break
          case '!':
          case 'not':
            builder.whereNot(col, v)
            break
          case '>':
          case '>=':
          case '<':
          case '<=':
            builder.where(col, op, v)
            break
          case '<>':
            builder.whereBetween(col, v)
            break
          case '><':
            builder.whereNotBetween(col, v)
            break
          default:
            break
        }
      } else {
        builder.where(k, v)
      }
    }
    // console.log({ sql: builder.toSQL().sql, bindings: builder.toSQL().bindings })
    return builder
  }

}