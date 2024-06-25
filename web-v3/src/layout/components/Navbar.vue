<template>
  <div class="h-12 overflow-hidden relative shadow-sm">
    <hamburger :is-active="sidebar.opened" class="hamburger-container h-full float-left leading-[48px] hover:bg-black/5 cursor-pointer transition-colors duration-300" @toggle-click="toggleSideBar" />

    <breadcrumb class="breadcrumb-container" />

    <div class="h-full float-right focus:outline-none">
      <!-- <span class="h-full text-center leading-[48px]">{{ nickname }}</span> -->
      <el-dropdown>
        <div class="h-full w-30 p-2">
          <span class="h-8 inline-block mr-1">{{ nickname }}</span>
          <img :src="avatar" class="h-8 w-8 inline-block ring-2 ring-white">
        </div>
        <template #dropdown>
          <el-dropdown-menu>
            <router-link to="/self/info">
              <el-dropdown-item>Info</el-dropdown-item>
            </router-link>
            <el-dropdown-item @click="logout">Logout</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>

<script setup>
import Breadcrumb from '~/components/Breadcrumb'
import Hamburger from '~/components/Hamburger'
import { useAppStore } from '~/store/app'
import { useUserStore } from '~/store/user'
import { ref } from 'vue'

const store = useAppStore()
const userStore = useUserStore()

const sidebar = ref(store.sidebar)
const nickname = ref(userStore.nickname)
const avatar = ref(userStore.avatar)
if (!avatar.value) {
  avatar.value = '/favicon.ico'
}

function toggleSideBar() {
  store.toggleSideBar()
}

function logout() {
  userStore.resetToken().then(() => {
    location.reload()
  })
}
</script>

<style lang="scss" scoped>
.hamburger-container {
  -webkit-tap-highlight-color:transparent;
}
</style>
