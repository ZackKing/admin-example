<template>
  <div class="app-container">
    <div class="container-header">
      <el-form :inline="true" :model="filter" class="header-filter">
        <el-form-item label="Account:">
          <el-input v-model.trim="filter.name" clearable placeholder="Account" />
        </el-form-item>
        <el-form-item >
          <el-button class="filter-item" :disabled="loading" type="primary" :icon="SearchIcon" @click="handleFilter">Search</el-button>
        </el-form-item>
        <el-form-item>
          <el-button class="filter-item" :disabled="loading" type="primary" :icon="PlusIcon" @click="handleCreate">Add</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="container-content">
      <el-table v-loading="loading" :data="list" border>
        <el-table-column prop="uid" label="ID" />
        <el-table-column prop="name" label="Account" />
        <el-table-column prop="real_name" label="Name" />
        <el-table-column prop="mobile" label="Phone" />
        <el-table-column prop="email" label="Email" />
        <el-table-column prop="group" label="Group">
          <template #default="{row}">
            <el-tag class="mr-1" v-for="g in row.group" :item="g">{{ g.name }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="Status">
          <template #default="{row}">
            <el-switch v-model="row.status" @click="switchStatus(row)" />
          </template>
        </el-table-column>
        <el-table-column prop="desc" label="Remark" />
        <el-table-column label="Actions" width="300px">
          <template #default="{row}">
            <el-button type="primary" @click="handleUpdate(row)">Edit</el-button>
            <el-button type="primary" @click="handleSetAuth(row)">Auth</el-button>
            <!-- <el-button type="danger" @click="deleteAccount(row)">Delete</el-button> -->
          </template>
        </el-table-column>
      </el-table>

      <el-pagination
        v-model:current-page="filter.page"
        v-model:page-size="filter.size"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total"
        @size-change="refresh"
        @current-change="refresh"
      />
    </div>

    <el-dialog v-model="dialogOne.visible" :title="dialogOne.dialogStatus === 'create' ? 'Add' : 'Edit'">
      <el-form ref="dialogOneRef" v-loading="dialogOne.loading" :rules="dialogOne.rules" :model="dialogOne.temp" label-position="left" label-width="100px" style="width: 400px; margin-left:50px;">
        <el-form-item label="Account" prop="name">
          <el-input v-model="dialogOne.temp.name" />
        </el-form-item>
        <el-form-item v-if="dialogOne.dialogStatus==='create'" label="Password" prop="password">
          <el-input v-model="dialogOne.temp.password" />
        </el-form-item>
        <el-form-item v-else label="Password">
          <el-button type="primary" @click="resetPwd(dialogOne.temp)">Reset</el-button>
        </el-form-item>
        <el-form-item label="Name" prop="real_name">
          <el-input v-model="dialogOne.temp.real_name" />
        </el-form-item>
        <el-form-item label="Phone" prop="mobile">
          <el-input v-model="dialogOne.temp.mobile" />
        </el-form-item>
        <el-form-item label="Email" prop="email">
          <el-input v-model="dialogOne.temp.email" />
        </el-form-item>
        <el-form-item label="Remark">
          <el-input v-model="dialogOne.temp.desc" :autosize="{ minRows: 2, maxRows: 4}" type="textarea" placeholder="Please input" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogOne.visible = false">Cancel</el-button>
        <el-button type="primary" :disabled="dialogOne.loading" @click="dialogOne.dialogStatus==='create' ? createData() : updateData()">Confirm</el-button>
      </template>
    </el-dialog>

    <el-dialog v-model="dialogTwo.visible" title="Auth" center>
      <el-form v-loading="dialogTwo.loading" ref="dialogTwoRef" label-position="left" class="px-8">
        <el-form-item label="Group">
          <el-select v-model="dialogTwo.temp.group" multiple filterable placeholder="select group">
            <el-option v-for="item in options.group" :key="item.id" :label="item.name" :value="item.id" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogTwo.visible = false">Cancel</el-button>
        <el-button type="primary" @click="toSetAuth()" :disabled="dialogTwo.loading">Confirm</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { Search as SearchIcon, Plus as PlusIcon } from '@element-plus/icons-vue'
import { getAccountList, addAccount, updateAccount, setStatus, setGroup } from '~/api/account'
import { getGroupList } from '~/api/group'
import { random } from 'lodash-es'
import { ElMessage, ElMessageBox } from 'element-plus'

const list = ref(null)
const total = ref(0)
const loading = ref(true)
const filter = reactive({ page: 1, size: 20, name: '' })

const dialogOne = reactive({
  loading: false,
  temp: {
    uid: undefined,
    name: '',
    real_name: '',
    mobile: '',
    email: '',
    password: '',
    desc: ''
  },
  visible: false,
  dialogStatus: '',
  rules: {
    name: [{ required: true, min: 6, max: 32, message: 'Account has 6 - 32 bit', trigger: 'change' }],
    real_name: [{ required: true, message: 'Name required', trigger: 'change' }],
    password: [{ required: true, min: 6, message: 'Password must more than 6 bit', trigger: 'change' }],
    mobile: [{ required: true, message: 'Phone required', trigger: 'change' }],
  }
})

const dialogTwo = reactive({
  loading: false,
  temp: {
    uid: undefined,
    group: [],
  },
  visible: false,
  dialogStatus: '',
  groupList: []
})

const options = reactive({
  group: [],
})

const dialogOneRef = ref(null)
const dialogTwoRef = ref(null)

onMounted(() => {
  refresh()
  initOptions()
})

function refresh() {
  loading.value = true
  getAccountList(filter).then(rs => {
    list.value = rs.data.list.map((item) => {
      item.status = !!item.status
      item.groupIds = item.group ? item.group.map(it => it.id) : []
      return item
    })
    console.log(list)
    total.value = rs.data.total
    loading.value = false
  })
}

function initOptions () {
  const allFilter = { page: 1, size: 9999 }
  getGroupList(allFilter).then(rs => {
    options.group = rs.data
  })
}

function handleFilter() {
  filter.page = 1
  refresh()
}

function resetTemp() {
  dialogOne.temp = {
    uid: undefined,
    name: '',
    real_name: '',
    mobile: '',
    email: '',
    password: '',
    desc: ''
  }
}

function handleCreate() {
  resetTemp()
  dialogOne.dialogStatus = 'create'
  dialogOne.visible = true
}

function createData() {
  dialogOneRef.value.validate((valid) => {
    if (valid) {
      dialogOne.loading = true
      addAccount(dialogOne.temp).then(() => {
        handleFilter()
        dialogOne.visible = false
        ElMessage({ title: 'Success', message: 'OK', type: 'success', duration: 2000 })
      }).finally(() => {
        dialogOne.loading = false
      })
    }
  })
}

function handleUpdate(row) {
  dialogOne.temp = Object.assign({}, row)
  dialogOne.visible = true
  dialogOne.dialogStatus = 'edit'
}

function updateData() {
  dialogOneRef.value.validate((valid) => {
    if (valid) {
      dialogOne.loading = true
      const tempData = Object.assign({}, dialogOne.temp)
      delete tempData.password
      updateAccount(tempData).then(() => {
        const index = list.value.findIndex(v => v.id === dialogOne.temp.id)
        list.value.splice(index, 1, dialogOne.temp)
        dialogOne.visible = false
        ElMessage({ title: 'Success', message: 'OK', type: 'success', duration: 2000})
      }).finally(() => {
        dialogOne.loading = false
      })
    }
  })
}

function switchStatus(row) {
  ElMessageBox.confirm('Please confirm the operation?', 'Warning', {
    confirmButtonText: 'Confirm',
    cancelButtonText: 'Cancel',
    type: 'warning'
  }).then(() => {
    setStatus({ status: !row.status ? 0 : 1, uid: row.uid }).then(() => {
      dialogOne.visible = false
      ElMessage({
        title: 'Success',
        message: 'OK',
        type: 'success',
        duration: 2000
      })
    })
  }).catch(err => {
    console.error(err)
    row.status = !row.status
  })
}

function deleteAccount(row) {
  ElMessageBox.confirm('Delete User: ' + row.name, 'Warning', { confirmButtonText: 'OK', cancelButtonText: 'Cancel', type: 'warning' }).then(() => {
    setStatus({ status: 2, uid: row.uid }).then(() => {
      ElMessage({ title: 'Success', message: 'OK', type: 'success' })
      refresh()
    })
  })
}

function resetPwd(row) {
  ElMessageBox.confirm('Confirm?', 'Warning', { confirmButtonText: 'OK', cancelButtonText: 'Cancel', type: 'warning' }).then(() => {
    const password = `Admin${random(100000, 999999)}`
    updateAccount({ uid: row.uid, password }).then(rs => {
      ElMessageBox.confirm(`New Password: ${password}`, 'New', { confirmButtonText: 'OK', cancelButtonText: 'Cancel', type: 'warning' })
    })
  })
}

function handleSetAuth(row) {
  dialogTwo.dialogStatus = 'setGroup'
  dialogTwo.temp = {
    uid: row.uid,
    group: row.groupIds,
  }
  dialogTwo.visible = true
}

async function toSetAuth() {
  dialogTwo.loading = true
  await setGroup({
    id: dialogTwo.temp.uid,
    group_ids: dialogTwo.temp.group
  })
  dialogTwo.loading = false
  dialogTwo.visible = false
  ElMessage({ title: 'Success', message: 'OK', type: 'success', duration: 2000 })
  refresh()
}

</script>
