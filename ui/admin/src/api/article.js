import { axios } from '@/utils/request'

export function publishArticle (parameter) {
  return axios({
    url: '/article/publish',
    method: 'post',
    data: parameter
  })
}
