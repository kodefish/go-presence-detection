import Vue from 'vue'
import Router from 'vue-router'
import { Toast } from 'buefy/dist/components/toast'
import SignupComponent from "../views/signup.vue"
import LoginComponent from "../views/login.vue"
import InsecureComponent from "../views/insecure.vue"
import SecureComponent from "../views/secure.vue"

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      redirect: {
          name: "login"
      }
    },
    {
      path: "/signup",
      name: "signup",
      component: SignupComponent
    },
    {
        path: "/login",
        name: "login",
        component: LoginComponent
    },
    {
      path: "/insecure",
      name: "insecure",
      component: InsecureComponent
    },
    {
      path: "/secure",
      name: "secure",
      component: SecureComponent
    }
  ]
})
