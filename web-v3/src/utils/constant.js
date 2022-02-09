import moment from 'moment'
moment.locale('zh-cn')

class Constant {
  constructor() {
    this.Map = {
      Type: {
        Admin: 1,
      }
    }
  }

  getMap(key) {
    return this.Map[key] || {}
  }

  getValueMap(key, stringify = false) {
    const _map = {}
    if (typeof this.Map[key] === 'undefined') {
      return _map
    }
    for (const k in this.Map[key]) {
      if (stringify) {
        _map[String(this.Map[key][k])] = String(k)
      } else {
        _map[this.Map[key][k]] = k
      }
    }
    return _map
  }

  /**
   * get map options for element format
   * @param {String} key
   * @return {Array}
   */
  getMapOptions(key, vIgnore = []) {
    const options = []
    for (const k in this.Map[key]) {
      if (vIgnore.indexOf(this.Map[key][k]) !== -1) {
        continue
      }
      options.push({
        value: this.Map[key][k],
        label: k
      })
    }
    return options
  }

  /**
   * get map options for element table column format
   * @param {String} key
   * @return {Array}
   */
  getTableFilters(key) {
    const filters = []
    for (const k in this.Map[key]) {
      filters.push({
        value: this.Map[key][k],
        text: k
      })
    }
    return filters
  }

  /**
   * return element picker options
   * @return {Object}
   */
  getDatePickerOptions() {
    return {
      shortcuts: [{
        text: 'yesterday',
        onClick(picker) {
          const end = moment().add(-1, 'days').endOf('day')
          const start = moment().add(-1, 'days').startOf('day')
          picker.$emit('pick', [start, end])
        }
      }, {
        text: 'Last 7 days',
        onClick(picker) {
          const end = moment().endOf('day')
          const start = moment().add(-7, 'days').startOf('day')
          picker.$emit('pick', [start, end])
        }
      }, {
        text: 'Last 15 days',
        onClick(picker) {
          const end = moment().endOf('day')
          const start = moment().add(-15, 'days').startOf('day')
          picker.$emit('pick', [start, end])
        }
      }, {
        text: 'Last 30 days',
        onClick(picker) {
          const end = moment().endOf('day')
          const start = moment().add(-30, 'days').startOf('day')
          picker.$emit('pick', [start, end])
        }
      }, {
        text: 'This week',
        onClick(picker) {
          const end = moment().weekday(6).endOf('day')
          const start = moment().weekday(0).startOf('day')
          picker.$emit('pick', [start, end])
        }
      }, {
        text: 'This month',
        onClick(picker) {
          const end = moment().endOf('month')
          const start = moment().startOf('month')
          picker.$emit('pick', [start, end])
        }
      }]
    }
  }
}

export default new Constant()
