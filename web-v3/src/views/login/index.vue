<template>
  <div class="login-container">
    <el-form ref="loginFormRef" :model="loginForm" :rules="loginRules" class="login-form" auto-complete="on" label-position="left">

      <div class="title-container">
        <h3 class="title">Admin System</h3>
      </div>

      <el-form-item prop="account">
        <span class="svg-container"><SvgIcon name="user" /></span>
        <el-input ref="account" v-model.trim="loginForm.account" placeholder="account" name="account" type="text" tabindex="1" auto-complete="on" />
      </el-form-item>

      <el-form-item prop="password">
        <span class="svg-container">
          <SvgIcon name="password" />
        </span>
        <el-input :key="passwordType" ref="passwordRef" v-model.trim="loginForm.password" :type="passwordType" placeholder="password" name="password" tabindex="2" auto-complete="on" @keyup.enter="handleLogin" />
        <span class="show-pwd" @click="showPwd"><SvgIcon :name="passwordType === 'password' ? 'eye' : 'eye-open'" /></span>
      </el-form-item>

      <el-button :loading="loading" type="primary" style="width:100%;margin-bottom:30px;" @click.enter.prevent="handleLogin">Login</el-button>

    </el-form>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useUserStore } from '@/store/user'

const validateUsername = (rule, value, callback) => {
  if (value.length === 0) {
    callback(new Error('Please enter the user name'))
  } else {
    callback()
  }
}

const validatePassword = (rule, value, callback) => {
  if (value.length === 0) {
    callback(new Error('Please enter the password'))
  } else {
    callback()
  }
}

const loginForm = ref({
  account: '',
  password: '',
})

const loginRules = {
  account: [{ required: true, min: 5, trigger: 'blur', validator: validateUsername }],
  password: [{ required: true, min: 6, trigger: 'blur', validator: validatePassword }],
}

const loading = ref(false)
const passwordType = ref('password')
const redirect = ref(undefined)

const router = useRouter()
const route = useRoute()
const store = useUserStore()

watch(
  () => route.query,
  (query) => {
    redirect.value = query && query.redirect
  },
  { immediate: true }
)

const showPwd = () => {
  passwordType.value = passwordType.value === 'password' ? '' : 'password'
}

const handleLogin = () => {
  loginFormRef.value.validate((valid) => {
    if (valid) {
      loading.value = true
      store.login(loginForm.value).then(rs => {
        router.push({ path: redirect.value || '/' })
        loading.value = false
      }).catch((err) => {
        console.log(err)
        loading.value = false
        return false
      })
    } else {
      return false
    }
  })
}

const loginFormRef = ref(null)
const passwordRef = ref(null)

</script>

<style lang="scss">
$bg: #1c1a1b;
$light_gray: #fff;
$cursor: #fff;

@supports (-webkit-mask: none) and (not (cater-color: $cursor)) {
  .login-container .el-input input {
    color: $cursor;
  }
}

/* reset element css */
.login-container {
  .el-input {
    height: 47px;
    width: 85%;

    input {
      background: transparent;
      border: 0px;
      border-radius: 0px;
      padding: 12px 5px 12px 15px;
      color: $light_gray;
      height: 47px;
      caret-color: $cursor;

      &:-webkit-autofill {
        box-shadow: 0 0 0px 1000px $bg inset !important;
        -webkit-text-fill-color: $cursor !important;
      }
    }
  }

  .el-form-item {
    border: 1px solid rgba(255, 255, 255, 0.1);
    background: rgba(0, 0, 0, 0.1);
    border-radius: 5px;
    color: #454545;
  }

  .el-input__wrapper {
    background: transparent;
    margin-left: 10px;
    box-shadow: 0 0 0 0px;
  }
}
</style>

<style lang="scss" scoped>
$bg: #2d3a4b;
$dark_gray: #889aa4;
$light_gray: #eee;

.login-container {
  min-height: 100%;
  width: 100%;
  background-color: $bg;
  overflow: hidden;
  background: url('/img/login_bg.jpg');
  background-size: 100%;

  .login-form {
    position: relative;
    width: 520px;
    max-width: 100%;
    padding: 160px 35px 0;
    margin: 0 auto;
    overflow: hidden;
  }

  .tips {
    font-size: 14px;
    color: #fff;
    margin-bottom: 10px;

    span {
      &:first-of-type {
        margin-right: 16px;
      }
    }
  }

  .svg-container {
    padding: 6px 5px 6px 15px;
    color: $dark_gray;
    vertical-align: middle;
    width: 30px;
    display: inline-block;
  }

  .title-container {
    position: relative;

    .title {
      font-size: 26px;
      color: $light_gray;
      margin: 0px auto 40px auto;
      text-align: center;
      font-weight: bold;
    }
  }

  .show-pwd {
    font-size: 16px;
    color: $dark_gray;
    cursor: pointer;
    user-select: none;
  }
}
</style>
