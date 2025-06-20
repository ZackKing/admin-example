
import dayjs from 'dayjs'
import utc from 'dayjs/plugin/utc'
import isoWeek from 'dayjs/plugin/isoWeek'

export default class time {

  static DS = 'YYYYMMDD'
  static FORMAT = {
    default: 'YYYY-MM-DD HH:mm:ss',
    ds: 'YYYYMMDD',
    date: 'YYYY.MM-DD'
  }

  static ds(date: dayjs.ConfigType = new Date(), tz = 0) {
    return time.formatDate(date, time.FORMAT.ds, tz)
  }

  static timeStamp(format = '', tz = 0) {
    dayjs.extend(utc)
    return dayjs().utcOffset(tz).format(format || time.FORMAT.default)
  }

  static unix() {
    return dayjs().unix()
  }

  static weekFirstDay(date: dayjs.ConfigType = new Date(), format = time.FORMAT.ds, tz = 0) {
    dayjs.extend(utc)
    dayjs.extend(isoWeek)
    return dayjs(date).utcOffset(tz).isoWeekday(1).format(format)
  }

  static lastFirday(date: dayjs.ConfigType = new Date(), fmt = time.FORMAT.ds, tz = 0) {
    const day = dayjs(date)
    const gap = (day.day() + 7 - 5) % 7
    const lastFriday = day.add(-1 * gap, 'day')
    dayjs.extend(utc)
    return lastFriday.utcOffset(tz).format(fmt)
  }

  static formatDate(date: dayjs.ConfigType = new Date(), format = time.FORMAT.default, tz = 0) {
    dayjs.extend(utc)
    return dayjs(date).utcOffset(tz).format(format)
  }

  static getDayEndOfMonth(date: string | Date = new Date(), format = time.FORMAT.date) {
    return dayjs(date).endOf('month').format(format)
  }

  static delay(ts: number) {
    return new Promise<void>((reslove) => {
      setTimeout(() => {
        reslove()
      }, ts)
    })
  }
}