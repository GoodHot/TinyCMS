<template>
  <div ref="markdownEditor">
  </div>
</template>
<script>
import 'codemirror/lib/codemirror.css'
import 'tui-editor/dist/tui-editor.css'
import 'tui-editor/dist/tui-editor-contents.css'
import 'highlight.js/styles/github.css'
import Editor from 'tui-editor'

export default {
  data() {
    return {
      markdown: null
    }
  },
  mounted() {
    this.initMarkdown('')
  },
  methods: {
    initMarkdown(initValue) {
      this.markdown = new Editor({
        el: this.$refs.markdownEditor,
        initialEditType: 'markdown',
        previewStyle: 'vertical',
        height: '500px',
        exts: ['scrollSync'],
        initialValue: initValue,
        // hooks: {
        //   addImageBlobHook (blob, callback) {
        //     _this.imageUpload(blob, callback)
        //   }
        // }
      })
    }
  }
}
</script>