<template>
  <div id="vditor" @click="clickHandler"></div>
</template>
<script>
import '../../assets/vditor/index.scss'
import Vditor from 'vditor'

export default {
  data() {
    return {
      markdown: null,
      contentEditor: null
    }
  },
  mounted() {
    this.initMarkdown()
  },
  methods: {
    initMarkdown() {
      this.contentEditor = new Vditor('vditor', {
        tab: '  ',
        cache: false,
      })
    },
    clickHandler(e) {
      this.$emit('click', e)
    },
    getValue() {
      return {
        markdown: this.contentEditor.getValue(),
        html: this.contentEditor.getHTML()
      }
    }
  }
}
</script>