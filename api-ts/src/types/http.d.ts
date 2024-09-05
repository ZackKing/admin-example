
interface ReqPage {
  page: Int
  size: Int
}

interface ResData {
  code: Int,
  msg: string,
  data: T | null,
  ts: Int,
}
