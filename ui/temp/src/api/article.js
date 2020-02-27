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

export function articlePage (page) {
  return request({
    url: `/article/page_${page}`,
    method: 'get'
  })
}

export function getArticleById (id) {
  return request({
    url: `/article/${id}`,
    method: 'get'
  })
}

export function getArticleCount () {
  return request({
    url: '/article/count',
    method: 'get'
  })
}