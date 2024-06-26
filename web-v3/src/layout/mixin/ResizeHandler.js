import { useAppStore } from '~/store/app'

const { body } = document
const WIDTH = 992 // refer to Bootstrap's responsive design

export default {
  watch: {
    $route(route) {
      const store = useAppStore()
      if (this.device === 'mobile' && this.sidebar.opened) {
        store.closeSideBar({ withoutAnimation: false })
      }
    }
  },
  beforeMount() {
    window.addEventListener('resize', this.$_resizeHandler)
  },
  beforeDestroy() {
    window.removeEventListener('resize', this.$_resizeHandler)
  },
  mounted() {
    const isMobile = this.$_isMobile()
    const store = useAppStore()
    if (isMobile) {
      store.toggleDevice('mobile')
      store.closeSideBar({ withoutAnimation: true })
    }
  },
  methods: {
    $_isMobile() {
      const rect = body.getBoundingClientRect()
      return rect.width - 1 < WIDTH
    },
    $_resizeHandler() {
      if (!document.hidden) {
        const isMobile = this.$_isMobile()
        const store = useAppStore()
        store.toggleDevice(isMobile ? 'mobile' : 'desktop')
        if (isMobile) {
          store.closeSideBar({ withoutAnimation: true })
        }
      }
    }
  }
}
