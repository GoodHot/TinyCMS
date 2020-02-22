import {request} from '@/utils/request'

/**
 * @param {user: String, password: String} parameter 
 */
export function searchTag (prefix) {
  return request({
    url: `/tag/search/${prefix}`,
    method: 'get'
  })
}