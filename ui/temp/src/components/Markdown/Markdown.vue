<template>
  <div id="vditor" @click="clickHandler"></div>
</template>
<script>
import '../../assets/vditor/index.scss'
import Vditor from 'vditor'

export default {
  data() {
    return {
      markdown: null
    }
  },
  mounted() {
    this.initMarkdown()
  },
  methods: {
    initMarkdown() {
      new Vditor('vditor')
    },
    clickHandler(e) {
      this.$emit('click', e)
    }
  }
}
</script>