<template>
  <!-- 
    Input
      theme: BASIC(default), alt
      size: Normal(default), Small, Large
      textAlign: left(default), right, center,
      state: success, error
  -->
  <textarea :class="classes" ></textarea>
</template>
<script>
export default {
  props: {
    theme: {
      type: String,
      default: ''
    },
    state: {
      type: String,
      default: null
    }
  },
  computed: {
    classes() {
      let cls = 'form-control '
      if (this.theme === 'alt') {
        cls += 'form-control-alt '
      }
      if (this.state) {
        if (this.state === 'success') {
          cls += 'is-valid '
        } else {
          cls += 'is-invalid '
        }
      }
      return cls
    }
  }
}
</script>