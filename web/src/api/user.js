import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/login',
    method: 'post',
    data
  })
}

export function getInfo() {
  return request({
    url: '/user/self',
    method: 'get'
  })
}

export function logout() {
  return request({
    url: '/logout',
    method: 'post'
  })
}

export function getCaptcha() {
  return request({
    url: '/login/captcha',
    method: 'get'
  })
}

export function editInfo(data) {
  return request({
    url: '/user/self',
    method: 'post',
    data
  })
}

export function editPassword(data) {
  return request({
    url: '/auth/password',
    method: 'post',
    data
  })
}

export function sendMailCode(data) {
  return request({
    url: '/login/sendMailCode',
    method: 'post',
    data
  })
}
