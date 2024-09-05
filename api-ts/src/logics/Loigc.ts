import { LogicError } from '~/types/error'

export default class Logic {

  private static instances: { [key: string]: Logic } = {}

  static instance<T extends Logic>(this: new () => T): T {
    const className = this.name
    if (!Logic.instances[className]) {
      Logic.instances[className] = new this()
    }
    return Logic.instances[className] as T
  }

  /**
   * @param code 
   * @param msg 
   * @param data 
   * @throws LogicError
   */
  throw(code: number, msg: string = '', data: KV = {}): never {
    throw new LogicError(code, msg, data)
  }

}
