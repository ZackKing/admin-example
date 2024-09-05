declare var allTasks: { [key: string]: Task }

type KV = {
  [key: string | number]: any
}

type Int = number

type RouteConf = {
  method: 'get' | 'post' | 'put' | 'del' | 'all'
  path: string
  ctl: Router.IMiddleware
}

type RouterMap = {
  [key: string]: RouteConf[]
}

interface Conf {
  app: string
  env: string
  debug: boolean
  host: string
  port: number
  timezone: number
  key: string

  db: KV
  redis: KV
  jwt: { secret: string, options: KV }
  tasks: TaskConf[]
}

// for task
interface TaskConf {
  name: string
  cls: string
  active: boolean
  [key: string]: any
}

interface Task {
  conf: TaskConf
  name: string
  status: boolean
  lastError: string
  get conf()
  async run()
  async stop()
}
