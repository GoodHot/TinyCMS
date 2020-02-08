import { axios } from '@/utils/request'

/**
 * login func
 * parameter: {
 *     username: '',
 *     password: '',
 *     remember_me: true,
 *     captcha: '12345'
 * }
 * @param parameter
 * @returns {*}
 */
export function getChannelPage (parameter) {
  return axios({
    url: `/channel/page_${parameter.page}?parentId=${parameter.parentId}`,
    method: 'get'
  })
}

export function saveChannel (param) {
  return axios({
    url: `/channel`,
    method: 'post',
    data: param
  })
}

export function deleteChannel (id) {
  return axios({
    url: `/channel/${id}`,
    method: 'delete'
  })
}

export function getChannel (id) {
  return axios({
    url: `/channel/${id}`,
    method: 'get'
  })
}
