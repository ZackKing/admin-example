import fs from 'fs'
import { log } from '~/services/log'
import { resolve } from 'path'

const env = process.env.NODE_ENV || 'dev'
const filePath = resolve(__dirname, `../configs/${env}.json`)
if (!fs.existsSync(filePath)) {
  log.error(`config file not found ! read ${filePath}`)
  process.exit(1)
}
const fileConf = JSON.parse(fs.readFileSync(filePath, 'utf8'))

const conf: Conf = {
  app: fileConf.app ?? 'admin-server',
  env,
  debug: fileConf.debug ?? false,
  host: fileConf.host ?? '',
  port: fileConf.port ?? 80,
  timezone: fileConf.timezone ?? 0,
  key: fileConf.app_key ?? '',

  db: fileConf.db,
  redis: fileConf.redis,
  jwt: fileConf.jwt,

  tasks: [
    // { name: 'name1', cls: 'class_name', active: true },
    // { name: 'name2', cls: 'class_name', active: true },
  ],
}
export default conf

export const constant: KV = {
  APP_NAME: 'admin-server',

  PWD_WRONG_LIMIT: 5,
}
