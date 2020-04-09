import {request} from '@/utils/request'

export function getTagByPage (page, keyword) {
  return request({
    url: `/tag/page_${page}?keyword=${keyword}`,
    method: 'get'
  })
}

export function saveTag (parameter) {
  return request({
    url: '/tag',
    method: 'post',
    data: parameter
  })
}

export function getTag (id) {
  return request({
    url: `/tag/${id}`,
    method: 'get'
  })
}