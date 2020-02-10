import { axios } from '@/utils/request'

export function uploadBase64 (parameter) {
  return axios({
    url: '/upload/base64',
    method: 'post',
    data: parameter
  })
}
