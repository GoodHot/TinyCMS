import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './assets/css/codebase.css'
import './assets/css/style.css'
import './assets/js/bootstrap'
import Components from './components'
import JQuery from 'jquery'
import dialog from './assets/js/dialog'
import Storage from 'vue-ls'

Vue.config.productionTip = false

for (const key in Components) {
  Vue.component(key, Components[key])
}

Vue.prototype.$ = JQuery
Vue.prototype.$dialog = dialog

// const options = {
//   namespace: 'vuejs__', // key prefix
//   name: 'ls', // name variable Vue.[ls] or this.[$ls],
//   storage: 'local', // storage name session, local, memory
// }

Vue.use(Storage)

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
