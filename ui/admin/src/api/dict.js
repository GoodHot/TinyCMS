import { axios } from '@/utils/request'

export function getAllDict () {
  return axios({
    url: `/dict/all`,
    method: 'get'
  })
}

export function editDict (param) {
  return axios({
    url: `/dict`,
    method: 'put',
    data: param
  })
}
