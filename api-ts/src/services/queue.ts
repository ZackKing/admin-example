import { Redis } from 'ioredis'
import redis from './redis'

class Queue {

  connKey: string
  queueName: string
  private _key: string
  private _conn: Redis
  private _connBlock: Redis

  constructor(key: string, ck: string) {
    this.connKey = ck
    this.queueName = key
    this._key = `queue:${key}`
    this._conn = redis.newConn(ck)
  }

  async push(item: any): Promise<number> {
    return await this._conn.rpush(this._key, JSON.stringify(item))
  }

  async pop(block: boolean = false): Promise<any | null> {
    let itemStr = null
    if (block) { // TODO: just for 1 conn blpop, need change to pool
      if (!this._connBlock) this._connBlock = redis.newConn(this.connKey)
      const res = await this._connBlock.blpop(this._key, 0)
      itemStr = res ? res[1] : null
    } else {
      itemStr = await this._conn.lpop(this._key) 
    }
    return itemStr ? JSON.parse(itemStr) : null
  }

  async size(): Promise<number> {
    return await this._conn.llen(this._key);
  }

  async clear(): Promise<void> {
    await this._conn.del(this._key);
  }

}

export const queues = {
  example: new Queue('name', 'queue'),
}
