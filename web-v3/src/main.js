
import '@/styles/index.scss' // global css

import { createApp } from 'vue'
import App from '@/App'
const app = createApp(App)

import { createPinia } from 'pinia'
app.use(createPinia())
import store from '@/store'
app.config.globalProperties.$store = store.install()

// router
import router from '@/router'
app.use(router)

// element
import ElementPlus from 'element-plus'
import ElementPlusZhCn from 'element-plus/dist/locale/zh-cn.mjs'
import 'element-plus/dist/index.css'
app.use(ElementPlus, { locale: ElementPlusZhCn, size: 'small' })

// icon
import svgIcon from '@/components/SvgIcon.vue'
app.component('SvgIcon', svgIcon)

import '@/permission' // permission control
import permission from '@/directive/permission/permission.js'
app.directive('permission', permission)
app.config.productionTip = false

app.mount('#app')
