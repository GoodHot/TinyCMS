import Vue from 'vue'
import axios from 'axios'
import store from '@/store'
import notification from 'ant-design-vue/es/notification'
import { VueAxios } from './axios'
import { ACCESS_TOKEN } from '@/store/mutation-types'

const requestConfig = {
  // baseURL: 'http://dev.jiemen.fun/api/admin',
  baseURL: 'http://jiemen.fun/api/admin',
  assetsURL: 'http://assets0.jiemen.fun'
}

// 创建 axios 实例
const service = axios.create({
  baseURL: requestConfig.baseURL, // api base_url
  timeout: 20000 // 请求超时时间
})

const err = (error) => {
  if (error.response) {
    const data = error.response.data
    const token = Vue.ls.get(ACCESS_TOKEN)
    if (error.response.status === 403) {
      notification.error({
        message: 'Forbidden',
        description: data.msg
      })
    }
    if (error.response.status === 401 && !(data.result && data.result.isLogin)) {
      const msg = data.msg || 'Authorization verification failed'
      notification.error({
        message: 'Unauthorized',
        description: msg
      })
      if (token) {
        store.dispatch('Logout').then(() => {
          setTimeout(() => {
            window.location.reload()
          }, 1500)
        })
      }
    }
  }
  return Promise.reject(error)
}

// request interceptor
service.interceptors.request.use(config => {
  const token = Vue.ls.get(ACCESS_TOKEN)
  if (token) {
    config.headers['Access-Token'] = token // 让每个请求携带自定义 token 请根据实际情况自行修改
  }
  return config
}, err)

// response interceptor
service.interceptors.response.use((response) => {
  const data = response.data
  if (data.status !== 200) {
    const msg = data.msg || '请求失败:\r\n' + response.config.url
    notification.error({
      message: 'Server:' + data.status,
      description: msg
    })
    return Promise.reject(data)
  }
  return data.data
}, err)

const installer = {
  vm: {},
  install (Vue) {
    Vue.use(VueAxios, service)
  }
}

export {
  installer as VueAxios,
  service as axios,
  requestConfig as config
}
