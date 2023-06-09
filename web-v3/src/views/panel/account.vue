<template>
  <div class="app-container">
    <div class="filter-container">
      <el-form :inline="true" :model="filter">
        <el-form-item label="Account:">
          <el-input
            v-model.trim="filter.name"
            clearable
            placeholder="Account"
            @keyup.enter="handleFilter"
          />
        </el-form-item>
        <el-form-item>
          <el-button
            class="filter-item"
            :disabled="listLoading"
            type="primary"
            icon="el-icon-search"
            @click="handleFilter"
          >Search</el-button>
        </el-form-item>
        <el-form-item>
          <el-button
            class="filter-item"
            :disabled="listLoading"
            type="primary"
            icon="el-icon-plus"
            @click="handleCreate"
          >Add</el-button>
        </el-form-item>
      </el-form>
    </div>

    <el-table v-loading="listLoading" :data="list" border style="width: 100%;">
      <el-table-column prop="uid" label="User ID" />
      <el-table-column prop="name" label="Account" />
      <el-table-column prop="real_name" label="Real Name" />
      <el-table-column prop="mobile" label="Mobile" />
      <el-table-column prop="email" label="Email" />
      <el-table-column prop="group" label="Groups" :formatter="groupFormatter" />
      <el-table-column label="Status">
        <template #default="{row}">
          <el-switch v-model="row.status" @click.enter="switchStatus(row)" />
        </template>
      </el-table-column>
      <el-table-column prop="desc" label="Remark" />
      <el-table-column label="Actions" width="300px">
        <template #default="{row}">
          <el-button type="primary" size="small" @click="handleUpdate(row)">Edit</el-button>
          <el-button type="primary" size="small" @click="handleSetGroup(row)">Set Group</el-button>
          <el-button type="primary" size="small" @click="deleteAccount(row)">Delete</el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-model:page="filter.page" v-model:limit="filter.size" :total="total" @pagination="getList" />

    <el-dialog v-model:visible="dialogOne.dialogFormVisible" :title="dialogOne.dialogStatus === 'create' ? 'Create' : 'Edit'">
      <el-form ref="dialogOne" :rules="dialogOne.rules" :model="dialogOne.temp" label-position="left" label-width="100px" style="width: 400px; margin-left:50px;">
        <el-form-item label="Account" prop="name">
          <el-input v-model="dialogOne.temp.name" />
        </el-form-item>
        <!-- <el-form-item label="Real Name" prop="real_name">
          <el-input v-model="dialogOne.temp.real_name" />
        </el-form-item> -->
        <el-form-item v-if="dialogOne.dialogStatus==='create'" label="Password" prop="password">
          <el-input v-model="dialogOne.temp.password" />
        </el-form-item>
        <el-form-item v-else label="Password">
          <el-button type="primary" @click="resetPwd(dialogOne.temp)">Reset Password</el-button>
        </el-form-item>
        <!-- <el-form-item label="Mobile" prop="mobile">
          <el-input v-model="dialogOne.temp.mobile" />
        </el-form-item> -->
        <el-form-item label="Department" prop="department">
          <el-select v-model="dialogOne.temp.department" placeholder="select department">
            <el-option v-for="item in options.department" :key="item.value" :label="item.value" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="Email" prop="email">
          <el-input v-model="dialogOne.temp.email" />
        </el-form-item>
        <el-form-item label="Remark">
          <el-input v-model="dialogOne.temp.desc" :autosize="{ minRows: 2, maxRows: 4}" type="textarea" placeholder="Please input" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div>
          <el-button @click="dialogOne.dialogFormVisible = false">
            Cancel
          </el-button>
          <el-button type="primary" @click="dialogOne.dialogStatus==='create' ? createData() : updateData()">
            Confirm
          </el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model:visible="dialogTwo.dialogFormVisible" title="Edit">
      <el-form ref="dialogTwo" label-position="left" label-width="150px">
        <el-form-item v-show="dialogTwo.dialogStatus==='setGroup'" label="Group List">
          <el-checkbox-group v-model="dialogTwo.temp.group">
            <el-checkbox v-for="item in dialogTwo.groupList" :key="item.id" style="min-width: 160px" :label="item.id">{{ item.name }}</el-checkbox>
          </el-checkbox-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <div>
          <el-button @click="dialogTwo.dialogFormVisible = false">Cancel</el-button>
          <el-button type="primary" @click="toSetGroup()">Confirm</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import { getAccountList, addAccount, updateAccount, setStatus, setGroup } from '@/api/account'
import { getGroupList } from '@/api/group'
import Pagination from '@/components/Pagination'
import _ from 'lodash'

export default {
  name: 'AccountManage',
  components: {
    Pagination
  },
  data() {
    return {
      list: null,
      total: 0,
      listLoading: true,
      filter: {
        page: 1,
        size: 20,
        name: ''
      },
      dialogOne: {
        temp: {
          uid: undefined,
          name: '',
          department: '',
          // real_name: '',
          // mobile: '',
          email: '',
          password: '',
          desc: ''
        },
        dialogFormVisible: false,
        dialogStatus: '',
        rules: {
          name: [{ required: true, min: 6, max: 32, message: 'account must be 6 to 32 digits', trigger: 'change' }],
          // real_name: [{ required: true, message: 'real name is required', trigger: 'change' }],
          password: [{ required: true, min: 6, message: 'password must be greater than or equal to 6 digits', trigger: 'change' }],
          // mobile: [{ required: true, message: 'mobile is required', trigger: 'change' }],
          email: [{ required: true, message: 'email is required', trigger: 'change' }],
          department: [{ required: true, message: 'Department is required', trigger: 'change' }],
        }
      },
      dialogTwo: {
        temp: {
          uid: undefined,
          group: []
        },
        dialogFormVisible: false,
        dialogStatus: '',
        groupList: []
      },
      options: {
        department: [
          { value: 'R&D / PM' },
          { value: 'R&D / Client' },
          { value: 'R&D / Server' },
          { value: 'R&D / Design' },
          { value: 'R&D / QA' },
          { value: 'Marketing / CS' },
          { value: 'Marketing / Sales' },
          { value: 'Admin' },
          { value: 'Others' },
        ]
      },
    }
  },
  created() {
    this.getList()
    this.getGroupList()
  },
  methods: {
    getList() {
      this.listLoading = true
      getAccountList(this.filter).then(response => {
        this.list = response.data.list.map((item) => {
          item.status = !!item.status
          item.groupList = item.group || []
          item.group = item.group ? item.group.map(it => {
            return it.id
          }) : []
          return item
        })
        this.total = response.data.total
        this.listLoading = false
      })
    },
    getGroupList() {
      getGroupList(this.filter).then(response => {
        this.dialogTwo.groupList = response.data
      })
    },
    handleFilter() {
      this.filter.page = 1
      this.getList()
    },
    resetTemp() {
      this.dialogOne.temp = {
        uid: undefined,
        name: '',
        real_name: '',
        mobile: '',
        email: '',
        password: '',
        desc: ''
      }
    },
    handleCreate() {
      this.resetTemp()
      this.dialogOne.dialogStatus = 'create'
      this.dialogOne.dialogFormVisible = true
      // this.$nextTick(() => {
      //   this.$refs['dialogOne'].clearValidate()
      // })
    },
    createData() {
      this.$refs['dialogOne'].validate((valid) => {
        if (valid) {
          addAccount(this.dialogOne.temp).then(() => {
            this.handleFilter()
            this.dialogOne.dialogFormVisible = false
            this.$notify({
              title: 'Success',
              message: 'Created Successfully',
              type: 'success',
              duration: 2000
            })
          })
        }
      })
    },
    handleUpdate(row) {
      this.dialogOne.temp = Object.assign({}, row)
      this.dialogOne.dialogFormVisible = true
      this.dialogOne.dialogStatus = 'edit'
      // this.$nextTick(() => {
      //   this.$refs['dialogOne'].clearValidate()
      // })
    },
    updateData() {
      this.$refs['dialogOne'].validate((valid) => {
        if (valid) {
          const tempData = Object.assign({}, this.dialogOne.temp)
          delete tempData.password
          updateAccount(tempData).then(() => {
            const index = this.list.findIndex(v => v.id === this.dialogOne.temp.id)
            this.list.splice(index, 1, this.dialogOne.temp)
            this.dialogOne.dialogFormVisible = false
            this.$notify({
              title: 'Success',
              message: 'Update Successfully',
              type: 'success',
              duration: 2000
            })
          })
        }
      })
    },
    switchStatus(row) {
      this.$confirm('Please confirm the operation?', 'Tips', {
        confirmButtonText: 'Confirm',
        cancelButtonText: 'Cancel',
        type: 'warning'
      }).then(() => {
        setStatus({ status: row.status ? 0 : 1, uid: row.uid }).then(() => {
          this.dialogOne.dialogFormVisible = false
          this.$notify({
            title: 'Success',
            message: 'Update Successfully',
            type: 'success',
            duration: 2000
          })
          row.status = !row.status
        })
      })
    },
    deleteAccount(row) {
      this.$confirm('Delete this account. Are you sure?', 'Tips', {
        confirmButtonText: 'Confirm',
        cancelButtonText: 'Cancel',
        type: 'warning'
      }).then(() => {
        setStatus({ status: 2, uid: row.uid }).then(() => {
          this.$notify({ title: 'Success', message: 'Update Successfully', type: 'success' })
          this.getList()
        })
      })
    },
    resetPwd(row) {
      this.$confirm('Reset this user\'s password. Are you sure?', 'Tips', {
        confirmButtonText: 'Confirm',
        cancelButtonText: 'Cancel',
        type: 'warning'
      }).then(() => {
        const password = `Admin${_.random(10000, 99999)}`
        updateAccount({ uid: row.uid, password }).then(rs => {
          this.$confirm(`New password: ${password}`, 'Tips', {
            confirmButtonText: 'Confirm',
            cancelButtonText: 'Cancel',
            type: 'warning'
          })
        })
      })
    },
    handleSetGroup(row) {
      this.dialogTwo.dialogStatus = 'setGroup'
      this.dialogTwo.temp = {
        uid: row.uid,
        group: row.group
      }
      this.dialogTwo.dialogFormVisible = true
    },
    toSetGroup() {
      setGroup({
        id: this.dialogTwo.temp.uid,
        group_ids: this.dialogTwo.temp.group
      }).then(() => {
        this.dialogTwo.dialogFormVisible = false
        const index = this.list.findIndex(v => v.uid === this.dialogTwo.temp.uid)
        this.list.splice(index, 1, { ...this.list[index], group: this.dialogTwo.temp.group })
        this.$notify({
          title: 'Success',
          message: 'Update Successfully',
          type: 'success',
          duration: 2000
        })
      })
    },
    groupFormatter(row) {
      let str = ''
      row.groupList.forEach(e => {
        str += e.name + ', '
      })
      return str.replace(/, $/gi, '')
    }
  }
}
</script>
