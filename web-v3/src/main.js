
import '@/styles/index.scss' // global css

import { createApp } from 'vue'
import App from '@/App'
const app = createApp(App)

import router from '@/router'
app.use(router)

import store from '@/store'
app.use(store)

import svgIcon from '@/components/SvgIcon.vue'
app.component('SvgIcon', svgIcon)

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
app.use(ElementPlus)

import '@/permission' // permission control
import permission from '@/directive/permission/permission.js'
app.directive('permission', permission)
app.config.productionTip = false

app.mount('#app')
