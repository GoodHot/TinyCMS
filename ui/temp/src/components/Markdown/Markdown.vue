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
      const htmlProm = this.contentEditor.getHTML()
      let html = ''
      if (typeof(htmlProm) === String) {
        html = htmlProm
      } else {
        html = this.contentEditor.vditor.lute.Md2HTML(this.contentEditor.getValue())
      }
      return {
        markdown: this.contentEditor.getValue(),
        html: html
      }
    }
  }
}
</script>