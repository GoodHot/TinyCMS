<template>
  <!--
    form-item
      label: string
      help: string
      state: success, error
  -->
  <div :class="classes">
    <label v-if="label" :style="label? '' : 'margin-bottom:1.5rem'">{{ label }}</label>
    <div>
      <slot v-bind:state="selfState"></slot>
    </div>
    <small class="form-text text-muted" v-if="help">{{ help }}</small>
  </div>
</template>
<script>
export default {
  name: 'TFormItem',
  props: {
    label: {
      type: String
    },
    help: {
      type: String,
      default: null
    },
    state: {
      type: String,
      default: null
    },
    col: {
      type: String,
      default: "12"
    }
  },
  data() {
    return {
      selfState: null
    }
  },
  computed: {
    classes() {
      let cls = 'form-group '
      if (this.col) {
        cls += 'col-md-' + this.col
      }
      return cls
    }
  },
  created() {
    this.selfState = this.state;
  }
}
</script>