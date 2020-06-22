// Libraries
import Vue from 'vue'
import VueRouter from 'vue-router'

// Main application window
import App from '@/app.vue'

// Views
import Dashboard from '@/views/dashboard.vue'
import Login from '@/views/login.vue'

// Components
// import Flex from '@/components/flex.vue'

// Load Vue.js libraries
Vue.use(VueRouter)

Vue.config.productionTip = false

const routes = [
  // Redirects
  // { path: '*', redirect: '/' },

  // Navigation entries
  { path: '/', name: 'dashboard', title: 'Dashboard', component: Dashboard },

  // Hidden routes
  { path: '/', name: 'login', title: 'Login', hidden: true, component: Login },
  // ...
].map(r => {
  // https://router.vuejs.org/api/#routes
  const reserved = [ 'path', 'component', 'name', 'components', 'redirect', 'props', 'alias', 'children', 'beforeEnter', 'meta', 'caseSensitive', 'pathToRegexpOptions' ]
  const route = { meta: { } }

  for (const p in r) {
    if (reserved.includes(p)) {
      route[p] = r[p]
    } else {
      route.meta[p] = r[p]
    }
  }

  return route
})

const router = new VueRouter({
  mode: 'history',
  root: '/',
  routes: routes,
  linkExactActiveClass: 'active',
  scrollBehavior: (to, from, savedPosition) => ({ x: 0, y: 0 }),
})

router.beforeEach((to, from, next) => {
  document.title = `${process.env.NAME} > ${to.meta.title}`
  next()
})

new Vue({
  render: h => h(App),
  router: router,
}).$mount('#app')
