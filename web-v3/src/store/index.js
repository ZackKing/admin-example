import { useAppStore } from '../store/app'
import { usePermissionStore } from '../store/permission'
import { useSettingStore } from '../store/settings'
import { useUserStore } from '../store/user'
import { useCommonStore } from '../store/common'
import { useTagsViewStore } from '../store/tagsView'

export default {
  install() {
    return {
      app: useAppStore(),
      common: useCommonStore(),
      settings: useSettingStore(),
      permission: usePermissionStore(),
      user: useUserStore(),
      tagsVies: useTagsViewStore(),
    }
  }
}
