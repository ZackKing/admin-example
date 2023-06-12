
<template>
  <component :is="is" v-if="isExt(to)" :href="to" target="_blank" rel="noopener">
    <slot />
  </component>
  <component :is="is" v-else :to="to">
    <slot />
  </component>
  <!-- <component v-bind="linkProps(to)">
    <slot />
  </component> -->
</template>

<script>
import { isExternal } from '@/utils'

export default {
  props: {
    to: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      is: ''
    }
  },
  methods: {
    isExt(to) {
      const is = isExternal(to)
      this.is = is ? 'a' : 'router-link'
      return is
    }
    // linkProps(url) {
    //   if (isExternal(url)) {
    //     return {
    //       is: 'a',
    //       href: url,
    //       target: '_blank',
    //       rel: 'noopener'
    //     }
    //   }
    //   return {
    //     is: 'router-link',
    //     // to: '/redirect' + url
    //     to: url
    //   }
    // }
  }
}
</script>
