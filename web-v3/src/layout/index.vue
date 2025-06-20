<template>
  <div :class="classObj" class="app-wrapper">
    <div v-if="device==='mobile'&&sidebar.opened" class="drawer-bg" @click="handleClickOutside" />
    <Sidebar class="sidebar-container" />
    <div :class="{hasTagsView:needTagsView}" class="main-container" v-watermark="watermark">
      <div :class="{'fixed-header':fixedHeader}">
        <Navbar />
        <TagsView v-if="needTagsView" />
      </div>
      <AppMain />
    </div>
  </div>
</template>

<script setup>
import { Navbar, Sidebar, AppMain, TagsView } from './components'
import { ref, computed, getCurrentInstance } from 'vue'

const { appContext } = getCurrentInstance()
const store = appContext.config.globalProperties.$store

const watermark = ref('')
watermark.value = `${store.user.name} - ${store.user.nickname}`

const sidebar = computed(() => store.app.sidebar)
const device = computed(() => store.app.device)
const fixedHeader = computed(() => store.settings.fixedHeader)
const needTagsView = computed(() => store.settings.tagsView)
const classObj = computed(() => ({
  hideSidebar: !sidebar.value.opened,
  openSidebar: sidebar.value.opened,
  withoutAnimation: sidebar.value.withoutAnimation,
  mobile: device.value === 'mobile'
}))

function handleClickOutside() {
  store.app.closeSideBar({ withoutAnimation: false })
}
</script>

<script>
import ResizeMixin from './mixin/ResizeHandler'
export default {
  mixins: [ResizeMixin],
}
</script>

<style lang="scss" scoped>
  @use "~/styles/mixin.scss" as *;

  .app-wrapper {
    @include clearfix;
    position: relative;
    height: 100%;
    width: 100%;
    &.mobile.openSidebar{
      position: fixed;
      top: 0;
    }
  }
  .drawer-bg {
    background: #000;
    opacity: 0.3;
    width: 100%;
    top: 0;
    height: 100%;
    position: absolute;
    z-index: 999;
  }

  .fixed-header {
    position: fixed;
    top: 0;
    right: 0;
    z-index: 9;
    width: calc(100% - #{210px});
    transition: width 0.28s;
  }

  .hideSidebar .fixed-header {
    width: calc(100% - 54px)
  }

  .mobile .fixed-header {
    width: 100%;
  }
</style>
