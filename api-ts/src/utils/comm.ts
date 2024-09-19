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

  static randomStr(len = 6, base: string = ''): string
  {
    if (!base) {
      base = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789'
    }
    let str = ''
    for (let i = 0; i < len; i++) {
      str += base[comm.randomInt(0, base.length - 1)]
    }
    return str
  }

  static randomHit(odds: number, loop: Int = 1): Int {
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

  static vsprintf(format: string, args: any[]): string {
    let i = 0
    return format.replace(/%(\d*\$)?([sdif])/g, (sub, pos, spec) => {
      let j = i++
      if (pos) { // Adjust for 1-based pos spec like %1$s, %2$d, etc.
        j = parseInt(pos) - 1
      }
      const arg = args[j]
      switch (spec) {
        case 's': // String
          return String(arg)
        case 'd': // Int
        case 'i': // Int
          return parseInt(arg, 10).toString()
        case 'f': // Float
          return parseFloat(arg).toString()
        default:
          return sub
      }
    })
  }

  static genTree(list: any[] = [], key: string = 'id', pkey: string = 'pid', skey: string = 'sub'): any[] {
    const tree: any[] = []
    const items: KV = {}
    list.map(v => items[v[key]] = v)

    for (const k in items) {
      if (items.hasOwnProperty(k)) {
        const item = items[k]
        if (items[item[pkey]]) {
          if (!items[item[pkey]][skey]) {
            items[item[pkey]][skey] = []
          }
          items[item[pkey]][skey].push(item)
        } else {
          tree.push(item)
        }
      }
    }
    return tree
  }

}