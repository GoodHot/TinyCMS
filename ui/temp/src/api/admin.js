import {request} from '@/utils/request'

/**
 * @param {user: String, password: String} parameter 
 */
export function getAllAdmin () {
  return request({
    url: '/admin/all',
    method: 'get'
  })
}

export function saveAdmin (param) {
  return request({
    url: '/admin',
    method: 'post',
    data: param
  })
}