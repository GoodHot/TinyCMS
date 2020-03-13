import Vue from 'vue'
import App from './App.vue'
import router from './router'
import { BootstrapVue } from 'bootstrap-vue'
// import 'bootstrap/dist/css/bootstrap.css'
// import 'bootstrap-vue/dist/bootstrap-vue.css'
import '@/assets/css/dashmix.css'
// import '@/assets/css/xwork.css'
import '@/assets/css/style.scss'
import Components from './components'

Vue.config.productionTip = false
Vue.use(BootstrapVue)

for (const key in Components) {
  Vue.component(key, Components[key])
}

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
