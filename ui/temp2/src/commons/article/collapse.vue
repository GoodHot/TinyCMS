<template>
<div class="block mb-0 t-article-collapse">
  <div class="block-header block-header-default p-0" role="tab" id="settings-accordion_h1">
    <a class="font-w600 d-block w-100 p-3" href="javascript:void(0)" @click="showHandler">{{ title }}</a>
  </div>
  <b-collapse :visible="visible">
    <div class="block-content">
      <slot></slot>
    </div>
  </b-collapse>
</div>
</template>
<script>
import { PropTypes } from '@/utils/types'

export default {
  props: {
    show: PropTypes.Boolean,
    title: PropTypes.String
  },
  data () {
    return {
      visible: false
    }
  },
  mounted () {
    this.visible = JSON.parse(JSON.stringify(this.show))
  },
  methods: {
    showHandler () {
      this.visible = !this.visible
      if (this.visible) {
        this.$emit('onshow')
      }
    }
  },
  watch: {
    show (val) {
      this.visible = JSON.parse(JSON.stringify(val))
    }
  }
}
</script>