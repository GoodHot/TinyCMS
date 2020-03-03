import {request} from '@/utils/request'

export function allPermission () {
  return request({
    url: '/role/permission',
    method: 'get'
  })
}

export function saveRole (param) {
  return request({
    url: '/role',
    method: 'post',
    data: param
  })
}

export function pageRole (page) {
  return request({
    url: `/role/${page}`,
    method: 'get'
  })
}
