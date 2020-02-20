<template>
  <!-- 
    Input
      theme: BASIC(default), alt
      size: Normal(default), Small, Large
      textAlign: left(default), right, center,
      state: success, error
  -->
  <input :type="type" :class="classes" @change="$emit('change.value', $event.target.value)" >
</template>
<script>
export default {
  name: "TInput",
  model: {
    prop: 'value',
    event: 'change.value'
  },
  props: {
    theme: {
      type: String,
      default: ''
    },
    type: {
      type: String,
      default: 'text'
    },
    state: {
      type: String,
      default: null
    },
    size: {
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
      if (this.size) {
        cls += `form-control-${this.size}`
      }
      return cls
    }
  },
  methods: {
    changeHandler(e) {
      console.log(e)
      // this.$emit('change.value', this.value)
    }
  }
}
</script>