<template>
  <div class="t-input-drawer">
    <b-input-group>
      <b-form-input v-model="dataValue" :readonly="readonly" :class="inputClass" :size="size" :placeholder="placeholder" @focus="showDrawer" @blur="hideDrawer" autocomplete="off"></b-form-input>
      <b-input-group-append>
        <b-button :variant="clearVariant" :size="size" @click="$emit('clear')">清空</b-button>
      </b-input-group-append>
    </b-input-group>
    <div class="drawer" v-show="drawerVisiable" tabindex="95270001" @focus="focusDrawer" @blur="drawerVisiable = false">
      <slot></slot>
    </div>
  </div>
</template>
<script>
import { PropTypes } from '@/utils/types'

export default {
  model: {
    prop: 'dataValue',
    event: 'change.value'
  },
  props: {
    size: PropTypes.String,
    placeholder: PropTypes.String,
    inputClass: PropTypes.String,
    clearVariant: PropTypes.String.def('light'),
    readonly: PropTypes.Boolean,
    dataValue: [Object, String, Number]
  },
  data () {
    return {
      drawerVisiable: false,
      hideTimer: null,
    }
  },
  mounted () {
  },
  methods: {
    showDrawer () {
      this.drawerVisiable = true
      this.$emit('onshow')
    },
    hideDrawer () {
      const _this = this
      this.hideTimer = setTimeout(() => {
        _this.drawerVisiable = false
      }, 50)
    },
    focusDrawer () {
      clearTimeout(this.hideTimer)
    }
  }
}
</script>