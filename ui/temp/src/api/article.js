import {request} from '@/utils/request'

/**
 * @param {user: String, password: String} parameter 
 */
export function saveArticle (parameter) {
  return request({
    url: '/article',
    method: 'post',
    data: parameter
  })
}