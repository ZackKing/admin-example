import Vue from 'vue'

import 'normalize.css/normalize.css' // A modern alternative to CSS resets
import '@/styles/index.scss' // global css
import '@/icons'

import ElementUI from 'element-ui'
// import 'element-ui/lib/theme-chalk/index.css'
import '@/styles/element-variables.scss'
import locale from 'element-ui/lib/locale/lang/en' // lang i18n
Vue.use(ElementUI, { locale, size: 'small' })

// register global utility filters
import * as filters from '@/filters'
Object.keys(filters).forEach(key => {
  Vue.filter(key, filters[key])
})

import ECharts from 'vue-echarts'
import 'echarts/lib/chart/bar'
import 'echarts/lib/component/tooltip'
Vue.component('v-chart', ECharts)

import '@/permission' // permission control
import permission from '@/directive/permission/index.js'
Vue.directive('permission', permission)
Vue.config.productionTip = false

import App from '@/App'
import store from '@/store'
import router from '@/router'

new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App)
})
