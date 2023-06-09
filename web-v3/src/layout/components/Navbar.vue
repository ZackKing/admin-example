<template>
  <div class="navbar">
    <hamburger :is-active="sidebar.opened" class="hamburger-container" @toggle-click="toggleSideBar" />

    <breadcrumb class="breadcrumb-container" />

    <div class="right-menu">
      <el-dropdown class="avatar-container" trigger="click">
        <div class="avatar-wrapper">
          <i class="el-icon-user-solid" />
          {{ name }}
          <!-- <img :src="avatar+'?imageView2/1/w/80/h/80'" class="user-avatar"> -->
          <i class="el-icon-caret-bottom" />
        </div>
        <template #dropdown>
          <el-dropdown-menu class="user-dropdown">
            <router-link to="/self/info">
              <el-dropdown-item>
                Info
              </el-dropdown-item>
            </router-link>
            <el-dropdown-item @click="logout">
              <span style="display:block;">Log Out</span>
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>

<script setup>
// import { mapState } from 'pinia'
import Breadcrumb from '@/components/Breadcrumb'
import Hamburger from '@/components/Hamburger'
import { useAppStore } from '@/store/app'
import { useUserStore } from '@/store/user'
import { ref } from 'vue'

const store = useAppStore()
const userStore = useUserStore()

const sidebar = ref(store.sidebar)
const name = ref(userStore.name)
// const avatar = ref(userStore.avatar)

function toggleSideBar() {
  store.toggleSideBar()
}

function logout() {
  userStore.resetToken().then(() => {
    location.reload()
  })
}

// export default {
//   components: {
//     Breadcrumb,
//     Hamburger
//   },
//   data() {
//     return {
//       sidebar: store.sidebar
//     }
//   },
//   computed: {
//     ...mapState(useAppStore, [
//       'avatar',
//       'name'
//     ])
//   },
//   methods: {
//     toggleSideBar() {
//       this.$store.app.toggleSideBar()
//     },
//     logout() {
//       this.$store.user.resetToken().then(() => {
//         location.reload()
//       })
//     }
//   }
// }
</script>

<style lang="scss" scoped>
.navbar {
  height: 50px;
  overflow: hidden;
  position: relative;
  background: #fff;
  box-shadow: 0 1px 4px rgba(0,21,41,.08);

  .hamburger-container {
    line-height: 46px;
    height: 100%;
    float: left;
    cursor: pointer;
    transition: background .3s;
    -webkit-tap-highlight-color:transparent;

    &:hover {
      background: rgba(0, 0, 0, .025)
    }
  }

  .breadcrumb-container {
    float: left;
  }

  .right-menu {
    float: right;
    height: 100%;
    line-height: 50px;

    &:focus {
      outline: none;
    }

    .right-menu-item {
      display: inline-block;
      padding: 0 8px;
      height: 100%;
      font-size: 18px;
      color: #5a5e66;
      vertical-align: text-bottom;

      &.hover-effect {
        cursor: pointer;
        transition: background .3s;

        &:hover {
          background: rgba(0, 0, 0, .025)
        }
      }
    }

    .avatar-container {
      margin-right: 30px;
      cursor: pointer;

      .avatar-wrapper {
        // margin-top: 5px;
        position: relative;

        .user-avatar {
          cursor: pointer;
          width: 40px;
          height: 40px;
          border-radius: 10px;
        }

        .el-icon-caret-bottom {
          cursor: pointer;
          position: absolute;
          right: -20px;
          top: 20px;
          font-size: 12px;
        }
      }
    }
  }
}
</style>
