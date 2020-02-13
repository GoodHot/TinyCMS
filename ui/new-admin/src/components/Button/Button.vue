<template>
  <!-- 
    按钮颜色分为：
    primary, secondary, success, info, warning, danger, dark, light
    按钮风格分为：
    default, alt, outline, 
    按钮大小分为：
    default, lg, sm
    按钮border-radio分为
    default, square, rounded
    按钮状态分为
    default, DISABLED
    按钮赋能
    带有图标的按钮
      按钮位置：before, after
  -->
  <button type="button" :class="btnClass()" @click="click">
    <TIcon :icon="icon" :pack="iconPack" class="mr-1 fa-fw" v-if="icon && iconPosition === 'before'" />
    <slot></slot>
    <TIcon :icon="icon" :pack="iconPack" class="ml-1 fa-fw" v-if="icon && iconPosition === 'after'" />
  </button>
</template>
<script>
export default {
  name: 'TButton',
  props: {
    type: {
      type: String, // default, primary, secondary, success, info, warning, danger, dark, light
      default: 'primary'
    },
    theme: {
      type: String, // default, alt, outline
      default: ''
    },
    size: {
      type: String, // default, lg, sm
      default: ''
    },
    border: {
      type: String, // default, square, rounded
      default: ''
    },
    block: {
      type: String , // default = '', value ==> ['','block']
      default: ''
    },
    iconPack: {
      type: String,
      default: 'fa'
    },
    iconPosition: {
      type: String,
      default: 'before'
    },
    icon: {
      type: String,
      default: null
    }
  },
  methods: {
    btnClass() {
      let cls = 'btn '
      if (this.theme !== '' && (this.theme === 'alt' || this.theme === 'outline')) {
        cls += `btn-${this.theme}-${this.type} `
      } else {
        cls += `btn-${this.type} `
      }
      if (this.size !== '') {
        cls += `btn-${this.size} `
      }
      if (this.border !== '') {
        cls += `btn-${this.border} `
      }
      if (this.block !== '') {
        cls += 'btn-block '
      }
      return cls
    },
    click(e) {
      this.$emit('click', e)
    }
  }
}
</script>