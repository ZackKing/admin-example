import request from '~/utils/request'

export function getAccountList(params) {
  return request({
    url: '/user/list',
    method: 'get',
    params
  })
}

export function addAccount(data) {
  return request({
    url: '/user/add',
    method: 'post',
    data
  })
}

export function updateAccount(data) {
  return request({
    url: '/user/update',
    method: 'post',
    data
  })
}

export function setStatus(data) {
  return request({
    url: '/user/status',
    method: 'post',
    data
  })
}

export function setGroup(data) {
  return request({
    url: '/user/group',
    method: 'post',
    data
  })
}
