
export class LogicError extends Error {
  code: number
  data: any = null
  constructor(code: number, msg?: string, data?: any) {
    super(msg || errMsg[code] || 'Unknown error')
    this.code = code
    this.data = data || null
  }
}

export const errMsg: { [code: number]: string } = {
  '-1': 'Unknown',
  0: 'OK',
  1: 'Error',
  2: 'Validator error',
  3: 'Locked',

  100: 'Token not found',
  101: 'Invalid Token',

}
