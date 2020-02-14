<template>
  <form class="form-row">
    <slot></slot>
  </form>
</template>
<script>
export default {
  methods: {
    validate() {
      console.log(123)
    }
  }
}
</script>