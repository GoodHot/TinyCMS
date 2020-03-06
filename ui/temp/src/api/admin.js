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

export function getAdmin (id) {
  return request({
    url: `/admin/${id}`,
    method: 'get'
  })
}

export function deleteAdmin (id) {
  return request({
    url: `/admin/${id}`,
    method: 'delete'
  })
}