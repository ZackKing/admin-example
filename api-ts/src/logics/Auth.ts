import conf from '~/conf'
import Logic from './Loigc'
import jwt from 'jsonwebtoken'
import { LogicError } from '~/types/error'

export default class Auth extends Logic {

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

}