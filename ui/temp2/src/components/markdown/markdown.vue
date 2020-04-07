<template>
  <div>
    <div id="editor"></div>
  </div>
</template>

<script>
import Vditor from 'vditor'
import '@/assets/vditor/index.scss'

export default {
  mounted () {
    const options = {
      tab: '  ',
      cache: false
      // upload: {
      //   url: config.baseURL + '/upload/markdown',
      //   headers: {},
      //   format: this.uploadFormat
      // }
    }
    this.contentEditor = new Vditor('editor', options)
  }
}
</script>