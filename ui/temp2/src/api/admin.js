import {request} from '@/utils/request'

export function getAllAdmin () {
  return request({
    url: '/admin/all',
    method: 'get'
  })
}
