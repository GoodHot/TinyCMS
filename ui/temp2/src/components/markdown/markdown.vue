<template>
  <div>
    <div id="editor">{{ value }}</div>
  </div>
</template>

<script>
import Vditor from 'vditor'
import '@/assets/vditor/index.scss'
import {ACCESS_TOKEN, config} from '@/utils/request'
import {PropTypes} from '@/utils/types'
import Vue from 'vue'

export default {
  model: {
    prop: 'value',
    event: 'value.change'
  },
  props: {
    value: PropTypes.String
  },
  data () {
    return {
      markdown: null,
      contentEditor: null
    }
  },
  mounted () {
    this.initMarkdown()
  },
  methods: {
    uploadFormat(files, response) {
      const respJson = JSON.parse(response)
      const result = {
        msg: respJson.msg,
        code: 0,
        data: {
          errFiles: [],
          succMap: {}
        }
      }
      for (let i in respJson.data.path) {
        const file = respJson.data.path[i]
        const name = file.path.substring(file.path.lastIndexOf('/') + 1)
        result.data.succMap[name] = this.assetsURL(file.url)
      }
      return JSON.stringify(result)
    },
    initMarkdown () {
      const options = {
        tab: '  ',
        cache: {
          enable: false
        },
        upload: {
          url: config.baseURL + '/upload/markdown',
          headers: {},
          format: this.uploadFormat
        }
      }
      options.upload.headers[ACCESS_TOKEN] = Vue.ls.get(ACCESS_TOKEN)
      this.contentEditor = new Vditor('editor', options)
    },
    getValue () {
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
  },
  watch: {
    value(val) {
      this.contentEditor.setValue(val)
    }
  }
}
</script>