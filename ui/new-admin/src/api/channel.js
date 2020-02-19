import {request} from '@/utils/request'

export function channelTree () {
  return request({
    url: '/channel/tree',
    method: 'get'
  })
}