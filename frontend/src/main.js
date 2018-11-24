// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import Vuex from 'vuex'
import Buefy from 'buefy'
import 'buefy/dist/buefy.css'
import VueResource from 'vue-resource'
import VueParticles from 'vue-particles'
import axios from 'axios'
import VueAxios from 'vue-axios'

Vue.config.productionTip = false
Vue.use(Buefy)
Vue.use(VueResource)
Vue.use(Vuex)
Vue.use(VueParticles)
Vue.use(VueAxios, axios)

const store = new Vuex.Store({
  state: {
    server: "http://192.33.206.40:4200",
    backgroundImagePath: require('./assets/background.jpg'),
    userIsLoggedIn: false,
    jwt: "",
    roleUser: false,
    roleAdmin: false,
  },
  mutations: { // possible transformations
    switch_url(state, url) {
      state.url_coreCA = url
    },
    switch_background (state, image) {
      state.backgroundImagePath = image
    },
    setLoggedIn (state, isLoggedIn) {
      state.userIsLoggedIn = isLoggedIn
    },
    setToken (state, token) {
      state.jwt = token
    },
    setRoleUser (state, role) {
      state.roleUser = role
    },
    setRoleAdmin (state, role) {
      state.roleAdmin = role
    }
  }
});

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
});