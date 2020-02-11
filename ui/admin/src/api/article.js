import { axios } from '@/utils/request'

export function publishArticle (parameter) {
  return axios({
    url: '/article/publish',
    method: 'post',
    data: parameter
  })
}

export function getArticlePage (parameter) {
  return axios({
    url: `/article/page_${parameter.page}`,
    method: 'get'
  })
}

export function deleteArticle (id) {
  return axios({
    url: `/article/${id}`,
    method: 'delete'
  })
}

export function getArticle (id) {
  return axios({
    url: `/article/${id}`,
    method: 'get'
  })
}
