<template>
  <div class="app-container">
    <div class="container-header">
      <el-form :inline="true" class="header-filter">
        <el-form-item >
          <el-button :disabled="loading" type="primary" :icon="RefreshIcon" @click="refresh">Refresh</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="container-content">
      <el-table v-loading="loading" :data="list" border fit highlight-current-row default-expand-all row-key="name"
        :tree-props="{ children: 'sub_menu', hasChildren: 'hasChildren' }">
        <el-table-column prop="name" label="Name" width="200" />
        <el-table-column prop="uri" label="Url" width="110" />
        <el-table-column prop="icon" label="Icon" width="110">
          <template #default="{ row }">
            <SvgIcon class="text-gray-500" :name="row.icon" />
          </template>
        </el-table-column>
        <el-table-column label="Status" width="110">
          <template #default="{ row }">
            <el-tag v-if="row.status === 1">valid</el-tag>
            <el-tag v-else type="info">invalid</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="remark" label="Remark" />
        <el-table-column prop="created_time" label="Created" width="150" />
        <el-table-column prop="updated_time" label="Updated" width="150" />
      </el-table>
    </div>
  </div>
</template>

<script setup>
import { getMenuList } from '@/api/menu'
import SvgIcon from '@/components/SvgIcon.vue'
import { onMounted, ref } from 'vue'
import { Refresh as RefreshIcon } from '@element-plus/icons-vue'

const list = ref([])
const loading = ref(true)

onMounted(async () => {
  await refresh()
})

async function refresh() {
  loading.value = true
  const res = await getMenuList()
  list.value = res.data
  loading.value = false
}

</script>
