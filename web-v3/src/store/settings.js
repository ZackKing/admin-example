import { defineStore } from 'pinia'
import { ref } from 'vue'
import defaultSettings from '~/settings'

export const useSettingStore = defineStore('setting', () => {
  const showSettings = ref(defaultSettings.showSettings)
  const fixedHeader = ref(defaultSettings.fixedHeader)
  const sidebarLogo = ref(defaultSettings.sidebarLogo)
  const tagsView = ref(defaultSettings.tagsView)

  function changeSetting({ k, val }) {
    switch (k) {
      case 'showSettings':
        showSettings.value = val
        break
      case 'fixedHeader':
        fixedHeader.value = val
        break
      case 'sidebarLogo':
        sidebarLogo.value = val
        break
      case 'tagsView':
        tagsView.value = val
        break
      default:
        break
    }
  }

  return {
    showSettings,
    fixedHeader,
    sidebarLogo,
    tagsView,
    changeSetting,
  }
})

