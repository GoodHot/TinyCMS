<template>
  <div id="page-container" :class="classes">
    <div id="page-overlay" @click="sideHide"></div>
    <aside id="side-overlay">
      <slot name="side"></slot>
    </aside>
    <nav id="sidebar">
      <slot name="nav"></slot>
    </nav>
     <header id="page-header">
       <slot name="header"></slot>
     </header>
     <main id="main-container">
       <slot></slot>
     </main>
     <footer id="page-footer" class="bg-body-light">
       <slot name="footer"></slot>
     </footer>
  </div>
</template>
<script>
import {PropTypes} from '@/utils/types'

export default {
  props: {
    sidebarShow: PropTypes.Boolean.def(true),
    sidebarPosition: PropTypes.String.def('left'),
    sidebarMin: PropTypes.Boolean,
    sidebarTheme: PropTypes.String.def('light'),
    sideVisiable: PropTypes.Boolean,
    sideHoverMode: PropTypes.Boolean,
    headerMode: PropTypes.String.def('fixed'),
    headerGlass: PropTypes.Boolean,
    headerTheme: PropTypes.String.def('dark'),
    footerMode: PropTypes.String.def('static'),
    contentMode: PropTypes.String.def('narrow')
  },
  data () {
    return {
      sideShow: false
    }
  },
  methods: {
    sideHide () {
      this.sideShow = false
      this.$emit('onSideHide')
    }
  },
  mounted () {
    this.sideShow = JSON.parse(JSON.stringify(this.sideVisiable))
  },
  computed: {
    classes () {
      return [
        'enable-page-overlay side-trans-enabled',
        {
          'sidebar-o': this.sidebarShow,
          'sidebar-r': this.sidebarPosition === 'right',
          'sidebar-mini': this.sidebarMin,
          'sidebar-dark': this.sidebarTheme === 'dark',
          'side-overlay-o': this.sideShow,
          'side-overlay-hover': this.sideHoverMode,
          'page-header-fixed': this.headerMode === 'fixed',
          'page-header-glass': this.headerGlass,
          'page-header-dark': this.headerTheme === 'dark',
          'page-footer-fixed': this.footerMode === 'fixed',
          'main-content-boxed': this.contentMode === 'boxed',
          'main-content-narrow': this.contentMode === 'narrow'
        }
      ]
    }
  },
  watch: {
    sideVisiable (val) {
      this.sideShow = JSON.parse(JSON.stringify(val))
    }
  }
}
</script>