import { usePermissionStore } from '@/store/permission'
import { ElMessage } from 'element-plus'

function checkPermission(el, binding) {
  const { value } = binding
  const store = usePermissionStore()
  const accessUrls = store.accessUrls

  if (value && value instanceof Array) {
    if (value.length > 0) {
      const permissionUrl = value

      const hasPermission = permissionUrl.every(url => {
        return accessUrls.includes(url)
      })

      // hidden
      // if (!hasPermission) {
      //   el.parentNode && el.parentNode.removeChild(el)
      // }
      // tips
      if (!hasPermission) {
        el.addEventListener('click', function (event) {
          event.stopPropagation()
          ElMessage({
            message: 'Access denied!',
            type: 'error',
            duration: 5 * 1000
          })
        }, true)
      }
    }
  } else {
    throw new Error(`need urls! Like v-permission="['/auth/menu','/auth/menu']"`)
  }
}

export default {
  inserted(el, binding) {
    checkPermission(el, binding)
  }
}
