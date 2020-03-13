<template>
  <div>
    <div :class="headerClasses">
      <h3 class="block-title" v-if="title || subtitle">
        {{title}}
        <small v-if="subtitle">{{subtitle}}</small>
      </h3>
      <div class="block-options" v-if="$slots.options">
        <slot name="options"></slot>
      </div>
    </div>
    <div class="block-content" v-if="contentVisible">
      <slot></slot>
    </div>
    <div v-if="$slots.footer" class="block-content block-content-full block-content-sm bg-body-light font-size-sm">
      <slot name="footer"></slot>
    </div>
    <div class="block-loading-mode" v-if="loading">
    </div>
  </div>
</template>
<script>
import {PropTypes} from '@/utils/types'

export default {
  props: {
    title: PropTypes.String,
    subtitle: PropTypes.String,
    border: PropTypes.Boolean,
    rounded: PropTypes.Boolean,
    transparent: PropTypes.Boolean,
    titleRight: PropTypes.Boolean,
    theme: PropTypes.String.def('none'), // 'primary-light', 'primary-dark', 'primary-darker', 'gd-dusk', 'gd-fruit', 'gd-aqua', 'gd-sublime', 'gd-sea', 'gd-leaf', 'gd-lake', 'gd-sun', 'default', 'default-light', 'default-dark', 'default-darker', 'xwork', 'xwork-light', 'xwork-dark', 'xwork-darker', 'xmodern', 'xmodern-light', 'xmodern-dark', 'xmodern-darker', 'xeco', 'xeco-light', 'xeco-dark', 'xeco-darker', 'xsmooth', 'xsmooth-light', 'xsmooth-dark', 'xsmooth-darker', 'xinspire', 'xinspire-light', 'xinspire-dark', 'xinspire-darker', 'success', 'info', 'warning', 'danger', 'gray', 'muted', 'gray-darker', 'black', 
    href: PropTypes.String,
    linkTheme: PropTypes.String, // 'rotate', 'pop', 'shadow'
    contentVisible: PropTypes.Boolean.def(true),
    loading: PropTypes.Boolean
  },
  computed: {
    headerClasses() {
      return [
        'block-header',
        {
          'block-header-rtl': this.titleRight,
          [`bg-${this.theme}`]: this.theme && this.theme != 'none',
          'block-header-default': !this.theme
        }
      ]
    }
  }
}
</script>