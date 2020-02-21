import {request} from '@/utils/request'

/**
 * @param {user: String, password: String} parameter 
 */
export function userLogin (parameter) {
  return request({
    url: '/auth/login',
    method: 'post',
    data: parameter
  })
}