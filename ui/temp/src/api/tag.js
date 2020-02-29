import {request} from '@/utils/request'

export function searchTag (prefix) {
  return request({
    url: `/tag/search/${prefix}`,
    method: 'get'
  })
}

export function hotTags () {
  return request({
    url: '/tag/hot',
    method: 'get'
  })
}