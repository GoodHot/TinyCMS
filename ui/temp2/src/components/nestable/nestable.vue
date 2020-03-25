<template>
<vue-nestable v-model="nestableItems" @input="inputHandler" @change="changeHandler" ref="nestable">
  <vue-nestable-handle slot-scope="{ item }" :item="item">
    <b-button class="expand" variant="none" v-if="item.children && item.children.length > 0" @click="expandHandler(item)">{{ item.expand }}</b-button>
    <slot v-bind:item="item"></slot>
  </vue-nestable-handle>
</vue-nestable>
</template>
<script>
import { VueNestable, VueNestableHandle } from 'vue-nestable'
import '@/assets/css/nestable.scss'
import {PropTypes} from '@/utils/types'

export default {
  model: {
    prop: 'data',
    event: 'data.change'
  },
  components: {
    VueNestable,
    VueNestableHandle
  },
  props: {
    data: PropTypes.Array
  },
  data () {
    return {
      nestableItems: []
    }
  },
  created () {
    const value = JSON.parse(JSON.stringify(this.data))
    this.initVal(value)
    this.nestableItems = value
  },
  methods: {
    initVal (val) {
      val.forEach(v => {
        if (v.children && v.children.length > 0) {
          v.hide = false
          v.expand = '-'
          this.initVal(v.children)
        } else {
          v.hide = false
          v.expand = '-'
          v.children = []
        }
      });
    },
    inputHandler (val) {
      this.$emit('data.change', val)
    },
    changeHandler (val, opt) {
      console.log(val, opt)
    },
    expandHandler (item) {
      const nestable = this.$refs.nestable
      const child = nestable.$el.querySelector(`.nestable-item-${item.id} .nestable-list`)
      if (item.hide) {
        child.style.display = 'block'
        item.hide = false
        item.expand = '-'
      } else {
        child.style.display = 'none'
        item.hide = true
        item.expand = '+'
      }
    }
  },
  watch: {
    data (val) {
      const value = JSON.parse(JSON.stringify(val))
      this.initVal(value)
      this.nestableItems = value
    }
  }
}
</script>