import { tLog } from './log'


export default new class Reporter {

  async info(msg: any) {
    tLog.info('reporter.info', msg)
  }

  async error(msg: any) {
    tLog.error('reporter.error', msg)
  }

}