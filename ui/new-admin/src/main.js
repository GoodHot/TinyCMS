import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './assets/css/codebase.css'
import './assets/css/style.css'
import Components from './components'

Vue.config.productionTip = false

for (const key in Components) {
  Vue.component(key, Components[key])
}

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
