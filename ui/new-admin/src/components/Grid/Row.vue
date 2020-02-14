<template>
  <!--
    row:
      gutter: ''(default), 'tiny', 'no'
  -->
  <div :class="rowClass()">
    <slot></slot>
  </div>
</template>
<script>
export default {
  props: {
    gutter: {
      type: String,
      default: null
    }
  },
  methods: {
    rowClass() {
      let cls = 'row '
      if (this.gutter) {
        if (this.gutter == 'tiny') {
          cls += 'gutters-tiny '
        } else {
          cls += 'no-gutters '
        }
      }
      return cls
    }
  }
}
</script>