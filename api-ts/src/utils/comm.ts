import crypto from 'crypto'

interface RandomItem extends KV {
  odds: number
}

export default class comm {

  private static readonly BASE_CHARS: string = 'abcdefghijklm0123456789nopqrstuvwxyz'
  private static readonly BASE_NUM: number = 20240828

  static md516(str: string) {
    const md5str = this.md5(str)
    return md5str.slice(8, 24)
  }

  static md5(str: string) {
    const generator = crypto.createHash('md5')
    generator.update(str)
    return generator.digest('hex')
  }

  static jsonDecode(str: string, defaultVal: any = ''): any {
    try {
      return JSON.parse(str)
    } catch (error) {
      return defaultVal
    }
  }

  static jsonEncode(data: any): string {
    return JSON.stringify(data)
  }

  static radnomHit(odds: number, loop: Int = 1): Int {
    let hit = 0
    for (let i = 0; i < loop; i++) {
      if (Math.random() < odds) {
        hit++
      }
    }
    return hit
  }

  static randomInt(min: Int = 0, max: Int = 100): Int {
    // make sure is int
    min = Math.ceil(min)
    max = Math.round(max)
    return Math.floor(Math.random() * (max - min + 1)) + min
  }

  static randomItem(items: RandomItem[]): KV | false {
    if (items.length == 0) {
      return false
    } else if (items.length == 1) {
      return items[0]
    }

    const random = Math.random()
    let odds = 0
    for (const item of items) {
      odds += item.odds
      if (odds <= random) {
        return item
      }
    }
    return items[items.length - 1]
  }

  static encodeID(id: number): string {
    if (id < 1) {
      return '0'
    }
    let str = ''
    id += this.BASE_NUM
    const base = this.BASE_CHARS.length
    while (id > 0) {
      const remainder = id % base
      str = this.BASE_CHARS.charAt(remainder) + str
      id = Math.floor(id / base)
    }
    return str
  }

  static decodeID(str: string): number | false {
    if (str.length < 1 || str == '') {
      return 0
    }
    str = str.toLowerCase()
    const base = this.BASE_CHARS.length
    let id = 0
    const len = str.length
    for (let i = 0; i < len; i++) {
      const charValue = this.BASE_CHARS.indexOf(str.charAt(i))
      if (charValue === -1) {
        return false
      }
      id = id * base + charValue
    }
    if (id <= this.BASE_NUM) {
      return false
    }
    return id - this.BASE_NUM
  }

}