<template>
  <!--
    Switch
      checked: true/false
      inline: true/false
      size: sm, lg
      type: primary, secondary, success, info, warning, danger, dark, light
  -->
  <div :class="classes">
    <input
      type="checkbox"
      class="custom-control-input"
      :name="name"
      :id="`switch-id-${name}`"
      :checked="checked"
    />
    <label class="custom-control-label" :for="`switch-id-${name}`">{{ label }}</label>
  </div>
</template>
<script>
export default {
  props: {
    label: {
      type: String,
      default: ''
    },
    checked: {
      type: Boolean,
      default: false
    },
    inline: {
      type: Boolean,
      default: false
    },
    size: {
      type: String,
      default: null
    },
    type: {
      type: String,
      default: 'primary'
    },
    name: {
      type: String
    }
  },
  computed: {
    classes() {
      let cls = `custom-control custom-switch custom-control-${this.type} `
      if (this.size) {
        cls += `custom-control-${this.size} `
      }
      if (this.inline) {
        cls += 'custom-control-inline '
      }
      return cls
    }
  }
}
</script>