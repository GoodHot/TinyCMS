import {request} from '@/utils/request'

export function categoryTree () {
  return request({
    url: '/category/tree',
    method: 'get'
  })
}

export function saveCategory (parameter) {
  return request({
    url: '/category',
    method: 'post',
    data: parameter
  })
}

export function getCategory (id) {
  return request({
    url: `/category/${id}`,
    method: 'get'
  })
}

export function getCategoryByPage (page, keyword) {
  return request({
    url: `/category/page_${page}?keyword=${keyword}`,
    method: 'get'
  })
}