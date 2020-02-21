import Vue from 'vue'
import App from './App.vue'
import router from './router'
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
import Components from './components'
import Storage from 'vue-ls'
import dialog from './assets/js/dialog'

// import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import './assets/css/codebase.css'
import './assets/css/style.css'

Vue.config.productionTip = false

for (const key in Components) {
  Vue.component(key, Components[key])
}

// Install BootstrapVue
Vue.use(BootstrapVue)
// Optionally install the BootstrapVue icon components plugin
Vue.use(IconsPlugin)
// vue.ls
Vue.use(Storage)
// dialog
Vue.prototype.$dialog = dialog

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
