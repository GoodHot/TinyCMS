<template>
  <!-- 
    菜单元素：
    {
      title = ''
      icon = {
        icon = ''
        pack = ''
      }
      badge = {

      }
      href = '' // a href link
      to = '' // router link
      children = []
    }
   -->
  <div :class="classes">
    <div class="d-lg-none">
      <b-button @click="showHorizontalNavigation" block variant="light">TinyCMS</b-button>
    </div>
    <div ref="horizontalNavigation" class="d-none d-lg-block mt-2 mt-lg-0">
      <div class="t-nav-between">
        <ul :class="'nav-main nav-main-horizontal nav-main-hover ' + (theme === 'dark' ? 'nav-main-dark' : '')" v-if="data">
          <TNavigationItem v-for="item of data" :key="item.title + item.icon" :data="item" />
        </ul>
        <span>
          <slot name="right"></slot>
        </span>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  props: {
    theme: {
      type: String,
      default: "light"
    },
    data: {
      type: Array,
      default () {
        return []
      }
    }
  },
  data() {
    return {
      horizontalNavigationDisplay: false
    }
  },
  computed: {
    classes() {
      if (this.theme === "light") {
        return "bg-white p-2 push"
      } else {
        return "bg-sidebar-dark p-2 push"
      }
    }
  },
  methods: {
    showHorizontalNavigation() {
      const nav = this.$refs.horizontalNavigation
      const clsList = nav.classList
      if (clsList.contains('d-none')) {
        clsList.remove('d-none')
      } else {
        clsList.add('d-none')
      }
    }
  }
}
</script>