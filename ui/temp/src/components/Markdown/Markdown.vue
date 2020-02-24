<template>
  <div id="vditor" @click="clickHandler">{{value}}</div>
</template>
<script>
import '../../assets/vditor/index.scss'
import {ACCESS_TOKEN, config} from '@/utils/request'
import Vditor from 'vditor'
import Vue from 'vue'

export default {
  model: {
    prop: 'value',
    event: 'value.change'
  },
  props: {
    value: {
      type: String,
      default: ''
    }
  },
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
    uploadFormat(files, response) {
      console.log('fmt')
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
    initMarkdown() {
      const options = {
        tab: '  ',
        cache: false,
        upload: {
          url: config.baseURL + '/upload/markdown',
          headers: {},
          format: this.uploadFormat
        }
      }
      options.upload.headers[ACCESS_TOKEN] = Vue.ls.get(ACCESS_TOKEN)
      this.contentEditor = new Vditor('vditor', options)
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