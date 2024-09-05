import pino from 'pino'
import time from '~/utils/time'
import fs from 'fs'
import { resolve } from 'path'

const logsPath = resolve(__dirname, '../../storages/logs')

const formatters = {
  level(label: string) {
    return { level: label }
  },
  log(info: any) {
    info.ts = time.timeStamp()
    return info
  },
}

const multistream = pino.multistream([
  { stream: process.stdout },
  { stream: fs.createWriteStream(`${logsPath}/all.log`) },
])

const log = pino({ level: 'info', formatters }, multistream)
const logger = log
const tLog = {
  error(tag: string, ...data: any) {
    log.error({ tag, data })
  },
  info(tag: string, ...data: any) {
    log.info({ tag, data })
  }
}

export { log, logger, tLog }