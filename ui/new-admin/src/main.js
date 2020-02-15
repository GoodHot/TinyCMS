import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './assets/css/codebase.css'
import './assets/css/style.css'
import './assets/js/bootstrap'
import Components from './components'
import JQuery from 'jquery'

Vue.config.productionTip = false

for (const key in Components) {
  Vue.component(key, Components[key])
}

Vue.prototype.$ = JQuery

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
