<template>
  <div class="app-container">
    <div class="filter-container">
      <el-form :inline="true" :model="filter">
        <el-form-item>
          <el-button class="filter-item" :disabled="listLoading" type="primary" :icon="PlusIcon" @click="handleCreate">Add</el-button>
          <el-button class="filter-item" :disabled="listLoading" type="primary" @click="handleFilter">Refresh</el-button>
        </el-form-item>
      </el-form>
    </div>

    <el-table v-loading="listLoading" :data="list" border style="width: 100%;">
      <el-table-column prop="id" label="ID" />
      <el-table-column prop="name" label="Name" />
      <el-table-column prop="remark" label="Remark" />
      <el-table-column label="Status">
        <template #default="{row}">
          <el-switch v-model="row.status" :disabled="row.id === 1" @click="switchStatus(row)" />
        </template>
      </el-table-column>
      <el-table-column prop="created_time" label="Created Time" />
      <el-table-column prop="updated_time" label="Updated Time" />
      <el-table-column label="Actions" width="300">
        <template #default="{row}">
          <el-button :disabled="row.id === 1" type="primary" @click="handleUpdate(row)">Edit</el-button>
          <el-button type="primary" @click="handleSetUser(row)">Set User</el-button>
          <el-button :disabled="row.id === 1" type="primary" @click="handleSetMenu(row)">Set Menu</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog v-model="dialogOne.dialogFormVisible" :title="dialogOne.dialogStatus === 'create' ? 'Create' : 'Edit'">
      <el-form ref="dialogOne" :rules="dialogOne.rules" :model="dialogOne.temp" label-position="left" label-width="100px">
        <el-form-item label="Name" prop="name">
          <el-input v-model="dialogOne.temp.name" />
        </el-form-item>
        <el-form-item label="Remark">
          <el-input
            v-model="dialogOne.temp.remark"
            :autosize="{ minRows: 2, maxRows: 4}"
            type="textarea"
            placeholder="Please input"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div>
          <el-button @click="dialogOne.dialogFormVisible = false">Cancel</el-button>
          <el-button
            type="primary"
            @click="dialogOne.dialogStatus==='create' ? createData() : updateData()"
          >Confirm</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="dialogTwo.dialogFormVisible" title="Edit">
      <el-form ref="dialogTwo" label-position="left" label-width="100px">
        <el-form-item v-if="dialogTwo.dialogStatus==='setMenu' && dialogTwo.dialogFormVisible" label="menuTree">
          <el-tree
            ref="menutree"
            :data="dialogTwo.menuTree"
            highlight-current
            show-checkbox
            default-expand-all
            check-on-click-node
            :default-checked-keys="filterMids"
            node-key="id"
            :props="dialogTwo.menuProps"
          />
        </el-form-item>
        <el-form-item v-show="dialogTwo.dialogStatus==='setUser'" label="userList">
          <el-checkbox-group v-model="dialogTwo.temp.uids">
            <el-checkbox v-for="item in dialogTwo.userList" :key="item.uid" :label="item.uid">{{ item.name }}</el-checkbox>
          </el-checkbox-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <div>
          <el-button @click="dialogTwo.dialogFormVisible = false">Cancel</el-button>
          <el-button type="primary" @click="dialogTwo.dialogStatus==='setUser' ? toSetUser() : toSetMenu()">Confirm</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { Plus as PlusIcon } from '@element-plus/icons-vue'
</script>

<script>
import {
  getGroupList,
  addGroup,
  updateGroup,
  setUser,
  setMenu
} from '@/api/group'
import { getMenuList } from '@/api/menu'
import { getAccountList } from '@/api/account'
import { ElMessage, ElMessageBox } from 'element-plus'

export default {
  name: 'GroupManage',
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
          id: undefined,
          name: '',
          remark: ''
        },
        dialogFormVisible: false,
        dialogStatus: '',
        rules: {
          name: [
            { required: true, message: 'name is required', trigger: 'change' }
          ]
        }
      },
      dialogTwo: {
        temp: {
          id: undefined,
          uids: [],
          mids: []
        },
        dialogFormVisible: false,
        dialogStatus: '',
        menuTree: [],
        menuTreeParent: [],
        menuProps: {
          children: 'sub_menu',
          label: 'name'
        },
        userList: []
      }
    }
  },
  computed: {
    filterMids: function () {
      return this.dialogTwo.temp.mids.filter(item => {
        return !this.dialogTwo.menuTreeParent.includes(item)
      })
    }
  },
  created() {
    this.getList()
    this.getMenuList()
    this.getAccountList()
  },
  methods: {
    getAccountList() {
      getAccountList({ size: 999 }).then(response => {
        this.dialogTwo.userList = response.data.list
      })
    },
    getMenuList() {
      getMenuList().then(response => {
        this.dialogTwo.menuTree = response.data
        const getParentKey = (arr) => {
          arr.forEach(item => {
            if (item.sub_menu) {
              this.dialogTwo.menuTreeParent.push(item.id)
              getParentKey(item.sub_menu)
            }
          })
        }
        getParentKey(this.dialogTwo.menuTree)
      })
    },
    getList() {
      this.listLoading = true
      getGroupList(this.filter).then(response => {
        this.list = response.data.map(item => {
          item.status = !!item.status
          return item
        })
        this.listLoading = false
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
      this.$refs['dialogOne'].validate(valid => {
        if (valid) {
          addGroup(this.dialogOne.temp).then(() => {
            this.handleFilter()
            this.dialogOne.dialogFormVisible = false
            ElMessage({
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
      this.$refs['dialogOne'].validate(valid => {
        if (valid) {
          const tempData = Object.assign({}, this.dialogOne.temp)
          delete tempData.password
          updateGroup(tempData).then(() => {
            const index = this.list.findIndex(
              v => v.id === this.dialogOne.temp.id
            )
            this.list.splice(index, 1, this.dialogOne.temp)
            this.dialogOne.dialogFormVisible = false
            ElMessage({
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
      ElMessageBox('Please confirm the operation?', 'Tips', {
        confirmButtonText: 'Confirm',
        cancelButtonText: 'Cancel',
        type: 'warning'
      }).then(() => {
        updateGroup({ status: row.status == 0 ? 0 : 1, id: row.id }).then(() => {
          this.dialogOne.dialogFormVisible = false
          ElMessage({ title: 'Success', message: 'Update Successfully', type: 'success', duration: 2000 })
          this.getList()
        })
      })
    },
    handleSetUser(row) {
      this.dialogTwo.dialogStatus = 'setUser'
      this.dialogTwo.temp = Object.assign({}, row)
      this.dialogTwo.dialogFormVisible = true
    },
    handleSetMenu(row) {
      this.dialogTwo.dialogStatus = 'setMenu'
      this.dialogTwo.temp = Object.assign({}, row)
      this.dialogTwo.dialogFormVisible = true
    },
    toSetUser() {
      setUser({
        id: this.dialogTwo.temp.id,
        uids: this.dialogTwo.temp.uids
      }).then(() => {
        this.dialogTwo.dialogFormVisible = false
        const index = this.list.findIndex(v => v.id === this.dialogTwo.temp.id)
        this.list[index].uids = this.dialogTwo.temp.uids
        ElMessage({
          title: 'Success',
          message: 'Update Successfully',
          type: 'success',
          duration: 2000
        })
      })
    },
    toSetMenu() {
      const current = this.$refs.menutree.getCheckedKeys().concat(this.$refs.menutree.getHalfCheckedKeys())
      setMenu({
        id: this.dialogTwo.temp.id,
        menu_ids: current
      }).then(() => {
        this.dialogTwo.dialogFormVisible = false
        const index = this.list.findIndex(v => v.id === this.dialogTwo.temp.id)
        this.list[index].mids = current
        ElMessage({
          title: 'Success',
          message: 'Update Successfully',
          type: 'success',
          duration: 2000
        })
      })
    }
  }
}
</script>
