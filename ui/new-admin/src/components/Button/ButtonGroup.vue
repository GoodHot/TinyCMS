<template>
  <!--
    ButtonGroup
      type: primary(default), secondary, success, info, warning, danger, dark, light
      direction: HORIZONTAL(default), VERTICAL
      size: Normal(default), Small, Large
      theme: Normal(default), alt, outline, 
  -->
  <div :class="classes">
    <slot></slot>
  </div>
</template>
<script>
export default {
  props: {
    type: {
      type: String, // primary(default), secondary, success, info, warning, danger, dark, light
      default: 'primary'
    },
    direction: {
      type: String, // hr, ve
      default: 'hr'
    },
    size: {
      type: String, // sm, lg
      default: ''
    },
    theme: {
      type: String, // alt, outline
      default: ''
    }
  },
  created() {
    this.renderChild()
  },
  computed: {
    classes() {
      let cls = ''
      if (this.direction === '' || this.direction === 'hr') {
        cls = 'btn-group '
      } else {
        cls = 'btn-group-vertical '
      }
      if (this.size !== '') {
        cls += `btn-group-${this.size} `
      }
      return cls
    }
  },
  methods: {
    renderChild() {
      const btns = this.$slots.default
      if (!btns) {
        return 
      }
      btns.forEach(btn => {
        if (btn.componentOptions.tag !== 'TButton') {
          return
        }
        btn.componentOptions.propsData.type = this.type
        btn.componentOptions.propsData.theme = this.theme
      })
    }
  }
}
</script>