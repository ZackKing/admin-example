<template>
  <el-select v-model="ppValue" multiple filterable :collapse-tags="collapseTags" @change="handleChange">
    <el-option :label="'All(' + list.length + ')'" value="_all_" />
    <el-option v-for="option in formatList" :key="option.value" :label="option.label" :value="option.value" />
  </el-select>
</template>

<script>
export default {
  name: 'PpSelect',
  props: {
    value: {
      type: Array,
      default: () => {
        return []
      }
    },
    list: {
      type: Array,
      default: () => {
        return []
      }
    },
    collapseTags: {
      type: Boolean,
      default: true
    },
    valueName: {
      type: String,
      default: 'value'
    },
    labelName: {
      type: String,
      default: 'lable'
    }
  },
  emits: ['input'],
  data() {
    return {
      ppValue: this.value.length ? this.value : ['_all_']
    }
  },
  computed: {
    formatList: function () {
      const _this = this
      return this.list.map(function (item) {
        if (typeof item !== 'object') {
          return {
            value: item,
            label: item
          }
        } else if ((item instanceof Object) && !(item instanceof Array) && !(item instanceof Function)) {
          const valueName = _this.valueName
          const labelName = _this.labelName
          return {
            value: item[valueName],
            label: item[labelName]
          }
        } else {
          alert('请传入正确的list')
          return {}
        }
      })
    }
  },
  methods: {
    handleChange(options) {
      if (options.length) {
        const newOptions = options[options.length - 1] !== '_all_' ? options.filter((item) => {
          return item !== '_all_'
        }) : ['_all_']

        this.ppValue = newOptions
        this.$emit('input', newOptions[0] === '_all_' ? [] : newOptions)
      } else {
        const newOptions = ['_all_']
        this.ppValue = newOptions
        this.$emit('input', [])
      }
    }
  }
}
</script>
