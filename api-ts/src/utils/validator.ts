
import { isString, isArray, isPlainObject, findIndex } from 'lodash'

export interface ValidatorRule {
  key: string
  type: string
  required?: boolean
  values?: any[]
  ignore?: any[]
  default?: any
  rename?: string
  min?: number
  max?: number
  [key: string | number]: any
}

export default class Validator {

  static check(data: any, rules: ValidatorRule[]) {
    const safeData: KV = {}

    rules.map(rule => {

      let value = data[rule.key]

      if (typeof rule.ignore != 'undefined' && findIndex(rule.ignore, value) != -1) {
        return
      }

      if (rule.required) {
        _assert(typeof value != 'undefined', `'${rule.key}' is required!`)
      }

      if (typeof value != 'undefined') {
        switch (rule.type) {
          case 'int':
            value = +value
            _assert(Number.isInteger(value), `'${rule.key}' should be int`)
            if (rule.min) _assert(value >= rule.min, `'${rule.key}' value should >= ${rule.min}`)
            if (rule.max) _assert(value <= rule.max, `'${rule.key}' value should <= ${rule.max}`)
            break

          case 'number':
            value = +value
            _assert(!isNaN(value), `'${rule.key}' should be number`)
            if (rule.min) _assert(value >= rule.min, `'${rule.key}' value should >= ${rule.min}`)
            if (rule.max) _assert(value <= rule.max, `'${rule.key}' value should <= ${rule.max}`)
            break

          case 'string':
            _assert(isString(value), `'${rule.key}' should be string`)
            if (rule.min) _assert(value.length >= rule.min, `'${rule.key}' length should >= ${rule.min}`)
            if (rule.max) _assert(value.length <= rule.max, `'${rule.key}' length should <= ${rule.max}`)
            break

          case 'array':
            _assert(isArray(value), `'${rule.key}' should be array`)
            break

          case 'boolean':
            if (value === 'true') value = true
            if (value === 'false') value = false
            _assert(typeof value == 'boolean', `'${rule.key}' should be boolean`)
            break

          case 'enum':
            value = `${value}`
            if (!rule.values) {
              rule.values = []
            }
            _assert(rule.values.indexOf(value) !== -1, `'${rule.key}' value should in [${rule.values.join(',')}]`)
            break

          case 'json':
            if (value.length == 0) {
              value = null
              break
            }
            try {
              value = JSON.parse(value)
            } catch (e) {
              _assert(false, `'${rule.key}' should be a json format string`)
            }
            break

          case 'object':
            _assert(isPlainObject(value), `'${rule.key}' should be plain object`)
            break

          default:
            break
        }
      } else {
        if (typeof rule.default == 'undefined') {
          return
        }
        value = rule.default
      }
      rule.rename ? safeData[rule.rename] = value : safeData[rule.key] = value
    })

    return safeData
  }

}

function _assert(ok: boolean, msg = '') {
  if (!ok) {
    throw new Error(msg)
  }
}
