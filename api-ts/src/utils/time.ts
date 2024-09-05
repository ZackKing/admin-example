
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

  static ds(date?: string | Date | dayjs.Dayjs, tz = 0) {
    return time.formatDate(date ?? new Date(), time.FORMAT.ds, tz)
  }

  static timeStamp(format = '', tz = 0) {
    dayjs.extend(utc)
    return dayjs().utcOffset(tz).format(format || time.FORMAT.default)
  }

  static unix() {
    return dayjs().unix()
  }

  static weekFirstDay(date?: string | number | Date | dayjs.Dayjs | null | undefined, format = time.FORMAT.ds, tz = 0) {
    dayjs.extend(utc)
    dayjs.extend(isoWeek)
    return dayjs(date).utcOffset(tz).isoWeekday(1).format(format)
  }

  static lastFirday(date?: string | number | Date | dayjs.Dayjs | null | undefined, fmt = time.FORMAT.ds, tz = 0) {
    const day = dayjs(date)
    const gap = (day.day() + 7 - 5) % 7
    const lastFriday = day.add(-1 * gap, 'day')
    dayjs.extend(utc)
    return lastFriday.utcOffset(tz).format(fmt)
  }

  static formatDate(date?: string | Date| dayjs.Dayjs, format = time.FORMAT.default, tz = 0) {
    dayjs.extend(utc)
    return dayjs(date).utcOffset(tz).format(format)
  }

  static getDayEndOfMonth(date?: string, format = time.FORMAT.date) {
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