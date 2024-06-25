import request from '~/utils/request'

export function getGroupList(params) {
  return request({
    url: '/group',
    method: 'get',
    params
  })
}

export function addGroup(data) {
  return request({
    url: '/group',
    method: 'post',
    data
  })
}

export function updateGroup(data) {
  return request({
    url: '/group/edit',
    method: 'post',
    data
  })
}

export function setMenu(data) {
  return request({
    url: '/group/menu',
    method: 'post',
    data
  })
}

export function setUser(data) {
  return request({
    url: '/group/user',
    method: 'post',
    data
  })
}
