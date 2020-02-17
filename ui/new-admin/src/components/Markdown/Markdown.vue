<template>
  <div id="vditor"></div>
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
      // this.markdown = new Editor({
      //   el: this.$refs.markdownEditor,
      //   initialEditType: 'markdown',
      //   previewStyle: 'vertical',
      //   height: '700px',
      //   exts: ['scrollSync'],
      //   initialValue: initValue,
      //   // hooks: {
      //   //   addImageBlobHook (blob, callback) {
      //   //     _this.imageUpload(blob, callback)
      //   //   }
      //   // }
      // })
    }
  }
}
</script>