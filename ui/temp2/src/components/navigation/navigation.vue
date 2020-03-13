<template>
  <div :class="classes">
    <div class="d-lg-none">
      <b-button block variant="light" class="d-flex justify-content-between align-items-center">Menu - Hover Normal <t-icon icon="bars" /></b-button>
    </div>
    <div class="d-none d-lg-block mt-2 mt-lg-0">
      <ul :class="ulClasses">
        <t-navigation-item
          ref="navItems"
          v-for="(item, index) in data"
          :key="index"
          :active="item.active"
          :href="item.href"
          :to="item.to"
          :title="item.title"
          :badgeText="item.badgeText"
          :badgeTheme="item.badgeTheme"
          :icon="item.icon"
          :iconPack="item.iconPack"
          :child="item.child"
          :showEvent="showEvent"
          :horizontal="horizontal"
          @onopen="onopenHandler"
        >
        </t-navigation-item>
      </ul>
    </div>
  </div>
</template>
<script>
import {PropTypes} from '@/utils/types'

export default {
  props: {
    horizontal: PropTypes.Boolean.def(true),
    rounded: PropTypes.Boolean.def(true),
    data: PropTypes.Array,
    theme: PropTypes.String.def('light'),
    layout: PropTypes.String.def('left'),
    showEvent: PropTypes.String.def('hover')
  },
  computed: {
    classes () {
      return [
        'push',
        {
          'rounded': this.rounded,
          'bg-sidebar-dark': this.theme === 'dark',
          'bg-white': this.theme === 'light',
        }
      ]
    },
    ulClasses () {
      return [
        'nav-main',
        {
          'nav-main-hover' : this.showEvent === 'hover',
          'nav-main-horizontal': this.horizontal,
          'nav-main-dark': this.theme === 'dark',
          'nav-main-horizontal-center': this.layout === 'center',
          'nav-main-horizontal-justify': this.layout === 'full'
        }
      ]
    }
  },
  methods: {
    onopenHandler (currItem) {
      this.$refs.navItems.map(item => {
        if (currItem !== item) {
          item.$close()
        }
      })
      this.$emit('onopen', currItem)
    }
  }
}
</script>