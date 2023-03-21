
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useTagsViewStore = defineStore('tagsView', () => {
  const visitedViews = ref([])
  const cachedViews = ref([])

  /**
   * @param {object} view
   * @param {int} handle -11: del all, -1: del, 0: add, 1: update
   */
  function _handleVisitedView(view, handle = 0) {
    switch (handle) {
      case -1: // del
        for (const [i, v] of visitedViews.value.entries()) {
          if (v.path === view.path) {
            visitedViews.value.splice(i, 1)
            break
          }
        }
        break
      case -11:
        const affixTags = visitedViews.value.filter(tag => tag.meta.affix)
        visitedViews.value = affixTags
        break
      case -12:
        visitedViews.value = visitedViews.value.filter(v => {
          return v.meta.affix || v.path === view.path
        })
        break
      case 0: // add
        if (visitedViews.value.some(v => v.path === view.path)) {
          return
        }
        visitedViews.value.push(
          Object.assign({}, view, {
            title: view.meta.title || 'no-name'
          })
        )
        break
      case 1: // update
        for (let v of visitedViews.value) {
          if (v.path === view.path) {
            v = Object.assign(v, view)
            break
          }
        }
        break
      default:
        break
    }
  }

  /**
   * @param {*} view
   * @param {int} handle 0: add, -1: del, -11: del all, -12: del other
   */
  function _handleCachedView(view, handle = 0) {
    switch (handle) {
      case -1: // del
        const index = cachedViews.value.indexOf(view.name)
        index > -1 && cachedViews.value.splice(index, 1)
        break
      case -11: // del all
        cachedViews.value = []
        break
      case -12: // del other
        const i = cachedViews.value.indexOf(view.name)
        if (i > -1) {
          cachedViews.value = cachedViews.value.slice(i, i + 1)
        } else { // if i = -1, there is no cached tags
          cachedViews.value = []
        }
        break
      case 0: // add
        if (cachedViews.value.includes(view.name)) {
          return
        }
        if (!view.meta.noCache) {
          cachedViews.value.push(view.name)
        }
        break
      default:
        break
    }
  }

  function addVisitedView(view) {
    _handleVisitedView(view, 0)
  }
  function delVisitedView(view) {
    _handleVisitedView(view, -1)
    return [...visitedViews]
  }

  function addCachedView(view) {
    _handleCachedView(view, 0)
  }
  function delCachedView(view) {
    _handleCachedView(view, -1)
    return [...cachedViews]
  }

  function addView(view) {
    _handleVisitedView(view, 0)
    _handleCachedView(view, 0)
  }

  function delView(view) {
    delVisitedView(view)
    delCachedView(view)
  }

  function delOthersViews(view) {
    _handleVisitedView(view, -12)
    _handleCachedView(view, -12)
  }

  function delAllViews(view) {
    _handleVisitedView(view, -11)
    _handleCachedView(view, -11)
  }

  function updateVisitedView(view) {
    _handleVisitedView(view, 2)
  }

  return {
    visitedViews,
    cachedViews,
    addView,
    delView,
    delAllViews,
    delOthersViews,
    addVisitedView,
    delVisitedView,
    addCachedView,
    delCachedView,
    updateVisitedView,
  }
})
