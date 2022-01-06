<template>
  <div class="app-container">
    <el-tabs tab-position="left">
      <el-tab-pane label="Info">
        <el-form ref="infoForm" :model="info" label-width="200px" :rules="infoRules">
          <el-form-item label="User Name: " prop="name">
            <el-input v-model="info.name" placeholder="login user name" :disabled="true" style="width: 200px" />
          </el-form-item>
          <el-form-item label="Real Name: " prop="real_name">
            <el-input v-model="info.real_name" placeholder="real name" style="width: 200px" />
          </el-form-item>
          <el-form-item label="Mobile: " prop="mobile">
            <el-input v-model="info.mobile" placeholder="mobile phone number" style="width: 200px" />
          </el-form-item>
          <el-form-item label="Email: " prop="email">
            <el-input v-model="info.email" placeholder="email address" style="width: 200px" />
          </el-form-item>
          <el-form-item label="Desc: " prop="desc">
            <el-input
              v-model="info.desc"
              type="textarea"
              placeholder="user desc"
              :autosize="{ minRows: 2, maxRows: 4}"
              style="width: 300px"
            />
          </el-form-item>
          <el-form-item>
            <el-button
              :disabled="infoLock"
              type="primary"
              @click="handleEditInfo"
            >Submit</el-button>
          </el-form-item>
        </el-form>
      </el-tab-pane>
      <el-tab-pane label="Password">
        <el-form ref="pwdForm" :model="pwd" label-width="200px" :rules="pwdRules">
          <el-form-item label="Current Password: " prop="old_password">
            <el-input v-model="pwd.old_password" placeholder="current password" style="width: 200px" show-password />
          </el-form-item>
          <el-form-item label="Password: " prop="password">
            <el-input v-model="pwd.password" placeholder="new password" style="width: 200px" show-password autocomplete="off" />
          </el-form-item>
          <el-form-item label="Retry Password: " prop="retry_password">
            <el-input v-model="pwd.retry_password" placeholder="retry password" style="width: 200px" show-password autocomplete="off" />
          </el-form-item>
          <el-form-item>
            <el-button :disabled="pwdLock" type="primary" @click="handleEditPassword">Submit</el-button>
          </el-form-item>
        </el-form>
      </el-tab-pane>
    </el-tabs>

  </div>
</template>

<script>
import { getInfo, editInfo, editPassword } from '@/api/user'

export default {
  name: 'SelfInfo',
  components: {},
  data() {
    return {
      infoLock: true,
      pwdLock: false,
      info: {},
      pwd: {
        old_password: null,
        password: null,
        retry_password: null
      },
      infoRules: {
        real_name: [{ required: true, min: 2, max: 64, trigger: 'change' }],
        mobile: [{ required: true, min: 11, max: 20, trigger: 'change' }],
        email: [{ required: true, type: 'email', min: 6, max: 128, trigger: 'change' }]
      },
      pwdRules: {
        old_password: [{ required: true, min: 6, max: 64, trigger: 'change' }],
        password: [{ required: true, min: 6, max: 64, trigger: 'change' }],
        retry_password: [{ required: true, min: 6, max: 64, trigger: 'change' }]
      }
    }
  },
  created() {
    this.refreshInfo()
  },
  methods: {
    refreshInfo() {
      getInfo().then(response => {
        this.info = response.data
        this.infoLock = false
      })
    },
    handleEditInfo() {
      this.infoLock = true
      this.$refs['infoForm'].validate((valid) => {
        if (valid) {
          editInfo({
            real_name: this.info.real_name,
            mobile: this.info.mobile,
            email: this.info.email,
            desc: this.info.desc
          }).then(response => {
            this.$notify({ title: 'Success', message: 'Update Successfully', type: 'success', duration: 2000 })
            this.refreshInfo()
          }).finally(() => {
            this.infoLock = false
          })
        } else {
          this.infoLock = false
          this.$notify({ title: 'Error', message: 'Info Data Valid Error!', type: 'error', duration: 5000 })
        }
      })
    },
    handleEditPassword() {
      this.pwdLock = true
      this.$refs['pwdForm'].validate((valid) => {
        if (valid) {
          if (this.pwd.password !== this.pwd.retry_password) {
            this.$notify({ title: 'Error', message: 'Retry password error', type: 'error', duration: 5000 })
            this.pwdLock = false
            return
          }
          if (this.pwd.password === this.pwd.old_password) {
            this.$notify({ title: 'Error', message: 'Password cannot be the same as the current password', type: 'error', duration: 5000 })
            this.pwdLock = false
            return
          }
          editPassword({ old_password: this.pwd.old_password, password: this.pwd.password }).then(response => {
            this.$notify({ title: 'Success', message: 'Update Password Successfully', type: 'success' })
            this.logout()
          }).finally(() => {
            this.pwdLock = false
          })
        } else {
          this.pwdLock = false
          this.$notify({ title: 'Error', message: 'Password Data Valid Error!', type: 'error', duration: 5000 })
        }
      })
    },
    resetPwd() {
      this.pwd = {
        old_password: null,
        password: null,
        retry_password: null
      }
    },
    logout() {
      this.$store.dispatch('user/resetToken').then(() => {
        location.reload()
      })
    }
  }
}
</script>
