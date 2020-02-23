import axios from 'axios'
import Vue from 'vue'

const ACCESS_TOKEN = 'ACCESS-TOKEN'

const requestConfig = {
  baseURL: 'http://localhost:9000/admin',
  assetsURL: 'http://localhost:9000'
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
      // notification.error({
      //   message: 'Forbidden',
      //   description: data.msg
      // })
      alert(data.msg)
    }
    if (error.response.status === 401 && !(data.result && data.result.isLogin)) {
      const msg = data.msg || 'Authorization verification failed'
      // notification.error({
      //   message: 'Unauthorized',
      //   description: msg
      // })
      window.location = "/#/auth/login"
      alert(msg)
      if (token) {
        // store.dispatch('Logout').then(() => {
        //   setTimeout(() => {
        //     window.location.reload()
        //   }, 1500)
        // })
      }
    }
  }
  return Promise.reject(error)
}

// request interceptor
service.interceptors.request.use(config => {
  const token = Vue.ls.get(ACCESS_TOKEN)
  if (token) {
    config.headers[ACCESS_TOKEN] = token // 让每个请求携带自定义 token 请根据实际情况自行修改
  }
  return config
}, err)

// response interceptor
service.interceptors.response.use((response) => {
  const data = response.data
  if (data.status !== 200) {
    const msg = data.msg || '请求失败:\r\n' + response.config.url
    // notification.error({
    //   message: 'Server:' + data.status,
    //   description: msg
    // })
    alert(msg)
    return Promise.reject(data)
  }
  return data.data
}, err)

export {
  service as request,
  ACCESS_TOKEN,
  requestConfig as config
}