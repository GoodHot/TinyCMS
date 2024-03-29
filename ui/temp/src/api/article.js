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

export function articlePage (query) {
  let url = `/article/page_${query.page}?q=q`
  if (query && query.type) {
    url += `&type=${query.type}`
  }
  if (query && query.category) {
    url += `&category=${query.category}`
  }
  if (query && query.keyword) {
    url += `&keyword=${query.keyword}`
  }
  if (query && query.tag) {
    url += `&tag=${query.tag}`
  }
  return request({
    url: url,
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

export function deleteArticles (ids) {
  return request({
    url: '/article',
    method: 'delete',
    data: ids
  })
}