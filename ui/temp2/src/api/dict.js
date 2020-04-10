import {request} from '@/utils/request'

export function getAllDict () {
  return request({
    url: '/dict',
    method: 'get'
  })
}

export function saveDict (param) {
  return request({
    url: '/dict',
    method: 'put',
    data: param
  })
}