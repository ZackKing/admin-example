
<script>
import { h } from 'vue'
import SvgIcon from '@/components/SvgIcon.vue'

export default {
  name: 'MenuItem',
  functional: true,
  props: {
    icon: {
      type: String,
      default: ''
    },
    title: {
      type: String,
      default: ''
    }
  },
  render() {
    const icon = this.icon
    const title = this.title
    const vnodes = []

    if (icon) {
      vnodes.push(h(SvgIcon, { name: icon }))
    }

    if (title) {
      vnodes.push(h('span', { slot: 'title' }, title))
    }
    return vnodes
  }
}
</script>
