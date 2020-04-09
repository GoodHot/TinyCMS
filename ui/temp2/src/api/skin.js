import {request} from '@/utils/request'

export function getAllSkin () {
  return request({
    url: '/skin',
    method: 'get'
  })
}