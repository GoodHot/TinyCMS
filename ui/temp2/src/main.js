import Vue from 'vue'
import App from './App.vue'
import router from './router'
import { BootstrapVue } from 'bootstrap-vue'
// import 'bootstrap/dist/css/bootstrap.css'
// import 'bootstrap-vue/dist/bootstrap-vue.css'
import '@/assets/css/dashmix.css'
// import '@/assets/css/xwork.css'
import '@/assets/css/style.scss'
import { config } from '@/utils/request'
import Components from './components'
import Storage from 'vue-ls'

Vue.config.productionTip = false
Vue.use(BootstrapVue)

Vue.use(Storage)

for (const key in Components) {
  Vue.component(key, Components[key])
}

Vue.prototype.assetsURL = function (src) {
  if (!src) {
    return 'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd66pwrkqij30hs0hsgmb.jpg'
  }
  if (src.startsWith('http')) {
    return src
  }
  return config.assetsURL + src
}

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
