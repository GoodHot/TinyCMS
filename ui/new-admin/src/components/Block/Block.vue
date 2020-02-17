<template>
  <!--
    block
      border: true/false
      head-bg: true/false
      border-round: true/false
      shadow: true/false
      transparent: true/false
      theme: ''(default),'primary','primary-light','primary-dark','primary-darker','success','info','warning','danger','gray','gray-dark','gray-darker','black','default','default-light','default-dark','default-darker','amethyst','amethyst-light','amethyst-dark','amethyst-darker','city','city-light','city-dark','city-darker','flat','flat-light','flat-dark','flat-darker','modern','modern-light','modern-dark','modern-darker','smooth','smooth-light','smooth-dark','smooth-darker'
      title: string
      subtitle: string
      icon: string
      iconPack: stirng
  -->
  <div :class="classes">
    <div :class="headerClass" v-if="title || subtitle">
      <h3 class="block-title">
        <TIcon :icon="icon" :pack="iconPack" v-if="icon" />
        {{ title }}
        <small v-if="subtitle">{{ subtitle }}</small>
      </h3>
      <div :class="`block-options ${contentClass}`">
        <slot name="options"></slot>
      </div>
    </div>
    <div class="block-content">
      <slot></slot>
    </div>
  </div>
</template>
<script>
export default {
  props: {
    border: {
      type: Boolean,
      default: false
    },
    headerBg: {
      type: Boolean,
      default: true
    },
    borderRound: {
      type: Boolean,
      default: false
    },
    shadow: {
      type: Boolean,
      default: false
    },
    transparent: {
      type: Boolean,
      default: false
    },
    theme: {
      type: String,
      default: null
    },
    title: {
      type: String,
      default: null
    },
    subtitle: {
      type: String,
      default: null
    },
    icon: {
      type: String,
      default: null
    },
    iconPack: {
      type: String,
      default: 'fa'
    },
    contentClass: {
      type: String,
      default: ''
    }
  },
  created() {
    this.renderOptions()
  },
  beforeUpdate() {
    this.renderOptions()
  },
  computed: {
    classes() {
      let cls = 'block '
      if (this.transparent) {
        return cls + 'block-transparent'
      }
      if (this.theme) {
        cls += 'block-themed '
      }
      if (this.border) {
        cls += 'block-bordered '
      }
      if (this.borderRound) {
        cls += 'block-rounded '
      }
      if (this.shadow) {
        cls += 'block-fx-shadow '
      }
      return cls
    },
    headerClass() {
      let cls = 'block-header '
      if (this.transparent) {
        return cls
      }
      if (this.headerBg && !this.theme) {
        return cls + 'block-header-default'
      }
      if (this.theme) {
        cls += `bg-${this.theme} `
      }
      return cls
    }
  },
  methods: {
    renderOptions() {
      const opts = this.$slots.options
      if (!opts) {
        return
      }
      opts.forEach(btn => {
        if (!btn.componentOptions || btn.componentOptions.tag !== 'TButton') {
          return
        }
        btn.componentOptions.propsData.isBlockOption = true
      })
    },
  }
}
</script>