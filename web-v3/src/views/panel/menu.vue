<template>
  <div class="app-container">

    <el-table
      v-loading="listLoading"
      :data="list"
      border
      fit
      highlight-current-row
      row-key="name"
      :tree-props="{children: 'sub_menu', hasChildren: 'hasChildren'}"
      style="width: 100%;"
    >
      <el-table-column prop="name" label="Name" width="200" />
      <el-table-column prop="uri" label="Url" width="110" />
      <el-table-column prop="icon" label="Icon" width="110" />
      <el-table-column label="Status" width="110">
        <template #default="{row}">
          <el-tag v-if="row.status === 1">valid</el-tag>
          <el-tag v-else type="info">invalid</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="remark" label="Remark" />
      <el-table-column prop="created_time" label="created_time" width="150" />
      <el-table-column prop="updated_time" label="updated_time" width="150" />
      <el-table-column label="Actions" width="150">
        <template #default="{row}">
          <el-button type="primary" @click="handleSetGroup(row)">Set Group</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog v-model="dialogTwo.dialogFormVisible" title="Edit">
      <el-form ref="dialogTwo" label-position="left" label-width="100px" style="width: 640px; margin-left:50px;">
        <el-form-item v-show="dialogTwo.dialogStatus==='setGroup'" label="groupList">
          <el-checkbox-group v-model="dialogTwo.temp.group_ids">
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
import { getMenuList, getMenuInfo, setGroup } from '@/api/menu'
import { getGroupList } from '@/api/group'

export default {
  name: 'MenuManage',
  data() {
    return {
      list: [],
      total: 0,
      listLoading: true,
      filter: {
        menu_id: ''
      },
      dialogTwo: {
        temp: {
          id: undefined,
          group_ids: []
        },
        dialogFormVisible: false,
        dialogStatus: '',
        groupList: []
      }
    }
  },
  created() {
    this.getList()
    this.getGroupList()
  },
  methods: {
    getList() {
      this.listLoading = true
      getMenuList(this.filter).then(response => {
        this.list = response.data

        this.listLoading = false
      })
    },
    getGroupList() {
      getGroupList(this.filter).then(response => {
        this.dialogTwo.groupList = response.data
      })
    },
    handleFilter() {
      this.getList()
    },
    handleSetGroup(row) {
      this.dialogTwo.dialogStatus = 'setGroup'
      getMenuInfo({ id: row.id }).then(response => {
        this.dialogTwo.temp = {
          id: row.id,
          group_ids: response.data.group_ids
        }
        this.dialogTwo.dialogFormVisible = true
      })
    },
    toSetGroup() {
      setGroup({
        id: this.dialogTwo.temp.id,
        group_ids: this.dialogTwo.temp.group_ids
      }).then(() => {
        this.dialogTwo.dialogFormVisible = false
        this.$notify({
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
