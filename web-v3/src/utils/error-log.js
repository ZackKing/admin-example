import Vue from 'vue'
import { isString, isArray } from '@/utils/validate'
import settings from '@/settings'

// you can set in settings.js
// errorLog:'prod' | ['prod', 'dev']
const { errorLog: needErrorLog } = settings

function checkNeed() {
  const env = process.env.NODE_ENV
  if (isString(needErrorLog)) {
    return env === needErrorLog
  }
  if (isArray(needErrorLog)) {
    return needErrorLog.includes(env)
  }
  return false
}

if (checkNeed()) {
  Vue.config.errorHandler = function (err, vm, info, a) {
    // just a hack
    Vue.nextTick(() => {
      console.error(err, info)
    })
  }
}
