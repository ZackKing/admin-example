import { defineStore } from 'pinia'
import Cookies from 'js-cookie'

// Option Store example
export const useAppStore = defineStore('app', {
  state: () => ({
    _sidebar: {
      opened: Cookies.get('sidebarStatus') ? !!+Cookies.get('sidebarStatus') : true,
      withoutAnimation: false
    },
    _device: 'desktop'
  }),
  getters: {
    sidebar: (state) => state._sidebar,
    device: (state) => state._device,
  },
  actions: {
    toggleSideBar() {
      this._sidebar.opened = !this._sidebar.opened
      this._sidebar.withoutAnimation = false
      if (this._sidebar.opened) {
        Cookies.set('sidebarStatus', 1)
      } else {
        Cookies.set('sidebarStatus', 0)
      }
    },
    closeSideBar({ withoutAnimation }) {
      Cookies.set('sidebarStatus', 0)
      this._sidebar.opened = false
      this._sidebar.withoutAnimation = withoutAnimation
    },
    toggleDevice(device) {
      this._device = device
    },
  }
})
