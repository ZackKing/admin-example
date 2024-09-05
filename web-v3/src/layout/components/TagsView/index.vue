<template>
  <div class="tags-view-container">
    <scroll-pane ref="scrollPane" class="tags-view-wrapper">
      <router-link
        v-for="tag in visitedViews"
        ref="tag"
        :key="tag.path"
        :class="`tags-view-item ${isActive(tag) ? 'active' : ''}`"
        :to="{ path: tag.path, query: tag.query, fullPath: tag.fullPath }"
        @contextmenu.prevent.enter="openMenu(tag, $event)"
      >
        {{ tag.title }}
        <el-icon v-if="isActive(tag)" class="h-px" ><CloseBold @click="closeSelectedTag(tag)" /></el-icon>
      </router-link>
    </scroll-pane>
    <ul v-show="visible" :style="{left:left+'px',top:top+'px'}" class="contextmenu">
      <li @click="closeSelectedTag(selectedTag)">Close</li>
      <li @click="closeOthersTags">Close Other</li>
      <li @click="closeAllTags(selectedTag)">Close All</li>
    </ul>
  </div>
</template>

<script>
import { CloseBold } from '@element-plus/icons-vue'
import ScrollPane from './ScrollPane.vue'
import path from 'path-browserify'

export default {
  components: { ScrollPane, CloseBold },
  data() {
    return {
      visible: false,
      top: 0,
      left: 0,
      selectedTag: {},
      affixTags: []
    }
  },
  computed: {
    visitedViews() {
      return this.$store.tagsView.visitedViews
    },
    routes() {
      return this.$store.permission.routes
    }
  },
  watch: {
    $route() {
      this.addTags()
      this.moveToCurrentTag()
    },
    visible(value) {
      if (value) {
        document.body.addEventListener('click', this.closeMenu)
      } else {
        document.body.removeEventListener('click', this.closeMenu)
      }
    }
  },
  mounted() {
    this.initTags()
    this.addTags()
  },
  methods: {
    isActive(route) {
      return route.path === this.$route.path
    },
    isAffix(tag) {
      return tag.meta && tag.meta.affix
    },
    filterAffixTags(routes, basePath = '/') {
      let tags = []
      routes.forEach(route => {
        if (route.meta && route.meta.affix) {
          const tagPath = path.resolve(basePath, route.path)
          tags.push({
            fullPath: tagPath,
            path: tagPath,
            name: route.name,
            meta: { ...route.meta }
          })
        }
        if (route.children) {
          const tempTags = this.filterAffixTags(route.children, route.path)
          if (tempTags.length >= 1) {
            tags = [...tags, ...tempTags]
          }
        }
      })
      return tags
    },
    initTags() {
      const affixTags = this.affixTags = this.filterAffixTags(this.routes)
      for (const tag of affixTags) {
        if (tag.name) { // Must have tag name
          this.$store.tagsView.addVisitedView(tag)
        }
      }
    },
    addTags() {
      const { name } = this.$route
      if (name) {
        this.$store.tagsView.addView(this.$route)
      }
      return false
    },
    moveToCurrentTag() {
      const tags = this.$refs.tag
      if (tags) {
        this.$nextTick(() => {
          for (const tag of tags) {
            if (tag.to.path === this.$route.path) {
              this.$refs.scrollPane.moveToTarget(tag)
              if (tag.to.fullPath !== this.$route.fullPath) { // when query is different then update
                this.$store.tagsView.updateVisitedView(this.$route)
              }
              break
            }
          }
        })
      }
    },
    refreshSelectedTag(view) {
      this.$store.tagsView.delCachedView(view)
      const { fullPath } = view
      this.$nextTick(() => {
        this.$router.replace({
          path: '/redirect' + fullPath
        })
      })
    },
    closeSelectedTag(view) {
      this.$store.tagsView.delView(view)
      if (this.isActive(view)) {
        this.toLastView(this.$store.tagsView.visitedViews, view)
      }
    },
    closeOthersTags() {
      this.$router.push(this.selectedTag)
      this.$store.tagsView.delOthersViews(this.selectedTag)
      this.moveToCurrentTag()
    },
    closeAllTags(view) {
      const { visitedViews } = this.$store.tagsView.delAllViews()
      if (this.affixTags.some(tag => tag.path === view.path)) {
        return
      }
      this.toLastView(visitedViews, view)
    },
    toLastView(visitedViews, view) {
      const latestView = visitedViews.slice(-1)[0]
      if (latestView) {
        this.$router.push(latestView.fullPath)
      } else {
        if (view.name === 'Dashboard') {
          this.$router.replace({ path: '/redirect' + view.fullPath })
        } else {
          this.$router.push('/')
        }
      }
    },
    openMenu(tag, e) {
      const menuMinWidth = 105
      const offsetLeft = this.$el.getBoundingClientRect().left
      const offsetWidth = this.$el.offsetWidth
      const maxLeft = offsetWidth - menuMinWidth
      const left = e.clientX - offsetLeft + 15

      if (left > maxLeft) {
        this.left = maxLeft
      } else {
        this.left = left
      }

      this.top = e.clientY
      this.visible = true
      this.selectedTag = tag
    },
    closeMenu() {
      this.visible = false
    }
  }
}
</script>

<style lang="scss" scoped>
.tags-view-container {
  height: 34px;
  width: 100%;
  background: #fff;
  border-bottom: 1px solid #d8dce5;
  box-shadow: 0 1px 3px 0 rgba(0, 0, 0, .12), 0 0 3px 0 rgba(0, 0, 0, .04);
  .tags-view-wrapper {
    .tags-view-item {
      display: inline-block;
      position: relative;
      cursor: pointer;
      height: 26px;
      min-width: 60px;
      line-height: 25px;
      border-radius: 3px;
      border: 1px solid #efefef;
      color: #97a8be;
      background: #fff;
      padding: 0 8px;
      font-size: 12px;
      text-align: center;
      margin-left: 5px;
      margin-top: 4px;
      &:first-of-type {
        margin-left: 15px;
      }
      &:last-of-type {
        margin-right: 15px;
      }
      &.active {
        background-color: #42b983;
        color: #fff;
        border-color: #42b983;
        &::before {
          content: '';
          background: #fff;
          display: inline-block;
          width: 8px;
          height: 8px;
          border-radius: 50%;
          position: relative;
          margin-right: 2px;
        }
      }
    }
  }
  .contextmenu {
    background: #fff;
    z-index: 3000;
    position: absolute;
    padding: 5px 0;
    border-radius: 4px;
    font-size: 12px;
    font-weight: 400;
    color: #333;
    box-shadow: 2px 2px 3px 0 rgba(0, 0, 0, .3);
    li {
      padding: 7px 16px;
      &:hover {
        background: #eee;
      }
    }
  }
}
</style>
