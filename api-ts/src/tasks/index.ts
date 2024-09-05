import conf from '~/conf'
import { log } from '~/services/log'
import schedule from 'node-schedule'
import time from '~/utils/time'
import { clone } from 'lodash'
import reporter from '~/services/reporter'

export default class Task {

  private static _instance: Task

  static instance(): Task {
    if (!Task._instance) {
      Task._instance = new Task()
    }
    return Task._instance
  }

  activedNames: string[] = []

  clsMap: KV = {
    // 'name': Class,
  }

  async init() {
    globalThis.allTasks = {}
    conf.tasks.map(async cfg => {
      if (!cfg.active) {
        return
      }
      if (typeof this.clsMap[cfg.cls] == 'undefined') {
        log.error(`unknow tasks class ${JSON.stringify(cfg)}`)
        return
      }
      const Cls = this.clsMap[cfg.cls]
      globalThis.allTasks[cfg.name] = new Cls(cfg)
      globalThis.allTasks[cfg.name].run()
      this.activedNames.push(cfg.name)
      log.info(`task ${cfg.name} has start at ${time.timeStamp()}`)
    })
    await this.holder()
  }

  async holder() {
    schedule.scheduleJob('*/1 * * * *', async () => {
      this.activedNames.map(async name => {
        const task = globalThis.allTasks[name]
        if (task.status) {
          return
        }
        let errMsg = `task ${name} has stop ! error: ${task.lastError}`
        log.error(errMsg)
        reporter.error(errMsg)
        try {
          const cfg = clone(task.conf)
          const Cls = this.clsMap[cfg.cls]
          const newTask = new Cls(cfg)
          delete globalThis.allTasks[name]
          globalThis.allTasks[name] = newTask
          globalThis.allTasks[name].run()
        } catch (error: any) {
          errMsg = `task ${name} has restart error ! error: ${error.message}`
          log.error(errMsg)
          reporter.error(errMsg)
        }
      })
    })
  }


}

