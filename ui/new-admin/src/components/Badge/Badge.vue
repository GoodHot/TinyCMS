<template>
  <!-- 
    Badge:
      type: primary(default), secondary, success, danger, warning, info, light, dark
      theme: ''(default), pill
  -->
  <span :class="classes"><slot></slot></span>
</template>
<script>
export default {
  props: {
    type: {
      type: String,
      default: 'primary'
    },
    theme: {
      type: String,
      default: ''
    }
  },
  computed: {
    classes() {
      let cls = `badge badge-${this.type} `
      if (this.theme !== '') {
        cls += `badge-${this.theme}`
      }
      return cls
    }
  }
}
</script>