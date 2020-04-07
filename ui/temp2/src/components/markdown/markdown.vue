<template>
  <div>
    <div id="editor">{{value}}</div>
  </div>
</template>

<script>
import Vditor from 'vditor'
import '@/assets/vditor/index.scss'
import {PropTypes} from '@/utils/types'

export default {
  model: {
    prop: 'value',
    event: 'value.change'
  },
  props: {
    value: PropTypes.String
  },
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