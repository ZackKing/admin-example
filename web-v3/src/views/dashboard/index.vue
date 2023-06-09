<template>
  <div class="dashboard-container">
    <div class="dashboard-text">name: {{ name }}</div>
    <!-- <div class="dashboard-text">roles: <span v-for="role in roles" :key="role">{{ role }}</span></div> -->
  </div>
</template>

<script setup>
import { useUserStore } from '@/store/user'
import { ref } from 'vue'

const name = ref('')
const store = useUserStore()
name.value = store.name

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
