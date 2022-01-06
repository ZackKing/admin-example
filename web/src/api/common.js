import request from '@/utils/request'

export function getPermissionMenu() {
  return request({
    url: '/auth/menu',
    method: 'get'
  })
}

