<template>
  <div class="app-container">
    <div class="container-content">
      <el-row>
        Welcom {{ name }}
      </el-row>
    </div>
  </div>
</template>


<script setup>
import { useUserStore } from '~/store/user'
import { ref } from 'vue'

const name = ref('')
const store = useUserStore()
name.value = store.nickname

</script>

<style lang="scss" scoped>
.dashboard {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>
