import permission from './permission/permission.js'
// import watermark from './watermark.js'

export default {
  install(app) {
    app.directive('permission', permission)
    // app.directive('watermark', watermark)
  }
}