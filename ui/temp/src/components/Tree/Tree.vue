<template>
  <div>
    <TTreeChild v-for="item of data" :key="item.name" :data="item" :expand="expand" @select="selectHandler"></TTreeChild>
  </div>
</template>
<script>
export default {
  props: {
    data: {
      type: Array,
      default() {
        return []
      }
    },
    expand: {
      type: Boolean,
      default: false
    }
  },
  methods: {
    selectHandler(data) {
      this.$emit('select', data)
    }
  }
}
</script>