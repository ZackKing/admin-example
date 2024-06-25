<template>
  <SvgIcon v-if="props.icon" :name="props.icon" />
  <span v-if="props.title" class="ml-1">{{ props.title }}</span>
</template>

<script setup>
import SvgIcon from '~/components/SvgIcon.vue'

const props = defineProps({
  icon: {
    type: String,
    default: ''
  },
  title: {
    type: String,
    default: ''
  }
})
</script>
