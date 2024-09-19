import comm from '~/utils/comm'

export class LogicError extends Error {
  code: number
  data: any = null
  constructor(code: number, msg?: string, data?: any, params?: any[]) {
    let message = msg || errMsg[code] || 'Unknown error'
    if (params) {
      message = comm.vsprintf(message, params)
    }
    super(message)
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

  10001: 'Token Not Found!',
  10002: 'Invalid Token',
  10003: 'Account not found / Password error!',
  10004: 'Old password error! / password change error!',
  10005: 'Account is disabled ! please contact administrator !',
  10006: 'Add account error! ',
  10007: 'Account user exist!',
  10008: 'Account password wrong! will ban with %s chance left !',
  10009: 'Account password retry too much ! set disabled !',

}
