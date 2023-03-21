
import '@/styles/index.scss' // global css

import { createApp } from 'vue'
import App from '@/App'
const app = createApp(App)

import { createPinia } from 'pinia'
app.use(createPinia())
import store from '@/store'
app.config.globalProperties.$store = store.install()
// app.use(store)

import router from '@/router'
app.use(router)

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
