
import _ from 'lodash'
import moment from 'moment'

export function numFormatter(row, col, cv, idx) {
  const num = +cv
  if (_.isNil(num) || _.isNaN(num)) {
    return 0
  }
  return num.toLocaleString()
}

export function timeFormatter(row, col, cv, idx) {
  if (_.isNil(cv)) {
    return '0s'
  }
  let str = ''
  if (cv > 3600) {
    str += Math.floor(cv / 3600) + 'h' + Math.floor((cv % 3600) / 60) + 'm'
  } else if (cv > 60) {
    str += Math.floor(cv / 60) + 'm'
  }
  str += (cv % 3600) % 60 + 's'
  return str
}

export function dateFormatter(row, col, cv, idx) {
  let str = '-'
  if (cv > 0) {
    str = moment.unix(cv).format('YYYY/MM/DD HH:mm:ss')
  }
  return str
}
