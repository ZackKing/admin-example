import request from '~/utils/request'

export function getMenuList(params) {
  return request({
    url: '/menu',
    method: 'get',
    params
  })
}

export function getMenuInfo(params) {
  return request({
    url: '/menu/info',
    method: 'get',
    params
  })
}

export function setGroup(data) {
  return request({
    url: '/menu/group',
    method: 'post',
    data
  })
}

