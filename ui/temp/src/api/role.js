import {request} from '@/utils/request'

export function allPermission () {
  return request({
    url: '/role/permission',
    method: 'get'
  })
}
