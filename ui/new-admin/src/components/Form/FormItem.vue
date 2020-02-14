<template>
  <!--
    form-item
      label: string
      help: string
      state: success, error
  -->
  <div class="form-group">
    <label>{{ label }}</label>
    <slot v-bind:state="formState"></slot>
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
    }
  },
  data() {
    return {
      formState: null
    }
  },
  created() {
    this.formState = this.state;
  }
}
</script>