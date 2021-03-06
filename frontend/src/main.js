// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import VueClipboard from 'vue-clipboard2'
import VueMathjax from 'vue-mathjax'
import axios from 'axios'

Vue.use(VueClipboard)
Vue.use(VueMathjax)

Vue.config.productionTip = false
axios.defaults.baseURL = process.env.BACKEND_URL;

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
