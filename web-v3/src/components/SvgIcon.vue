<template>
  <svg :class="svgClass" v-bind="$attrs" :style="{color: color}">
    <use :xlink:href="iconName" />
  </svg>
</template>

<script setup>

import { computed } from 'vue'

const props = defineProps({
  name: {
    type: String,
    required: true
  },
  color: {
    type: String,
    default: ''
  }
})

const iconName = computed(() => `#icon-${props.name}`)
const svgClass = computed(() => {
  return `w-3.5 h-3.5 fill-current align-middle ${props.name ? 'icon-' + props.name : ''}`
})
</script>